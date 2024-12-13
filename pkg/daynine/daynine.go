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

func equalSlice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func compressZeros(inMatrix [][]int64) (newMatrix [][]int64) {
	var tmpRow []int64
	var cacheMatrix [][]int64
	backInMatrix := inMatrix
	cacheMatrix = [][]int64{{}}
	for i := 1; i < len(backInMatrix); i++ {
		if len(backInMatrix[i]) > 0 {
			if len(backInMatrix[i-1]) > 0 {
				if int64(backInMatrix[i-1][0]) == int64(0) && int64(backInMatrix[i][0]) == int64(0) {
					tmpRow = backInMatrix[i-1]
					tmpRow = append(tmpRow, backInMatrix[i]...)
					cacheMatrix = backInMatrix[:i-1]
					cacheMatrix = append(cacheMatrix, tmpRow)
					if len(backInMatrix[i:]) > 1 {
						cacheMatrix = append(cacheMatrix, backInMatrix[i+1:]...)
					}
					backInMatrix = cacheMatrix
					i = 1
				}
			}
		}
	}
	if len(cacheMatrix) > 1 {
		backInMatrix = cacheMatrix
	}
	for i := 0; i < len(backInMatrix); i++ {
		if len(backInMatrix[i]) > 0 {
			newMatrix = append(newMatrix, backInMatrix[i])
		}
	}
	return
}

func procPartTwo(dTwoSlice [][]int64) (dTwoSliceB [][]int64) {
	var lenV, lenR int
	dTwoSliceB = dTwoSlice // make copy of the 2D slice
	backupDTslice := [][]int64{{}}
	keepGoing := true
	appendBack := true
	sameArray := true
	// var newInvQ [][]int64
	tmpDTwo := [][]int64{{}}         // set a new 2D slice that is empty
	invertedQ := invQueue(dTwoSlice) // invertedQueue of additions to empty registers
	for keepGoing {
		for _, vals := range invertedQ {
			lenV = len(vals)
			backupDTslice = dTwoSliceB
			tmpDTwo = [][]int64{{}}
			appendBack = true
			for _, regi := range dTwoSliceB {
				lenR = len(regi)
				if lenR >= lenV && lenR > 0 {
					if (regi[0] == int64(0)) && appendBack { // this block is for when a register is full of 0s and large enough to fit the mem segment
						tmpDTwo = append(tmpDTwo, vals)
						tmpDTwo = append(tmpDTwo, regi[lenV:])
						// tmpDTwo = append(tmpDTwo, dTwoSliceB[j+1:]...)
						appendBack = false
					} else if equalSlice(regi, vals) { // if it rus into the current memory segment
						if appendBack == false { // if it runs into it after moving it, set the segment to 0s
							var tmpNewRegi []int64
							for i := 0; i < len(regi); i++ {
								tmpNewRegi = append(tmpNewRegi, int64(0))
							}
							tmpDTwo = append(tmpDTwo, tmpNewRegi)
						} else { // appendBack is true still, which means that it found itself before finding a register
							tmpDTwo = append(tmpDTwo, regi)
							appendBack = false
						}
					} else {
						tmpDTwo = append(tmpDTwo, regi)
					}
				} else {
					tmpDTwo = append(tmpDTwo, regi)
				}
			}
			tmpDTwo = tmpDTwo[1:]
			tmpDTwo = compressZeros(tmpDTwo)
			dTwoSliceB = tmpDTwo
		}
		// invertedQ = newInvQ
		sameArray = true
		for i, valsA := range backupDTslice {
			if !(equalSlice(valsA, dTwoSliceB[i])) {
				sameArray = false
			}
		}
		if sameArray {
			keepGoing = false
		}
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
	dTwoSlice, firstSlice := decompressTwo(diskSlice)
	outSlice := procPartTwo(dTwoSlice)
	finalMatrix := [][]int64{{}}
	finalMatrix[0] = firstSlice
	finalMatrix = append(finalMatrix, outSlice...)
	var finalSlice []int64
	for _, row := range finalMatrix {
		finalSlice = append(finalSlice, row...)
	}
	outSum = checkSum(finalSlice)
	fmt.Println(outSum)
}
