type Car struct{
	pos int
	time float64
}
func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	// calculate time of cars
	for i, v:= range position{
		car := Car {
			pos:v,
			time: float64(target-v)/float64(speed[i]),
		}
		cars = append(cars, car)
	}
	// target  = 10 
	//position = [4,1  ,0, 7], 
	//   speed = [2,2  ,1, 1]
	//   time  = [3,4.5,10,3]
	// when become a fleet,
	// if previous car time > ahead car time
	// sort the position and time
	sort.Slice(cars, func(i, j int) bool {
        return cars[i].pos > cars[j].pos
    })

	//position = [0,1  ,4, 7], 
	//   speed = [2,2  ,1, 1]
	//   time  = [10,4.5,3,3]
	fleets := 0
	maxTime :=0.0
	for _, car := range cars{
		if car.time > maxTime{
			fleets ++
			maxTime = car.time
		}
	}
	return fleets
}
