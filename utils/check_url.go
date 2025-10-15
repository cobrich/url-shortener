package utils

import (
	"net/http"
	"net/url"
	"time"
)

func isValidUrl(rawUrl string) bool {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return false // Ошибка парсинга - точно невалидный URL
	}

	// Проверяем, что есть схема (http, https, ftp и т.д.) и хост (google.com)
	// Это отсекает относительные пути типа "/some/path"
	if u.Scheme == "" || u.Host == "" {
		return false
	}

	// Дополнительная проверка, если нужны только определенные схемы
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

// IsUrlReachable теперь принимает HTTPClient
func IsUrlReachable(rawUrl string, client HTTPClient) bool {
	if !isValidUrl(rawUrl) {
		return false
	}
    
    // Используем переданный клиент
	resp, err := client.Head(rawUrl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}

// Создадим "прод" версию, которую будем использовать в реальном коде
func IsUrlReachableProd(rawUrl string) bool {
    client := &http.Client{Timeout: 3 * time.Second}
    return IsUrlReachable(rawUrl, client)
}