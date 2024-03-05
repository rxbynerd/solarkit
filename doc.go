package solarkit

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"github.com/rxbynerd/solarkit/weather"
)

func init() {
	// register function handler
	functions.HTTP("solarkitWeather", weather.HandleWeatherRequest)
}
