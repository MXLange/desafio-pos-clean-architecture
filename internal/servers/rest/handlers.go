package rest

import (
	"encoding/json"
	"net/http"

	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/dto"
	usecases "github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/use_cases"
	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
)

type Handlers struct {
	createOrderUseCase *usecases.CreateOrderUseCase
	listOrdersUseCase  *usecases.ListOrdersUseCase
}

func NewHandler(createOrderUseCase *usecases.CreateOrderUseCase, listOrdersUseCase *usecases.ListOrdersUseCase) (*Handlers, error) {
	if createOrderUseCase == nil {
		return nil, e.ErrNilCreateOrderUseCase
	}
	if listOrdersUseCase == nil {
		return nil, e.ErrNilListOrdersUseCase
	}
	return &Handlers{
		createOrderUseCase: createOrderUseCase,
		listOrdersUseCase:  listOrdersUseCase,
	}, nil
}

func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var orderRequest dto.OrderCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	res, err := h.createOrderUseCase.Execute(r.Context(), &orderRequest)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func (h *Handlers) ListOrders(w http.ResponseWriter, r *http.Request) {
	res, err := h.listOrdersUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, "Failed to list orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
