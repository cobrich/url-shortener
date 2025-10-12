package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/cobrich/url-shortener/dtos"
	"github.com/cobrich/url-shortener/storage"
	"github.com/cobrich/url-shortener/utils"
)

type Handler struct {
	Urls *map[string]string
	Mu   *sync.Mutex
}

func NewHandler() *Handler {
	return &Handler{Urls: &storage.Urls, Mu: &sync.Mutex{}}
}

func (h *Handler) GetLongURLHundler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	key := r.PathValue("short_code")

	h.Mu.Lock()
	value, ok := (*h.Urls)[key]
	h.Mu.Unlock()

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Long URL Not Found for short key!"))
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(value))
}

func (h *Handler) CreateShortURLHundler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := dtos.ResponseCreateShortURLDTO{}
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	url := resp.Url

	ok := utils.IsUrlReachable(url)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short_code, err := utils.Generate_Short_Code()
	if err != nil {
		req := dtos.RequestErrorDTO{Error: err.Error()}
		json.NewEncoder(w).Encode(req)
		return
	}
	h.Mu.Lock()
	(*h.Urls)[short_code] = url
	h.Mu.Unlock()

	req := dtos.RequestCreateShortURLDTO{Short_code: short_code}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
