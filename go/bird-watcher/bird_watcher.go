package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	rt := 0
	for i := 0; i < len(birdsPerDay); i++ {
		rt += birdsPerDay[i]
	}
	return rt
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekrt := 0
	for i := week*7 - 7; i < week*7; i++ {
		weekrt += birdsPerDay[i]
	}
	return weekrt
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
