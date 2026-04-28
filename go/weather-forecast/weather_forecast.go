// Package weather provides tools
// to get current forecast for any city.
package weather

// CurrentCondition represents the current weather for the city.
var CurrentCondition string

// CurrentLocation represents the city of forecast.
var CurrentLocation string

// Forecast takes city and condition to return a string explaining the current weather condition for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
