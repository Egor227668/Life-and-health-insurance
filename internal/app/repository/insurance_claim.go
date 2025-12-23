package repository

import (
    "time"
    "life-and-health-insurance/internal/app/model"
    "github.com/jmoiron/sqlx"
)

type InsuranceClaimRepository struct {
    db *sqlx.DB
}

func NewInsuranceClaimRepository(db *sqlx.DB) *InsuranceClaimRepository {
    return &InsuranceClaimRepository{db: db}
}

func (r *InsuranceClaimRepository) SubmitInsuranceClaim(claimNumber string, policyID int, submissionDate time.Time, description string, status string, claimAmount float64) error {
    query := `
    INSERT INTO insurance_claims(claim_number, policy_id, submission_date, description, status, claim_amount)
    VALUES (:claim_number, :policy_id, :submission_date, :description, :status, :claim_amount)`
    
    params := map[string]interface{}{
        "claim_number":    claimNumber,
        "policy_id":       policyID,
        "submission_date": submissionDate,
        "description":     description,
        "status":          status,
        "claim_amount":    claimAmount,
    }
    
    _, err := r.db.NamedExec(query, params)
    if err != nil {
        return err
    }
    return nil
}

func (r *InsuranceClaimRepository) CalculateCompensation(claimID int) (float64, error) {
    query := `
    SELECT claim_amount 
    FROM insurance_claims 
    WHERE claim_id = :claim_id`
    
    params := map[string]interface{}{
        "claim_id": claimID,
    }
    
    var amount float64
    err := r.db.Get(&amount, query, params)
    if err != nil {
        return 0, err
    }
    return amount, nil
}