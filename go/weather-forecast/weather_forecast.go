// Package weather provides a function that returns
// CurrentCondition and CurrentLocation in a nice human-readable format.
package weather

// CurrentCondition variable stores Weather Condition to pass to Forecast function.
var CurrentCondition string

// CurrentLocation variable stores Weather to pass into Forecast function.
var CurrentLocation string

// Forecast function returns Goblin readable weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
