package utils

import "errors"

func GetIndex(p []int, city int) (int, error) {
	for i := range p {
		if p[i] == city {
			return i, nil
		}
	}
	return 0, errors.New("[ERR] City is not in genotype")
}
