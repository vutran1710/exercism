package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")
	var total int

	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
	var total int
	startIdx := 7 * (week - 1)
	endIdx := startIdx + 7

	for _, value := range birdsPerDay[startIdx:endIdx] {
		total += value
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
	for day, value := range birdsPerDay {
		birdsPerDay[day] = value + (day+1)%2
	}

	return birdsPerDay
}
