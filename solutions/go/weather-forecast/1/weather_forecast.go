// Package weather provides current weather conditions for cities.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string
// CurrentLocation stores the current city.
var CurrentLocation string

// Forecast returns the weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}