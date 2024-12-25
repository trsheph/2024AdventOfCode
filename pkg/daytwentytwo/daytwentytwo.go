package daytwentytwo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Secret struct {
	Prev, Curr     uint64
	A, B, C, D     int16
	Da, Db, Dc, Dd int16
}

func (s *Secret) loadSecret(str string) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	} else {
		s.Curr = num
	}
}

func (s *Secret) fillVals() {
	var err error
	var holdInt int64
	s.A = s.B
	s.Da = s.Db
	s.B = s.C
	s.Db = s.Dc
	s.C = s.D
	s.Dc = s.Dd
	holdStr := strconv.FormatUint(s.Curr, 10)
	holdOnes := string(holdStr[len(holdStr)-1:])
	holdInt, err = strconv.ParseInt(holdOnes, 10, 8)
	if err != nil {
		panic(err)
	}
	s.D = int16(holdInt)
	s.Dd = s.D - s.C
}

func (s *Secret) getNext(multiA, modBaseA, divA, multiB uint64) {
	s.Prev = s.Curr
	stepA := (s.Curr ^ (s.Curr * multiA)) % modBaseA
	stepB := (stepA ^ (stepA / divA)) % modBaseA
	s.Curr = (stepB ^ (stepB * multiB)) % modBaseA
}

func readDayTwentyTwo(filename string) (secrets []Secret) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileByte)
	fileRows := strings.Split(fileString, "\n")
	for _, row := range fileRows {
		if len(row) > 0 {
			var secr Secret
			secr.loadSecret(row)
			secrets = append(secrets, secr)
		}
	}
	return
}

func ProcDayTwentyTwo(filename string) {
	numSteps := 2000
	var tensor [][][][]int16
	total := uint64(0)
	irows := 19
	jrows := 19
	krows := 19
	lrows := 19
	tensor = make([][][][]int16, irows)
	for i := range tensor {
		tensor[i] = make([][][]int16, jrows)
		for j := range tensor[i] {
			tensor[i][j] = make([][]int16, krows)
			for k := range tensor[i][j] {
				tensor[i][j][k] = make([]int16, lrows)
			}
		}
	}
	secrets := readDayTwentyTwo(filename)
	for i := 0; i < 3; i++ {
		for j := 0; j < len(secrets); j++ {
			secrets[j].getNext(64, 16777216, 32, 2048)
			secrets[j].fillVals()
		}
	}
	for j := 0; j < len(secrets); j++ {
		cacheTensor := make([][][][]int16, irows)
		for i := range cacheTensor {
			cacheTensor[i] = make([][][]int16, jrows)
			for j := range cacheTensor[i] {
				cacheTensor[i][j] = make([][]int16, krows)
				for k := range cacheTensor[i][j] {
					cacheTensor[i][j][k] = make([]int16, lrows)
				}
			}
		}
		for i := 3; i < numSteps; i++ {
			secrets[j].getNext(64, 16777216, 32, 2048)
			secrets[j].fillVals()
			if cacheTensor[secrets[j].Da+9][secrets[j].Db+9][secrets[j].Dc+9][secrets[j].Dd+9] == 0 {
				tmpVal := secrets[j].D + 1
				cacheTensor[secrets[j].Da+9][secrets[j].Db+9][secrets[j].Dc+9][secrets[j].Dd+9] += tmpVal
			}
		}
		for a := 0; a < irows; a++ {
			for b := 0; b < jrows; b++ {
				for c := 0; c < krows; c++ {
					for d := 0; d < lrows; d++ {
						if cacheTensor[a][b][c][d] > 0 {
							tensor[a][b][c][d] = tensor[a][b][c][d] + cacheTensor[a][b][c][d] - 1
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(secrets); i++ {
		total = total + secrets[i].Curr
	}
	fmt.Println(total)
	var maxVal int16
	maxVal = int16(0)
	for i := 0; i < irows; i++ {
		for j := 0; j < jrows; j++ {
			for k := 0; k < krows; k++ {
				for l := 0; l < lrows; l++ {
					if tensor[i][j][k][l] > maxVal {
						maxVal = tensor[i][j][k][l]
					}
				}
			}
		}
	}
	fmt.Println(maxVal)
}
