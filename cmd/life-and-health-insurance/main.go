package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


func CreatePolicy(w http.ResponseWriter, r *http.Request) {
    policy := mux.Vars(r)
    w.Write([]byte("Оформлен новый страховой полис " + policy["id"]))
}

func SubmitInsuranceClaim(w http.ResponseWriter, r *http.Request) {
    claim := mux.Vars(r)
    w.Write([]byte("Подано заявление о страховом случае " + claim["id"]))
}

func CalculateCompensation(w http.ResponseWriter, r *http.Request) {
    calc := mux.Vars(r)
    w.Write([]byte("Рассчитана сумма выплаты для случая " + calc["id"]))
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/policies", CreatePolicy).Methods("POST")
    r.HandleFunc("/claims", SubmitInsuranceClaim).Methods("POST")
    r.HandleFunc("/calculate/{id}", CalculateCompensation).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}
