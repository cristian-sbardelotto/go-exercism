package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (temperatureUnit TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[temperatureUnit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temperature Temperature) String() string {
	return fmt.Sprintf("%v %s", temperature.degree, temperature.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[speedUnit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return fmt.Sprintf("%v %s", speed.magnitude, speed.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteorologyData MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %v%% Humidity", meteorologyData.location, meteorologyData.temperature.String(), meteorologyData.windDirection, meteorologyData.windSpeed.String(), meteorologyData.humidity)
}
