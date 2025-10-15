package utils

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect bool
	}{
		{
			name:   "Валидный HTTPS URL",
			input:  "https://google.com",
			expect: true,
		},
		{
			name:   "Валидный HTTP URL с путем и query-параметрами",
			input:  "http://localhost:8080/path?query=123",
			expect: true,
		},
		{
			name:   "Невалидный: другая схема (ftp)",
			input:  "ftp://files.example.com",
			expect: false,
		},
		{
			name:   "Невалидный: отсутствует схема",
			input:  "google.com",
			expect: false,
		},
		{
			name:   "Невалидный: относительный путь",
			input:  "/relative/path",
			expect: false,
		},
		{
			name:   "Невалидный: полная белиберда",
			input:  "это не url",
			expect: false,
		},
		{
			name:   "Невалидный: отсутствует хост",
			input:  "http://",
			expect: false,
		},
		{
			name:   "Невалидный: пустая строка",
			input:  "",
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidUrl(tc.input)
			if result != tc.expect {
				t.Errorf("Для URL '%s' ожидалось %v, получено %v", tc.input, tc.expect, result)
			}
		})
	}
}

// MockHTTPClient - наша фальшивая реализация HTTP-клиента
type MockHTTPClient struct {
	// Задаем, какой ответ и ошибку он должен вернуть
	StatusCode int
	Err        error
}

// Head реализует интерфейс HTTPClient
func (c *MockHTTPClient) Head(url string) (*http.Response, error) {
	if c.Err != nil {
		return nil, c.Err
	}
	// Создаем фальшивый ответ
	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       io.NopCloser(strings.NewReader("")), // Пустое тело
	}, nil
}

func TestIsUrlReachable(t *testing.T) {
	testCases := []struct {
		name         string
		inputURL     string
		mockClient   HTTPClient // Будем передавать разные моки
		expectResult bool
	}{
		{
			name:     "Успех: URL доступен (200 OK)",
			inputURL: "https://valid.com",
			mockClient: &MockHTTPClient{
				StatusCode: 200,
			},
			expectResult: true,
		},
		{
			name:     "Успех: URL доступен (302 Redirect)",
			inputURL: "https://valid.com",
			mockClient: &MockHTTPClient{
				StatusCode: 302,
			},
			expectResult: true,
		},
		{
			name:     "Неудача: Сервер вернул 404 Not Found",
			inputURL: "https://valid.com",
			mockClient: &MockHTTPClient{
				StatusCode: 404,
			},
			expectResult: false,
		},
		{
			name:     "Неудача: Ошибка сети (например, таймаут)",
			inputURL: "https://valid.com",
			mockClient: &MockHTTPClient{
				Err: errors.New("connection timeout"),
			},
			expectResult: false,
		},
		{
			name:     "Неудача: URL синтаксически невалиден",
			inputURL: "not-a-url",
			// Клиент даже не будет вызван, так что можно передать nil
			mockClient:   nil,
			expectResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsUrlReachable(tc.inputURL, tc.mockClient)
			if result != tc.expectResult {
				t.Errorf("Ожидаемый результат %v, получено %v", tc.expectResult, result)
			}
		})
	}
}
