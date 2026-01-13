package handler

import (
    "encoding/json"
    "net/http"
    "life-and-health-insurance/internal/app/handler/dto"
    "life-and-health-insurance/internal/app/service"
)

type InsuranceClaimHandler struct {
    claimService *service.InsuranceClaimService
}

func NewInsuranceClaimHandler(claimService *service.InsuranceClaimService) *InsuranceClaimHandler {
    return &InsuranceClaimHandler{claimService: claimService}
}

// SubmitInsuranceClaim godoc
// @Summary Подача заявления о страховом случае
// @Description Подает новое заявление о страховом случае
// @Tags Страховые случаи
// @Accept json
// @Produce plain
// @Param request body dto.SubmitInsuranceClaimRequest true "Данные страхового случая"
// @Success 200 {string} string "Claim submitted"
// @Failure 400 {string} string "Invalid json submit"
// @Failure 500 {string} string "Claim cant submit"
// @Router /claims [post]

func (h *InsuranceClaimHandler) SubmitInsuranceClaim(w http.ResponseWriter, r *http.Request) {
    var request dto.SubmitInsuranceClaimRequest
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        w.WriteHeader(400)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Invalid json submit"))
        return
    }

    err = h.claimService.SubmitInsuranceClaim(
        request.ClaimNumber,
        request.PolicyID,
        request.SubmissionDate,
        request.Description,
        request.Status,
        request.ClaimAmount,
    )
    if err != nil {
        w.WriteHeader(500)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Claim cant submit"))
        return
    }

    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Claim submitted"))
}

func (h *InsuranceClaimHandler) CalculateCompensation(w http.ResponseWriter, r *http.Request) {
    var request dto.CalculateCompensationRequest
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        w.WriteHeader(400)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Invalid json calculate"))
        return
    }

    amount, err := h.claimService.CalculateCompensation(request.ClaimID)
    if err != nil {
        w.WriteHeader(500)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Cant calculate compensation"))
        return
    }

    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Compensation calculated: " + string(amount)))
}
