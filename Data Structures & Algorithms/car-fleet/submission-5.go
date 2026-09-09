type Car struct {
	position int
	speed int
}
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i] = Car{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})
	
	var maxTime float64
	fleets := 0
	for _, car := range cars {
		time := float64(target - car.position) / float64(car.speed)
		if time > maxTime {
			fleets++
			maxTime = time
		}
	}
	return fleets
}