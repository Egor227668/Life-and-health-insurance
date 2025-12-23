package models

import "time"

type InsuranceType struct {
    TypeID      int       `db:"type_id"`
    Name        string    `db:"name"`
    Description string    `db:"description"`
    Risks       string    `db:"risks"`
    BaseCost    float64   `db:"base_cost"`
}