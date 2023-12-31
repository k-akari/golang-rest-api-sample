package handler

import (
	"encoding/json"
	"net/http"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/pkg/validator"
)

type CompanyHandler struct {
	companyUsecase companyUsecase
}

func NewCompanyHandler(
	companyUsecase companyUsecase,
) *CompanyHandler {
	return &CompanyHandler{
		companyUsecase: companyUsecase,
	}
}

func (ch *CompanyHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Name            string `json:"name" validate:"required"`
		Representative  string `json:"representative" validate:"required"`
		TelephoneNumber string `json:"telephone_number" validate:"required"`
		PostalCode      string `json:"postal_code" validate:"required"`
		Address         string `json:"address" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	err := validator.Struct(b)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	company := domain.Company{
		Name:            b.Name,
		Representative:  b.Representative,
		TelephoneNumber: b.TelephoneNumber,
		PostalCode:      b.PostalCode,
		Address:         b.Address,
	}

	cid, err := ch.companyUsecase.CreateCompany(ctx, &company)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID domain.CompanyID `json:"id"`
	}{ID: cid}

	respondJSON(ctx, w, resp, http.StatusOK)
}

func (ch *CompanyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	coid, err := getCompanyIDFromCtx(ctx)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	c, err := ch.companyUsecase.GetCompanyByID(ctx, domain.CompanyID(coid))
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	respondJSON(ctx, w, c, http.StatusOK)
}
