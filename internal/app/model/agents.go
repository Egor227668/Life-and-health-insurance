package models

import "time"

type Agent struct {
    AgentID     int       `db:"agent_id"`
    FullName    string    `db:"full_name"`
    Position    string    `db:"position"`
    Phone       string    `db:"phone"`
    Email       string    `db:"email"`
    WorkShift   string    `db:"work_shift"`
    EmployeeID  string    `db:"employee_id"`
}