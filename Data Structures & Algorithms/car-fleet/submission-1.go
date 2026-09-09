func carFleet(target int, position []int, speed []int) int {
    type Car struct {
        position int
        speed    int
    }

    cars := make([]Car, len(position))

    for i := range position {
        cars[i] = Car{
            position: position[i],
            speed:    speed[i],
        }
    }

    sort.Slice(cars, func(i, j int) bool {
        return cars[i].position > cars[j].position
    })

    fleets := []float64{}

    for _, car := range cars {
        time := float64(target-car.position) / float64(car.speed)

        if len(fleets) == 0 || time > fleets[len(fleets)-1] {
            fleets = append(fleets, time)
        }
        // Otherwise, this car joins the fleet ahead.
    }

    return len(fleets)
}