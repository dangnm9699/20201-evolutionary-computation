package tsp

import "math"

func fitness(ind []int) (totalDistance float64) {
	for i, a := range ind {
		//diffX := LocationData[a-1][0] - LocationData[ind[(i+1)%NumCities]-1][0]
		//diffY := LocationData[a-1][1] - LocationData[ind[(i+1)%NumCities]-1][1]
		//totalDistance += distance(float64(diffX), float64(diffY))
		totalDistance += DistanceMatrix[a][ind[(i+1)%NumCities]]
	}
	return
}

func distance(diffX, diffY float64) float64 {
	return math.Sqrt(diffX*diffX + diffY*diffY)
}
