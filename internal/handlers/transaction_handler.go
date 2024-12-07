package handlers

import (
	"encoding/json"
	"github.com/achepkin/banklite/internal/domain/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandler struct {
	transactionService TransactionService
}

func NewTransactionHandler(transactionService TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	accountID := vars["id"]

	var req struct {
		Type   string
		Amount float64
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Type != "deposit" && req.Type != "withdraw" {
		http.Error(w, "invalid type", http.StatusBadRequest)
		return
	}

	txType := entity.TxType(req.Type)

	tx, err := h.transactionService.CreateTransaction(r.Context(), accountID, txType, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(tx.ID)

}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]

	transactions, err := h.transactionService.GetTransactions(nil, accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FromAccountID string  `json:"from_account_id"`
		ToAccountID   string  `json:"to_account_id"`
		Amount        float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.FromAccountID == "" || req.ToAccountID == "" {
		http.Error(w, "invalid from or to", http.StatusBadRequest)
		return
	}

	transfer, err := h.transactionService.Transfer(r.Context(), req.FromAccountID, req.ToAccountID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(transfer)
}
