/*
Package weather provides tools for us to give weather forecast for Goblinocus.
*/
package weather

/*
CurrentCondition holds the value for the current weather condition in Goblinocus.
*/
var CurrentCondition string

/*
CurrentLocation holds the value for the current location in Goblinocus.
*/
var CurrentLocation string

/*
Forecast function that tells us the current weather condition for a location in Goblinocus.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
