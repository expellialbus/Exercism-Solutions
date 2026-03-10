// Package weather provides tools to get weather forecasts.
package weather

var (
    // CurrentCondition variable stores the current condition of the weather.
	CurrentCondition string
	// CurrentLocation variable stores the current location which the CurrentCondition belongs to.
    CurrentLocation  string
)

// Forecast function take city and condition values and returns the weather forecast for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
