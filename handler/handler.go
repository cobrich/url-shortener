package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/cobrich/url-shortener/dtos"
	"github.com/cobrich/url-shortener/storage"
	"github.com/cobrich/url-shortener/utils"
)

type Handler struct {
	storage *storage.Storage
}

func NewHandler(st *storage.Storage) *Handler {
	return &Handler{storage: st}
}

func (h *Handler) GetLongURLHundler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.PathValue("short_code")

	url, ok := h.storage.Get(code)

	if !ok {
		http.NotFound(w, r)
		return
	}

	go func() {
		log.Printf("INFO: redirected %s to %s", code, url)
	}()

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Handler) CreateShortURLHundler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req := dtos.RequestCreateShortURLDTO{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	url := req.Url

	ok := utils.IsUrlReachable(url)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code, err := utils.GenerateShortCode()
	if err != nil {
		req := dtos.RequestErrorDTO{Error: err.Error()}
		json.NewEncoder(w).Encode(req)
		return
	}

	h.storage.Save(code, url)

	resp := dtos.ResponseCreateShortURLDTO{ShortCode: code}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
