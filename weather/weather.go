package weather

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type Client struct {
}

func init() {
	// register function handler
	functions.HTTP("sk-weather", handleWeatherRequest)
}

func handleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, SolarKit!"))
}
