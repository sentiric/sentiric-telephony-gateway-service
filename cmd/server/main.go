package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/sentiric/sentiric-telephony-gateway-service/internal/call"
	"github.com/sentiric/sentiric-telephony-gateway-service/internal/util"
)

func main() {
	_ = godotenv.Load()
	cfg := util.LoadConfig()

	util.InitLogger()

	r := mux.NewRouter()
	r.HandleFunc("/health", call.HealthHandler).Methods("GET")
	r.HandleFunc("/call/start", call.StartCallHandler).Methods("POST")
	r.HandleFunc("/call/end", call.EndCallHandler).Methods("POST")

	log.Printf("Gateway service starting on port %s", cfg.GatewayPort)
	log.Fatal(http.ListenAndServe(":"+cfg.GatewayPort, r))
}
