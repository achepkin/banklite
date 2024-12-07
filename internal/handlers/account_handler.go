package handlers

import (
	"encoding/json"
	"errors"
	"github.com/achepkin/banklite/internal/domain"
	"github.com/achepkin/banklite/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type AccountHandler struct {
	accService AccountService
}

func NewAccountHandler(accService AccountService) *AccountHandler {
	return &AccountHandler{accService: accService}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Owner          string  `json:"owner"`
		InitialBalance float64 `json:"initial_balance"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := &entity.Account{
		ID:        uuid.New().String(),
		Owner:     req.Owner,
		Balance:   req.InitialBalance,
		CreatedAt: time.Now(),
	}

	err := h.accService.CreateAccount(nil, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")

	vars := mux.Vars(r)
	id := vars["id"]
	account, err := h.accService.GetAccount(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_ = json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accService.ListAccounts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(accounts)
}
