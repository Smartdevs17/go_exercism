// Package weather is use for checking the current weather condition current location that is provided.
package weather

var (
    // CurrentCondition and CurrentLocation are the vairables are used in the application.
	CurrentCondition string
    // CurrentLocation describes the current location of the person. CurrentCondition describes the current condition of the weather.
	CurrentLocation  string
)

// Forecast takes in a city and the weather condition and returns a string with is the joining of the current location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
