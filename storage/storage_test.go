package storage

import "testing"

func TestSave_Table(t *testing.T) {
	st := NewStorage()
	newUrl := "https://example.com"
	newCode := "newurl"
	code := "testcode"
	url := "https://test.com"

	testCases := []struct {
		name      string
		inputCode string
		inputUrl  string
	}{
		{"first", newCode, newUrl},
		{"second", code, url},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			st.Save(tc.inputCode, tc.inputUrl)

			// Здесь мы ВЫНУЖДЕНЫ проверить внутреннее состояние
			// Это единственный способ проверить Save в изоляции
			st.mu.RLock() // Нужно заблокировать мьютекс, чтобы избежать гонки данных
			defer st.mu.RUnlock()

			savedURL, ok := st.urls[tc.inputCode]
			if !ok {
				t.Fatalf("Save() не сохранил URL, ключ не найден")
			}
			if savedURL != tc.inputUrl {
				t.Errorf("Ожидаемый URL '%s', получено '%s'", tc.inputUrl, savedURL)
			}
		})
	}
}

func TestGet_Table(t *testing.T) {
	st := NewStorage()
	existingCode := "exists123"
	existingURL := "https://example.com"
	emptyURLCode := "emptyurl"
	st.Save(existingCode, existingURL)
	st.Save(emptyURLCode, "")

	testCases := []struct {
		name      string
		inputCode string
		expectURL string
		expectOk  bool
	}{
		{
			name:      "Успешное получение существующего URL",
			inputCode: existingCode,
			expectURL: existingURL,
			expectOk:  true,
		},
		{
			name:      "Попытка получить несуществующий URL",
			inputCode: "nonexistent",
			expectURL: "",    // Ожидаем пустую строку, так как ничего не найдено
			expectOk:  false, // Ожидаем, что ok будет false
		},
		{
			name:      "Получение существующего ключа с пустой строкой в качестве URL",
			inputCode: emptyURLCode,
			expectURL: "",
			expectOk:  true, // Ключ есть, поэтому ok должен быть true
		},
	}

	go func() {

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Выполняем тестируемую функцию
				url, ok := st.Get(tc.inputCode)

				// Проверяем результат
				if ok != tc.expectOk {
					t.Errorf("Ожидаемый 'ok' = %v, получено = %v", tc.expectOk, ok)
				}

				if url != tc.expectURL {
					t.Errorf("Ожидаемый URL = '%s', получено = '%s'", tc.expectURL, url)
				}
			})
		}
	}()
}
