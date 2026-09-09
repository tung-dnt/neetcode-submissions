type Car struct {
	pos  int
	spd int
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
	sort.Slice(cars, func(i,j int) bool {
		return cars[i].pos > cars[j].pos
	})
	fleets := 0
	var maxTime float64
	for _, car := range cars {
		eta := float64(target - car.pos) / float64(car.spd)
		if eta > maxTime {
			maxTime = eta 
			fleets++
		}
	}
	return fleets
}