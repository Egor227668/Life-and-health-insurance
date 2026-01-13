package handler

import (
    "encoding/json"
    "net/http"
    "life-and-health-insurance/internal/app/handler/dto"
    "life-and-health-insurance/internal/app/service"
)

type PolicyHandler struct {
    policyService *service.PoliciesService
}

func NewPolicyHandler(policyService *service.PoliciesService) *PolicyHandler {
    return &PolicyHandler{policyService: policyService}
}

// CreatePolicy godoc
// @Summary Создание нового страхового полиса
// @Description Создает новый страховой полис с указанными параметрами
// @Tags Полисы
// @Accept json
// @Produce plain
// @Param request body dto.CreatePolicyRequest true "Данные для создания полиса"
// @Success 200 {string} string "Policy created"
// @Failure 400 {string} string "Invalid json create"
// @Failure 500 {string} string "Policy cant create"
// @Router /policies [post]

func (h *PolicyHandler) CreatePolicy(w http.ResponseWriter, r *http.Request) {
    var request dto.CreatePolicyRequest
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        w.WriteHeader(400)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Invalid json create"))
        return
    }

    err = h.policyService.CreatePolicy(
        request.PolicyNumber,
        request.ClientID,
        request.AgentID,
        request.TypeID,
        request.StartDate,
        request.EndDate,
        request.Status,
        request.CoverageAmount,
        request.Premium,
    )
    if err != nil {
        w.WriteHeader(500)
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Policy cant create"))
        return
    }

    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Policy created"))
}
