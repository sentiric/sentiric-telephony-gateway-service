package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "github.com/sentiric/sentiric-telephony-gateway-service/internal/call"
    "github.com/sentiric/sentiric-telephony-gateway-service/internal/util"
)

func main() {
    _ = godotenv.Load()

    port := os.Getenv("TELEPHONY_GATEWAY_SERVICE_PORT")
    if port == "" {
        port = "8080"
    }

    util.InitLogger()

    r := mux.NewRouter()
    r.HandleFunc("/health", call.HealthHandler).Methods("GET")
    r.HandleFunc("/call/start", call.StartCallHandler).Methods("POST")
    r.HandleFunc("/call/end", call.EndCallHandler).Methods("POST")

    log.Printf("Gateway service starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
