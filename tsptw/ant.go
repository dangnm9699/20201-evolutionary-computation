package tsptw

import (
	"log"
	"math"
	"math/rand"
)

type Ant struct {
	Position  int
	Path      []int
	Cost      float64
	Time      float64
	Unvisited []int
	Delta     [][]float64
}

func (ant *Ant) construct_route() {
	for len(ant.Unvisited) > 0 {
		next_idx := ant.next_node()
		if next_idx < 0 {
			break
		}
		curr_node := ant.Position
		next_node := ant.Unvisited[next_idx]
		//
		ant.Path = append(ant.Path, next_node)
		ant.Cost += DM[curr_node][next_node]
		ant.Time = math.Max(ant.Time+DM[curr_node][next_node], TW[next_node][0])
		ant.Unvisited = remove_index(ant.Unvisited, next_idx)
		ant.Position = next_node
	}
	if len(ant.Path) == NC {
		updated = true
		ant.update_local()
		last_city := ant.Path[NC-1]
		ant.Cost += DM[last_city][0]
		ant.Time += DM[last_city][0]
		// Update best
		if ant.Time < BestTime {
			log.Printf("Cost: %7.2f ; Time: %7.2f; Path: %v\n", ant.Cost, ant.Time, ant.Path)
			BestTime = ant.Time
			BestCost = ant.Cost
			BestTour = ant.Path
		}
	}
}

func (ant *Ant) next_node() int {
	l := len(ant.Unvisited)
	//
	G := make([]float64, l)
	g := make([]float64, l)
	count_G := 0
	mean_G := 0.0
	//
	H := make([]float64, l)
	h := make([]float64, l)
	count_H := 0
	mean_H := 0.0
	//
	nex := -1
	cur := ant.Position
	//
	for i, adj := range ant.Unvisited {
		travel_time := ant.Time + DM[cur][adj]
		G[i] = TW[adj][1] - travel_time
		if G[i] >= 0 {
			count_G++
			mean_G += G[i]
		}
		H[i] = TW[adj][0] - travel_time
		if H[i] > 0 {
			count_H++
			mean_H += H[i]
		}
	}
	if count_G > 0 {
		mean_G /= float64(count_G)
	} else {
		//le := len(ant.Path) - 1
		//P[le-1][le] *= 0.
		return nex
	}
	if count_H >= 0 {
		mean_H /= float64(count_H)
	}
	//
	for i := range ant.Unvisited {
		if G[i] >= 0 {
			g[i] = 1.0 / (1.0 + math.Exp(delta*(G[i]-mean_G)))
		} else {
			g[i] = 0.0
		}
		if H[i] > 0 {
			h[i] = 1.0 / (1.0 + math.Exp(lamda*(H[i]-mean_H)))
		} else {
			h[i] = 1.0
		}
	}
	//
	//log.Println(g)
	q := rand.Float64()
	if q < Q0 {
		// Exploitation
		max_heuristic := -1.0
		for i, adj := range ant.Unvisited {
			heuristic := P[cur][adj] *
				math.Pow(g[i], beta) *
				math.Pow(h[i], gamma)
			if heuristic > max_heuristic {
				max_heuristic = heuristic
				nex = i
			}
		}
	} else {
		// Exploration
		sum_heuristic := float64(0)
		arr_heurictic := make([]float64, 0)
		for i, adj := range ant.Unvisited {
			sum_heuristic += P[cur][adj] *
				math.Pow(g[i], beta) *
				math.Pow(h[i], gamma)
			arr_heurictic = append(arr_heurictic, sum_heuristic)
		}
		random := rand.Float64()
		if sum_heuristic > 0 {
			for i := range ant.Unvisited {
				arr_heurictic[i] /= sum_heuristic
				if arr_heurictic[i] >= random {
					nex = i
					break
				}
			}
		}
	}
	return nex
}

func (ant *Ant) update_local() {
	l := len(ant.Path)
	for i := 1; i < l; i++ {
		cur, nex := ant.Path[i-1], ant.Path[i]
		P[cur][nex] = (1-omega)*P[cur][nex] +
			omega*ant.Delta[cur][nex]
	}
}
