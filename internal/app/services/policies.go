package service

import (
	"life-and-health-insurance/internal/app/repository"
    "time"
)

type PoliciesService struct {
    policyRepository *repository.PolicyRepository
}

func NewPoliciesService(policyRepository *repository.PolicyRepository) *PoliciesService {
    return &PoliciesService{policyRepository: policyRepository}
}

func (s *PoliciesService) CreatePolicy(policyNumber string, clientID int, agentID int, typeID int, startDate time.Time, endDate time.Time, status string, coverageAmount float64, premium float64) error {
    return s.policyRepository.CreatePolicy(policyNumber, clientID, agentID, typeID, startDate, endDate, status, coverageAmount, premium)
}