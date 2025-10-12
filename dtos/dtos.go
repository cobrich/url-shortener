package dtos

type ResponseCreateShortURLDTO struct {
	ShortCode string `json:"short_code"`
}

type RequestCreateShortURLDTO struct {
	Url string `json:"url"`
}