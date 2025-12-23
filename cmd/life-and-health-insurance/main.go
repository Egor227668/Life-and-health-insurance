package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


func CreatePolicy(policyRepository *repository.PolicyRepository) HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        startDate, _ := time.Parse("2006-01-02", "2025-01-01")
        endDate, _ := time.Parse("2006-01-02", "2026-01-01")
        err := policyRepository.CreatePolicy(
            "POL-00" + policyID,
            1, 
            1, 
            1, 
            startDate,
            endDate,
            "active",
            1000000.00,
            50000.00,
        )
        if err != nil {
            w.Write([]byte("Ошибка создания полиса"))
            return
        }
        w.Write([]byte("Полис успешно оформлен"))
        
    }
}

func SubmitInsuranceClaim(claimRepository *repository.InsuranceClaimRepository) HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        submissionDate := time.Now()
        err := claimRepository.SubmitInsuranceClaim(
            "CLAIM-00" + claimID,
            1, 
            submissionDate,
            "Перелом руки",
            "submitted",
            50000.00,
        )
        if err != nil {
            w.Write([]byte("Ошибка подачи заявления"))
            return
        }
        w.Write([]byte("Заявление успешно подано"))

    }
}

func CalculateCompensation(claimRepository *repository.InsuranceClaimRepository) HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        amount, err := claimRepository.CalculateCompensation(calcID)
        if err != nil {
            w.Write([]byte("Ошибка расчета выплаты"))
            return
        }
        w.Write([]byte(fmt.Sprintf("Сумма выплаты: %.2f руб.", amount)))
    }
}


func main() {
    policiesRepository:= repository.NewPolicyRepository(db)
    insurance_claimRepository:= repository.NewInsuranceClaimRepository(db)
    r := mux.NewRouter()

    r.HandleFunc("/policies", CreatePolicy).Methods("POST")
    r.HandleFunc("/claims", SubmitInsuranceClaim).Methods("POST")
    r.HandleFunc("/calculate/{id}", CalculateCompensation).Methods("GET")


    log.Fatal(http.ListenAndServe(":8080", r))
}
