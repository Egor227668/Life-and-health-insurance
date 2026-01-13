package dto

import "time"

type CreatePolicyRequest struct {
    PolicyNumber   string    `json:"policy_number"`
    ClientID       int       `json:"client_id"`
    AgentID        int       `json:"agent_id"`
    TypeID         int       `json:"type_id"`
    StartDate      time.Time `json:"start_date"`
    EndDate        time.Time `json:"end_date"`
    Status         string    `json:"status"`
    CoverageAmount float64   `json:"coverage_amount"`
    Premium        float64   `json:"premium"`
}

type CreatePolicyResponse struct {
    PolicyID       int       `json:"policy_id"`
    PolicyNumber   string    `json:"policy_number"`
    ClientID       int       `json:"client_id"`
    AgentID        int       `json:"agent_id"`
    TypeID         int       `json:"type_id"`
    StartDate      time.Time `json:"start_date"`
    EndDate        time.Time `json:"end_date"`
    Status         string    `json:"status"`
    CoverageAmount float64   `json:"coverage_amount"`
    Premium        float64   `json:"premium"`
    CreatedAt      time.Time `json:"created_at"`
}