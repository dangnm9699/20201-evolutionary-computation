package tsptw

import (
	"fmt"
	"os"
)

const (
	BaseSolomonPotvinBengio = "tsptw/data/SolomonPotvinBengio/"
)

type BSPB string

var (
	RC201_1 BSPB = "rc_201.1.txt"
	RC201_2 BSPB = "rc_201.2.txt"
	RC201_4 BSPB = "rc_201.4.txt"
	RC202_2 BSPB = "rc_202.2.txt"
	RC203_1 BSPB = "rc_203.1.txt"
	RC203_4 BSPB = "rc_203.4.txt"
	RC205_1 BSPB = "rc_205.1.txt"
	RC206_1 BSPB = "rc_206.1.txt"
	RC207_4 BSPB = "rc_207.4.txt"
)

func getData(file BSPB) {
	f, _ := os.Open(string(BaseSolomonPotvinBengio + file))
	_, _ = fmt.Fscanln(f, &NC)
	DM = make([][]float64, NC)
	for i := range DM {
		DM[i] = make([]float64, NC)
		for j := range DM[i] {
			_, _ = fmt.Fscan(f, &DM[i][j])
		}
	}
	TW = make([][2]float64, NC)
	for i := range TW {
		_, _ = fmt.Fscanln(f, &TW[i][0], &TW[i][1])
	}
}
