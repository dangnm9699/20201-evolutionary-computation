package tsptw

import (
	"log"
	"math"
)

var (
	DM [][]float64  // Distance matrix
	NC int          // Number of cities
	TW [][2]float64 // Time windows
	P  [][]float64  // Pheromone matrix
	// Parameters
	m            int     = 3 // Number of ants
	Q0           float64 = 0.85
	theta, omega float64 = 0.1, 0.1
	beta         float64 = 0.5
	gamma        float64 = 3.0
	lamda, delta float64 = 0.05, 0.05
	BestCost     float64 = math.MaxFloat64
	BestTour     []int
	BestTime     float64 = math.MaxFloat64
	is_updated   bool    = false
)

func Config(bspb BSPB) {
	getData(bspb)
	log.Println(NC)
	for i := 0; i < NC; i++ {
		log.Println(DM[i])
	}
	for i := 0; i < NC; i++ {
		log.Println(TW[i])
	}
	P = make([][]float64, NC)
	for i := 0; i < NC; i++ {
		P[i] = make([]float64, NC)
		for j := 0; j < NC; j++ {
			P[i][j] = 1.0
		}
	}
}

func ACO(file BSPB) {
	Config(file)
	for it := 0; true ; it++ {
		ants := make([]Ant, m)
		for k := 0; k < m; k++ {
			ants[k] = new_ant()
		}
		for k := 0; k < m; k++ {
			ants[k].construct_route()
		}
		// Global update
		if is_updated {
			for i := 1; i < len(BestTour); i++ {
				cur := BestTour[i-1]
				nex := BestTour[i]
				P[cur][nex] = (1-theta)*P[cur][nex] + theta/BestTime
			}
			return
		}
	}
}
