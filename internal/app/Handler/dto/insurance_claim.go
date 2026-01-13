package dto

import "time"

type CreateInsuranceClaimRequest struct {
    ClaimNumber    string    `json:"claim_number"`
    PolicyID       int       `json:"policy_id"`
    SubmissionDate time.Time `json:"submission_date"`
    Description    string    `json:"description"`
    Status         string    `json:"status"`
    ClaimAmount    float64   `json:"claim_amount"`
}

type CreateInsuranceClaimResponse struct {
    ClaimID        int       `json:"claim_id"`
    ClaimNumber    string    `json:"claim_number"`
    PolicyID       int       `json:"policy_id"`
    SubmissionDate time.Time `json:"submission_date"`
    Description    string    `json:"description"`
    Status         string    `json:"status"`
    ClaimAmount    float64   `json:"claim_amount"`
    CreatedAt      time.Time `json:"created_at"`
}
type CalculateCompensationRequest struct {
    ClaimID int `json:"claim_id"`
}

type CalculateCompensationResponse struct {
    ClaimID      int       `json:"claim_id"`
    ClaimNumber  string    `json:"claim_number"`
    Amount       float64   `json:"amount"`
    CalculatedAt time.Time `json:"calculated_at"`
}