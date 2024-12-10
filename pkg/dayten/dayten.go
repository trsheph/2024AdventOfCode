package dayten

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const HighValue int = 9
const LowValue int = 0

type height struct {
	x int
	y int
	h int
}

type hiker struct {
	x        int
	y        int
	h        int
	finished bool
	pathID   string
}

func takeSteps(mh hiker, outMatrix [][]int, pathway map[string][]hiker) (hikerWave []hiker) {
	lastPathID := mh.pathID
	if mh.y-1 >= 0 {
		if outMatrix[mh.y-1][mh.x]-mh.h == 1 {
			var newHiker hiker
			newHiker.x = mh.x
			newHiker.y = mh.y - 1
			newHiker.h = outMatrix[mh.y-1][mh.x]
			newHiker.pathID = uuid.New().String()
			if newHiker.h == HighValue {
				newHiker.finished = true
			}
			hikerWave = append(hikerWave, newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], pathway[lastPathID]...)
		}
	}
	if mh.y+1 < len(outMatrix) {
		if outMatrix[mh.y+1][mh.x]-mh.h == 1 {
			var newHiker hiker
			newHiker.x = mh.x
			newHiker.y = mh.y + 1
			newHiker.h = outMatrix[mh.y+1][mh.x]
			newHiker.pathID = uuid.New().String()
			if newHiker.h == HighValue {
				newHiker.finished = true
			}
			hikerWave = append(hikerWave, newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], pathway[lastPathID]...)
		}
	}
	if mh.x-1 >= 0 {
		if outMatrix[mh.y][mh.x-1]-mh.h == 1 {
			var newHiker hiker
			newHiker.x = mh.x - 1
			newHiker.y = mh.y
			newHiker.h = outMatrix[mh.y][mh.x-1]
			newHiker.pathID = uuid.New().String()
			if newHiker.h == HighValue {
				newHiker.finished = true
			}
			hikerWave = append(hikerWave, newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], pathway[lastPathID]...)
		}
	}
	if mh.x+1 < len(outMatrix[0]) {
		if outMatrix[mh.y][mh.x+1]-mh.h == 1 {
			var newHiker hiker
			newHiker.x = mh.x + 1
			newHiker.y = mh.y
			newHiker.h = outMatrix[mh.y][mh.x+1]
			newHiker.pathID = uuid.New().String()
			if newHiker.h == HighValue {
				newHiker.finished = true
			}
			hikerWave = append(hikerWave, newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], newHiker)
			pathway[newHiker.pathID] = append(pathway[newHiker.pathID], pathway[lastPathID]...)
		}
	}
	delete(pathway, lastPathID)
	return
}

func findTrailheads(topomap []height) (trailstarts []height) {
	for _, topo := range topomap {
		if topo.h == LowValue {
			trailstarts = append(trailstarts, topo)
		}
	}
	return
}

func readDayTen(filename string) (outMatrix [][]int, topomap []height) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileString, "\n")
	fileRows = fileRows[:len(fileRows)-1]
	for i, row := range fileRows {
		var rowVals []int
		for j, val := range row {
			var localtopo height
			valInt, err := strconv.Atoi(string(val))
			if err != nil {
				panic(err)
			}
			localtopo.h = valInt
			localtopo.x = j
			localtopo.y = i
			topomap = append(topomap, localtopo)
			rowVals = append(rowVals, valInt)
		}
		outMatrix = append(outMatrix, rowVals)
	}
	return
}

func initDayTen(trailheads []height, pathway map[string][]hiker) (hikers []hiker) {
	for _, trailH := range trailheads {
		var tmpHiker hiker
		tmpHiker.x = trailH.x
		tmpHiker.y = trailH.y
		tmpHiker.h = LowValue
		tmpHiker.finished = false
		tmpHiker.pathID = uuid.New().String()
		hikers = append(hikers, tmpHiker)
		pathway[tmpHiker.pathID] = hikers
	}
	return
}

func ProcDayTen(filename string) {
	var hikers, newHikers, currentHikers []hiker
	var currFlag, contFlag bool
	pathway := make(map[string][]hiker)
	pathFinish := make(map[string][]hiker)
	outMatrix, topomap := readDayTen(filename)
	trailheads := findTrailheads(topomap)
	hikers = initDayTen(trailheads, pathway)
	for _, initH := range hikers {
		contFlag = true
		currentHikers = []hiker{}
		currentHikers = append(currentHikers, initH)
		for contFlag {
			currFlag = false
			newHikers = []hiker{}
			for _, mh := range currentHikers {
				if len(pathway) == 0 {
					currFlag = true
				} else {
					if mh.finished {
						pathFinish[mh.pathID] = pathway[mh.pathID]
						delete(pathway, mh.pathID)
					} else {
						tmpHikers := takeSteps(mh, outMatrix, pathway)
						newHikers = append(newHikers, tmpHikers...)
					}
				}
			}
			if len(newHikers) == 0 {
				currFlag = true
			}
			currentHikers = newHikers
			contFlag = !(currFlag)
		}
	}
	var pathStarts []hiker
	var pathEnds []hiker
	for pathid := range pathFinish {
		pathStart := pathFinish[pathid][len(pathFinish[pathid])-1]
		pathEnd := pathFinish[pathid][0]
		pathStarts = append(pathStarts, pathStart)
		pathEnds = append(pathEnds, pathEnd)
	}
	var countUPath int
	var addFlag bool
	for i := range pathStarts {
		addFlag = true
		for j := i + 1; j < len(pathStarts); j++ {
			if ((pathStarts[i].x == pathStarts[j].x) && (pathStarts[i].y == pathStarts[j].y)) && ((pathEnds[i].x == pathEnds[j].x) && (pathEnds[i].y == pathEnds[j].y)) {
				addFlag = false
			}
		}
		if addFlag {
			countUPath++
		}
	}
	fmt.Println(countUPath)
	fmt.Println(len(pathFinish))
}
