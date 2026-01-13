package models

import "time"

type Coefficient struct {
    CoefficientID   int       `db:"coefficient_id"`
    Name            string    `db:"name"`
    Value           float64   `db:"value"`
    Description     string    `db:"description"`
    IsActive        bool      `db:"is_active"`
}