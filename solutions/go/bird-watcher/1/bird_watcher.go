package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int

	for i := 0; i <= len(birdsPerDay)-1; i++ {
		sum += birdsPerDay[i]
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysToRemove := (week - 1) * 7
	var sum int

	newSlice := birdsPerDay[daysToRemove : daysToRemove+7]

	for i := 0; i <= len(newSlice)-1; i++ {
		sum += newSlice[i]
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i <= len(birdsPerDay)-1; i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
