func carFleet(target int, position []int, speed []int) int {
	// Map positions directly to speeds using target as upper bound
	// 0 means no car at that position
	timeToTarget := make([]float64, target+1)
	for i := 0; i < len(position); i++ {
		timeToTarget[position[i]] = float64(target-position[i]) / float64(speed[i])
	}

	fleets := 0
	var maxTime float64

	// Iterate backwards from the target position to position 0
	for pos := target; pos >= 0; pos-- {
		time := timeToTarget[pos]
		if time > 0 { // Car exists at this position
			if time > maxTime {
				maxTime = time
				fleets++
			}
		}
	}

	return fleets
}