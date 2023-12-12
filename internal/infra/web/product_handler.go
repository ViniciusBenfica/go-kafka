package web

import (
	"encoding/json"
	"mensageria/internal/usecase"
	"net/http"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductUseCase   *usecase.ListProductUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductUseCase:   listProductUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) ListProducsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
