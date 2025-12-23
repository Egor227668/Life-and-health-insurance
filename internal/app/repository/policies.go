package repository

import (
    "time"
    "life-and-health-insurance/internal/app/model"
    "github.com/jmoiron/sqlx"
)

type PolicyRepository struct {
    db *sqlx.DB
}

func NewPolicyRepository(db *sqlx.DB) *PolicyRepository {
    return &PolicyRepository{db: db}
}

func (r *PolicyRepository) CreatePolicy(policyNumber string, clientID int, agentID int, typeID int, startDate time.Time, endDate time.Time, status string, coverageAmount float64, premium float64) error {
    query := `
    INSERT INTO policies(policy_number, client_id, agent_id, type_id, start_date, end_date, status, coverage_amount, premium)
    VALUES (:policy_number, :client_id, :agent_id, :type_id, :start_date, :end_date, :status, :coverage_amount, :premium)`
    
    params := map[string]interface{}{
        "policy_number":   policyNumber,
        "client_id":       clientID,
        "agent_id":        agentID,
        "type_id":         typeID,
        "start_date":      startDate,
        "end_date":        endDate,
        "status":          status,
        "coverage_amount": coverageAmount,
        "premium":         premium,
    }
    
    _, err := r.db.NamedExec(query, params)
    if err != nil {
        return err
    }
    return nil
}