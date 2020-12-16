package tsp

import "math/rand"

func genPop() (population [][]int) {
	population = make([][]int, PopSize)
	for i := 0; i < PopSize; i++ {
		population[i] = genInd()
	}
	return
}

func genInd() (individual []int) {
	individual = rand.Perm(NumCities)
	for i := range individual {
		individual[i] += 1
	}
	return
}