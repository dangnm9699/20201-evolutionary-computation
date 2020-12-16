package tsp

import "math/rand"

func swapMutation(gene []int, mutationProb float64) {
	if mutationProb < MutationThreshold {
		muPos1 := rand.Intn(NumCities)
		muPos2 := rand.Intn(NumCities)
		gene[muPos1], gene[muPos2] = gene[muPos2], gene[muPos1]
	}
}