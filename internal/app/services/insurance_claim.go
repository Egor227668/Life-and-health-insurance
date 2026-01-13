package service

import (
    "life-and-health-insurance/internal/app/model"
    "life-and-health-insurance/internal/app/repository"    
)

type InsuranceClaimService struct {
    claimRepository *repository.InsuranceClaimRepository
}

func NewInsuranceClaimService(claimRepository *repository.InsuranceClaimRepository) *InsuranceClaimService {
    return &InsuranceClaimService{claimRepository: claimRepository}
}

func (s *InsuranceClaimService) GetInsuranceClaims() []model.InsuranceClaim {
    return s.claimRepository.GetInsuranceClaims()
}