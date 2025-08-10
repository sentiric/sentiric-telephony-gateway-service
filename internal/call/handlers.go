package call

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func StartCallHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("MOCK_MODE") == "true" {
		fmt.Fprintln(w, "Mock call started.")
		return
	}

	log.Println("StartCallHandler: initiating real SIP call")

	// Gerçek SIP isteği burada gönderilecek
	fmt.Fprintln(w, "Call started via SIP gateway.")
}

func EndCallHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("MOCK_MODE") == "true" {
		fmt.Fprintln(w, "Mock call ended.")
		return
	}

	log.Println("EndCallHandler: ending real SIP call")

	// Gerçek SIP isteği burada gönderilecek
	fmt.Fprintln(w, "Call ended via SIP gateway.")
}
