package dtos

type ResponseCreateShortURLDTO struct{
	Url string `json:"url"` 
}

type RequestCreateShortURLDTO struct{
	Short_code string `json:"short_code"`
}

type RequestErrorDTO struct{
	Error string `json:"error"`
}