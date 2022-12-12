/*
Package weather: Because we're dealing with a code for the Goblinocus weather report, this is the package title of this go file.
The package contains details about the weather condition in Goblinocus.
*/
package weather

// CurrentCondition holds the value for the current weather condition in Goblinocu.
var CurrentCondition string

// CurrentLocation holds the value for the current location in Goblinocus.
var CurrentLocation string

// Forecast function that tells us the current weather condition for a location in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
