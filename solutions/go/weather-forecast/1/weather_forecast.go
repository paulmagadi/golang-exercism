// Package weather provides the tools for forcasting the weather.
package weather

var (
    // CurrentCondition variable holds the current weather condition of the city.
	CurrentCondition string
    // CurrentLocation holds the name of the city.
	CurrentLocation  string
)

// Forecast returns the current weather condition of the current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
