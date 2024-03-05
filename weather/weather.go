package weather

import (
	"net/http"
)

type Client struct {
}

func HandleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, SolarKit!"))
}
