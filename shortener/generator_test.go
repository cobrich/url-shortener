package shortener

import (
	"testing"
	"unicode/utf8"
)

func TestGenerateShortCode_CorrectLength(t *testing.T) {
	testCases := []struct {
		name         string
		inputLength  int
		expectLength int
	}{
		{"Длина 8", 8, 8},
		{"Длина 10", 10, 10},
		{"Длина 0", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Вызываем вашу новую, исправленную функцию
			result, err := GenerateSecureString(tc.inputLength)
			
			// Проверяем, что не было ошибки
			if err != nil {
				t.Fatalf("Функция вернула неожиданную ошибку: %v", err)
			}
			
			resultLength := utf8.RuneCountInString(result)
			
			if resultLength != tc.expectLength {
				t.Errorf("Ожидаемая длина: %d, получено: %d", tc.expectLength, resultLength)
			}
		})
	}
}