package dayseven

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readDaySeven(filename string) (fileRows []string) {
	fileBytes, err := os.ReadFile(filename)
	for err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileRows = strings.Split(fileString, "\n")
	return
}

func textToEqs(fileRows []string) ([]int, map[int][]int) {
	xs := make(map[int][]int)
	var vals []int
	for i, row := range fileRows {
		if len(row) > 0 {
			tmpSStr := strings.Split(row, ": ")
			tmpVal, err := strconv.Atoi(string(tmpSStr[0]))
			if err != nil {
				panic(err)
			}
			vals = append(vals, tmpVal)
			tmpSRight := strings.Split(string(tmpSStr[1]), " ")
			for _, rightSideSplit := range tmpSRight {
				tmpX, err := strconv.Atoi(string(rightSideSplit))
				if err != nil {
					panic(err)
				}
				xs[i] = append(xs[i], tmpX)
			}
		}
	}
	return vals, xs
}

func removeDuplicates(slice []int) []int {
	encountered := map[int]bool{}
	result := []int{}
	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

func powInt(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func concat(x int, y int) int {
	outVal, err := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
	if err != nil {
		panic(err)
	}
	return outVal
}

func ProcDaySeven(filename string, ptBS string) {
	var ptB bool
	if ptBS == "true" || ptBS == "t" {
		ptB = true
	}
	var totalVal int
	var goodRows []int
	fileRows := readDaySeven(filename)
	lside, rside := textToEqs(fileRows)
	for i := range lside {
		var allOps []int
		allOps = append(allOps, rside[i][0])
		var topOps int
		for j := 1; j < len(rside[i]); j++ {
			limMe := len(allOps)
			for k := 0; k < limMe; k++ {
				sOp := allOps[k] + rside[i][j]
				mOp := allOps[k] * rside[i][j]
				allOps = append(allOps, sOp)
				allOps = append(allOps, mOp)
				if ptB {
					cOp := concat(allOps[k], rside[i][j])
					allOps = append(allOps, cOp)
				}
			}
		}
		topOps = len(allOps) - powInt(2, len(rside[i])-1)
		if ptB {
			topOps = len(allOps) - powInt(3, len(rside[i])-1)
		}
		allOps = allOps[topOps:]
		for j := 0; j < len(allOps); j++ {
			if lside[i] == allOps[j] {
				goodRows = append(goodRows, i)
			}
		}
	}
	goodRows = removeDuplicates(goodRows)
	for _, goodRow := range goodRows {
		totalVal = totalVal + lside[goodRow]
	}
	fmt.Println(totalVal)
}
