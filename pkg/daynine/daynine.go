package daynine

import (
	"fmt"
	"os"
	"strconv"
)

func readDayNine(filename string) (diskString string) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	diskString = string(fileByte)
	return
}

func initDayNine(diskString string) (diskSlice []int64) {
	for _, chara := range diskString {
		if string(chara) != "\n" {
			tmpIn, err := strconv.Atoi(string(chara))
			if err != nil {
				panic(err)
			}
			diskSlice = append(diskSlice, int64(tmpIn))
		}
	}
	return
}

func decompressOne(diskSlice []int64) (mainSlice []int64) {
	fileID := int64(0)
	for i, dOut := range diskSlice {
		if i%2 == 0 {
			// number of fileIDs
			for j := 0; j < int(dOut); j++ {
				mainSlice = append(mainSlice, fileID)
			}
			fileID = fileID + 1
		} else {
			// number of empty spaces
			for j := 0; j < int(dOut); j++ {
				mainSlice = append(mainSlice, -1)
			}
		}
	}
	return
}

func decompressTwo(diskSlice []int64) (mainSlice [][]int64, preSlice []int64) {
	fileID := int64(0)
	for i, dOut := range diskSlice {
		var tmpSlice []int64
		if i%2 == 0 && i == 0 {
			for j := 0; j < int(dOut); j++ {
				preSlice = append(preSlice, fileID)
			}
			fileID = fileID + 1
		} else if i%2 == 0 && i > 0 {
			for j := 0; j < int(dOut); j++ {
				tmpSlice = append(tmpSlice, fileID)
			}
			fileID = fileID + 1
		} else {
			for j := 0; j < int(dOut); j++ {
				tmpSlice = append(tmpSlice, 0)
			}
		}
		mainSlice = append(mainSlice, tmpSlice)
	}
	return
}

func procPartOne(mainSlice []int64) (outSlice []int64) {
	lenMS := len(mainSlice)
	numPop := 1
	numPut := 1
	for _, val := range mainSlice {
		if val == -1 {
			if mainSlice[lenMS-numPop] == -1 {
				for mainSlice[lenMS-numPop] == -1 {
					numPop = numPop + 1
					numPut = numPut + 1
				}
				outSlice = append(outSlice, mainSlice[lenMS-numPop])
				numPop = numPop + 1
			} else {
				outSlice = append(outSlice, mainSlice[lenMS-numPop])
				numPop = numPop + 1
			}
		} else {
			outSlice = append(outSlice, val)
		}
	}
	outSlice = outSlice[:lenMS-numPop+numPut]
	return
}

func invQueue(dTwoSlice [][]int64) (tOutSlice [][]int64) {
	lenMS := len(dTwoSlice)
	for invInd := 0; invInd < lenMS-1; invInd++ {
		if invInd%2 == 0 {
			popSlice := dTwoSlice[lenMS-1-invInd]
			tOutSlice = append(tOutSlice, popSlice)
		}
	}
	return
}

func procPartTwo(dTwoSlice [][]int64) (dTwoSliceB [][]int64) {
	dTwoSliceB = dTwoSlice
	tmpDTwo := [][]int64{{}}
	shuffleQueue := invQueue(dTwoSlice)
	fmt.Println(shuffleQueue)
	for _, vals := range shuffleQueue {
		lenV := len(vals)
		for j, regi := range dTwoSliceB {
			lenR := len(regi)
			if lenR >= lenV && lenR > 0 {
				if regi[0] == int64(0) {
					tmpDTwo = append(tmpDTwo, vals)
					tmpDTwo = append(tmpDTwo, regi[lenV:])
					for k := j + 1; k < len(dTwoSliceB); k++ {
						tmpDTwo = append(tmpDTwo, dTwoSliceB[k])
					}
				} else {
					tmpDTwo = append(tmpDTwo, regi)
				}
			} else {
				tmpDTwo = append(tmpDTwo, regi)
			}
		}
		fmt.Println(tmpDTwo)
		dTwoSliceB = tmpDTwo
		tmpDTwo = [][]int64{{}}
	}
	return
}

func checkSum(outSlice []int64) (outSum int64) {
	for i, val := range outSlice {
		outSum = outSum + int64(i)*val
	}
	return
}

func ProcDayNine(filename string) {
	diskSlice := initDayNine(readDayNine(filename))
	outSum := checkSum(procPartOne(decompressOne(diskSlice)))
	fmt.Println(outSum)
	dTwoSlice, firstslice := decompressTwo(diskSlice)
	fmt.Println(dTwoSlice)
	fmt.Println(firstslice)
	outSlice := procPartTwo(dTwoSlice)
	fmt.Println(outSlice)
	// outSum = checkSum(outSlice)
}
