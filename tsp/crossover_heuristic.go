package tsp

import (
	"ec/utils"
	"fmt"
	"math/rand"
	"os"
)

func heuristicCrossover(p1, p2 []int) (c []int) {
	if len(p1) != len(p2) {
		panic("[PAN] Not the same length")
	}
	randCity := rand.Intn(NumCities) + 1
	indexInP1, err := utils.GetIndex(p1, randCity)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	indexInP2, err := utils.GetIndex(p2, randCity)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c = make([]int, NumCities)
	prev := 0
	c[prev] = randCity
	rightCircumvolve(p1, indexInP1, 0)
	rightCircumvolve(p2, indexInP2, 0)
	for i := 1; i < NumCities; i++ {
		prev += 1
		if DistanceMatrix[c[prev-1]][p1[i]] < DistanceMatrix[c[prev-1]][p2[i]] {
			c[prev] = p1[i]
			index, err := utils.GetIndex(p2, p1[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			rightCircumvolve(p2, index, i)
		} else {
			c[prev] = p2[i]
			index, err := utils.GetIndex(p1, p2[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			rightCircumvolve(p1, index, i)
		}
	}
	return
}

func rightCircumvolve(p []int, src int, dst int) {
	if !isValid(src) || !isValid(dst) {
		panic("[PAN] Invalid parameters")
	}
	prev := make([]int, len(p[dst:src]))
	next := make([]int, len(p[src:]))
	//copy(prev, p[dst:dst+len(p)-len(p[src:])])
	copy(prev, p[dst:src])
	copy(next, p[src:])
	for i := 0; i < len(next); i++ {
		p[dst+i] = next[i]
	}
	for i := 0; i < len(prev); i++ {
		p[dst+len(next)+i] = prev[i]
	}
}
