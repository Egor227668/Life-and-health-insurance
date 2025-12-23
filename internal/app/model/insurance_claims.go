package models

import "time"

type InsuranceClaim struct {
    ClaimID         int       `db:"claim_id"`
    ClaimNumber     string    `db:"claim_number"`
    Policy          Policy    `db:"policy"`
    SubmissionDate  time.Time `db:"submission_date"`
    Description     string    `db:"description"`
    Status          string    `db:"status"`
    ClaimAmount     float64   `db:"claim_amount"`
}