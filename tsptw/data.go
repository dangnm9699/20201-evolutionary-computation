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
	//
	RC201_1 BSPB = "rc_201.1.txt"
	RC201_2 BSPB = "rc_201.2.txt"
	RC201_3 BSPB = "rc_201.3.txt"
	RC201_4 BSPB = "rc_201.4.txt"
	//
	RC202_1 BSPB = "rc_202.1.txt"
	RC202_2 BSPB = "rc_202.2.txt"
	RC202_3 BSPB = "rc_202.3.txt"
	RC202_4 BSPB = "rc_202.4.txt"
	//
	RC203_1 BSPB = "rc_203.1.txt"
	RC203_2 BSPB = "rc_203.2.txt"
	RC203_3 BSPB = "rc_203.3.txt"
	RC203_4 BSPB = "rc_203.4.txt"
	//
	RC204_1 BSPB = "rc_204.1.txt"
	RC204_2 BSPB = "rc_204.2.txt"
	RC204_3 BSPB = "rc_204.3.txt"
	//
	RC205_1 BSPB = "rc_205.1.txt"
	RC205_2 BSPB = "rc_205.2.txt"
	RC205_3 BSPB = "rc_205.3.txt"
	RC205_4 BSPB = "rc_205.4.txt"
	//
	RC206_1 BSPB = "rc_206.1.txt"
	RC206_2 BSPB = "rc_206.2.txt"
	RC206_3 BSPB = "rc_206.3.txt"
	RC206_4 BSPB = "rc_206.4.txt"
	//
	RC207_1 BSPB = "rc_207.1.txt"
	RC207_2 BSPB = "rc_207.2.txt"
	RC207_3 BSPB = "rc_207.3.txt"
	RC207_4 BSPB = "rc_207.4.txt"
	//
	RC208_1 BSPB = "rc_208.1.txt"
	RC208_2 BSPB = "rc_208.2.txt"
	RC208_3 BSPB = "rc_208.3.txt"
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
