// Package weather provides information about the forecast for your city.
package weather

// CurrentCondition represents the current condition of your city.
var CurrentCondition string

// CurrentLocation represents the current location that you are in.
var CurrentLocation string

// Forecast function gather all the Forecast information and show it to the user.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
