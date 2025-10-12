package handler

import (
	"encoding/json"
	"log"
	"net/http"

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
	req := dtos.RequestCreateShortURLDTO{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	url := req.Url

	ok := utils.IsUrlReachable(url)
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "The provided URL is not reachable")
		return
	}

	var code string
	for i := 0; i < 10; i++ {
		code, err = utils.GenerateShortCode()
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate short code")
			return
		}

		_, ok = h.storage.Get(code)
		if !ok {
			break
		}

		if i == 9 {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate a unique short code after several attempts")
			return
		}
	}

	h.storage.Save(code, url)

	resp := dtos.ResponseCreateShortURLDTO{ShortCode: code}
	utils.RespondWithJSON(w, http.StatusCreated, resp)
}
