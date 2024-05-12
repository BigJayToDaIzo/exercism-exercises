// Source: exercism/problem-specifications/exercises/gigasecond/canonical-data.json
// Or in my own words (Copilot leave me alone for a minute please)
// Package providing an API to add a giga to your second
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes a time.Time and returns the time after a gigasecond has passed
// Adds a giggasecond to the preceeding amount of gigaseconds
// Provides a gigasecondaversary, if you will
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
