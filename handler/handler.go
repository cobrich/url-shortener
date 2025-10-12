package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cobrich/url-shortener/dtos"
	"github.com/cobrich/url-shortener/storage"
	"github.com/cobrich/url-shortener/utils"
)

type Handler struct {
	storage *storage.Storage
	Mu   *sync.RWMutex
}

func NewHandler(st *storage.Storage) *Handler {
	return &Handler{storage: st, Mu: &sync.RWMutex{}}
}

func (h *Handler) GetLongURLHundler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	code := r.PathValue("short_code")

	url, ok := h.storage.Get(code)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Long URL Not Found for short key!"))
	}

	go func() {
		log.Printf("INFO: redirected %s to %s", code, url)
	}()

	w.WriteHeader(http.StatusFound)
	w.Write([]byte(url))
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

	code, err := utils.Generate_Short_Code()
	if err != nil {
		req := dtos.RequestErrorDTO{Error: err.Error()}
		json.NewEncoder(w).Encode(req)
		return
	}

	h.storage.Save(code, url)

	req := dtos.RequestCreateShortURLDTO{Short_code: code}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
