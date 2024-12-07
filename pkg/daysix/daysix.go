package daysix

import (
	"fmt"
	"os"
	"strings"
)

type Obstacle struct {
	Xpos  int
	Ypos  int
	Rules []RulePlate
}

type RulePlate struct {
	Xpos  int
	Ypos  int
	NewDx int
	NewDy int
	InDx  int
	InDy  int
}

type GuardData struct {
	Xpos   int
	Ypos   int
	Nx     int
	Ny     int
	Dx     int
	Dy     int
	Unique bool
}

var Obstacles []Obstacle
var Guard GuardData
var GuardPath []GuardData

/*
	func countUniquePositions(totalSteps int) (totalCnt int) {
		totalCnt = totalSteps
		for i := range GuardPath {
			for j := i + 1; j < len(GuardPath); j++ {
				if (GuardPath[i].Xpos == GuardPath[j].Xpos) && (GuardPath[i].Ypos == GuardPath[j].Ypos) {
					totalCnt = totalCnt - 1
				}
			}
		}
		return
	}
*/
func genNewMap(xob int, yob int) (localObs []Obstacle) {
	var skipFlag bool
	skipFlag = false
	if (xob == Guard.Xpos) && (yob == Guard.Ypos) {
		skipFlag = true
	}
	for _, tmpobs := range Obstacles {
		localObs = append(localObs, tmpobs)
		if tmpobs.Xpos == xob && tmpobs.Ypos == yob {
			skipFlag = true
		}
	}
	if !(skipFlag) {
		var tmpRule RulePlate
		var tmpRules []RulePlate
		tmpRule.Xpos = xob + 1
		tmpRule.Ypos = yob
		tmpRule.NewDx = 0
		tmpRule.NewDy = -1
		tmpRule.InDx = -1
		tmpRule.InDy = 0
		tmpRules = append(tmpRules, tmpRule)
		tmpRule.Xpos = xob - 1
		tmpRule.Ypos = yob
		tmpRule.NewDx = 0
		tmpRule.NewDy = 1
		tmpRule.InDx = 1
		tmpRule.InDy = 0
		tmpRules = append(tmpRules, tmpRule)
		tmpRule.Xpos = xob
		tmpRule.Ypos = yob + 1
		tmpRule.NewDx = 1
		tmpRule.NewDy = 0
		tmpRule.InDx = 0
		tmpRule.InDy = -1
		tmpRules = append(tmpRules, tmpRule)
		tmpRule.Xpos = xob
		tmpRule.Ypos = yob - 1
		tmpRule.NewDx = -1
		tmpRule.NewDy = 0
		tmpRule.InDx = 0
		tmpRule.InDy = 1
		tmpRules = append(tmpRules, tmpRule)
		var tmpOb Obstacle
		tmpOb.Xpos = xob
		tmpOb.Ypos = yob
		tmpOb.Rules = tmpRules
		localObs = append(localObs, tmpOb)
	} else {
		localObs = []Obstacle{}
	}
	return
}

func takeStep(xmax int, ymax int, localObs []Obstacle) (stepCount int, err string) {
	isObstructed := true
	Guard.Xpos = Guard.Nx
	Guard.Ypos = Guard.Ny
	Guard.Nx = Guard.Xpos + Guard.Dx
	Guard.Ny = Guard.Ypos + Guard.Dy
	for _, guardPos := range GuardPath {
		if Guard.Xpos == guardPos.Xpos && Guard.Ypos == guardPos.Ypos {
			Guard.Unique = false
			if Guard.Nx == guardPos.Nx && Guard.Ny == guardPos.Ny {
				err = "guard is looping"
				// fmt.Println("Guard is looping:")
				return
			}
		}
	}
	if Guard.Unique {
		stepCount = stepCount + 1
	}
	GuardPath = append(GuardPath, Guard)
	for isObstructed {
		for _, obstruction := range localObs {
			if (Guard.Nx == obstruction.Xpos) && (Guard.Ny == obstruction.Ypos) {
				for _, plate := range obstruction.Rules {
					if (Guard.Xpos == plate.Xpos && Guard.Ypos == plate.Ypos) && (Guard.Dx == plate.InDx && Guard.Dy == plate.InDy) {
						Guard.Dx = plate.NewDx
						Guard.Dy = plate.NewDy
						Guard.Nx = Guard.Xpos + Guard.Dx
						Guard.Ny = Guard.Ypos + Guard.Dy
						for _, guardPos := range GuardPath {
							if Guard.Xpos == guardPos.Xpos && Guard.Ypos == guardPos.Ypos {
								if Guard.Nx == guardPos.Nx && Guard.Ny == guardPos.Ny {
									err = "guard is looping"
									// fmt.Println("Guard is looping")
									return
								}
							}
						}
						GuardPath = append(GuardPath, Guard)
					}
				}
			} else {
				isObstructed = false
			}
		}
	}
	if (Guard.Nx <= 0 || Guard.Nx >= xmax) || (Guard.Ny <= 0 || Guard.Ny >= ymax) {
		err = "guard is out"
		return
	}
	return
}

func ReadDaySix(filename string) (charMatrix [][]rune) {
	GuardPath = []GuardData{}
	Obstacles = []Obstacle{}
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileSlice := strings.Split(fileString, "\n")
	rowsNum := len(fileSlice) + 1
	colsNum := len(fileSlice[0]) + 2
	outMatrix := make([][]rune, rowsNum)
	for i := range outMatrix {
		outMatrix[i] = make([]rune, colsNum)
	}
	for i := range outMatrix {
		for j := range outMatrix[i] {
			if i == 0 || i == len(outMatrix)-1 {
				outMatrix[i][j] = '.'
			} else if j == 0 || j == len(outMatrix[i])-1 {
				outMatrix[i][j] = '.'
			}
		}
	}
	var sumAll int
	for i, row := range fileSlice {
		for j, rcVal := range row {
			sumAll = sumAll + 1
			if string(rcVal) == "." {
				outMatrix[i+1][j+1] = '.'
			} else if string(rcVal) == "#" {
				outMatrix[i+1][j+1] = '#'
				var tmpRule RulePlate
				var tmpRules []RulePlate
				tmpRule.Xpos = j + 1 + 1
				tmpRule.Ypos = i + 1 + 0
				tmpRule.NewDx = 0
				tmpRule.NewDy = -1
				tmpRule.InDx = -1
				tmpRule.InDy = 0
				tmpRules = append(tmpRules, tmpRule)
				tmpRule.Xpos = j + 1 - 1
				tmpRule.Ypos = i + 1 + 0
				tmpRule.NewDx = 0
				tmpRule.NewDy = 1
				tmpRule.InDx = 1
				tmpRule.InDy = 0
				tmpRules = append(tmpRules, tmpRule)
				tmpRule.Xpos = j + 1 + 0
				tmpRule.Ypos = i + 1 + 1
				tmpRule.NewDx = 1
				tmpRule.NewDy = 0
				tmpRule.InDx = 0
				tmpRule.InDy = -1
				tmpRules = append(tmpRules, tmpRule)
				tmpRule.Xpos = j + 1 + 0
				tmpRule.Ypos = i + 1 - 1
				tmpRule.NewDx = -1
				tmpRule.NewDy = 0
				tmpRule.InDx = 0
				tmpRule.InDy = 1
				tmpRules = append(tmpRules, tmpRule)
				var tmpOb Obstacle
				tmpOb.Xpos = j + 1
				tmpOb.Ypos = i + 1
				tmpOb.Rules = tmpRules
				Obstacles = append(Obstacles, tmpOb)
			} else if string(rcVal) == "^" {
				outMatrix[i+1][j+1] = '^'
				var tmpGuard GuardData
				tmpGuard.Xpos = j + 1
				tmpGuard.Ypos = i + 1
				tmpGuard.Dx = 0
				tmpGuard.Dy = -1
				tmpGuard.Nx = tmpGuard.Xpos + tmpGuard.Dx
				tmpGuard.Ny = tmpGuard.Ypos + tmpGuard.Dy
				Guard = tmpGuard
				GuardPath = append(GuardPath, Guard)
			} else if string(rcVal) == "<" {
				outMatrix[i+1][j+1] = '<'
				var tmpGuard GuardData
				tmpGuard.Xpos = j + 1
				tmpGuard.Ypos = i + 1
				tmpGuard.Dx = -1
				tmpGuard.Dy = 0
				tmpGuard.Nx = tmpGuard.Xpos + tmpGuard.Dx
				tmpGuard.Ny = tmpGuard.Ypos + tmpGuard.Dy
				Guard = tmpGuard
				GuardPath = append(GuardPath, Guard)
			} else if string(rcVal) == ">" {
				outMatrix[i+1][j+1] = '>'
				var tmpGuard GuardData
				tmpGuard.Xpos = j + 1
				tmpGuard.Ypos = i + 1
				tmpGuard.Dx = 1
				tmpGuard.Dy = 0
				tmpGuard.Nx = tmpGuard.Xpos + tmpGuard.Dx
				tmpGuard.Ny = tmpGuard.Ypos + tmpGuard.Dy
				Guard = tmpGuard
				GuardPath = append(GuardPath, Guard)
			} else if string(rcVal) == "v" {
				outMatrix[i+1][j+1] = 'v'
				var tmpGuard GuardData
				tmpGuard.Xpos = j + 1
				tmpGuard.Ypos = i + 1
				tmpGuard.Dx = 0
				tmpGuard.Dy = 1
				tmpGuard.Nx = tmpGuard.Xpos + tmpGuard.Dx
				tmpGuard.Ny = tmpGuard.Ypos + tmpGuard.Dy
				Guard = tmpGuard
				GuardPath = append(GuardPath, Guard)
			}
		}
	}
	fmt.Println(sumAll)
	charMatrix = outMatrix
	return
}

func gInZone(charMatrix [][]rune, localObs []Obstacle) (totalSteps int, steperr string) {
	totalSteps = 2
	guardInZone := true
	for guardInZone {
		var stepCount int
		stepCount, steperr = takeStep(len(charMatrix[0])-1, len(charMatrix)-1, localObs)
		if steperr != "" {
			if steperr == "guard is out" {
				guardInZone = false
				break
			} else if steperr == "guard is looping" {
				guardInZone = false
				break
			} else {
				err := fmt.Errorf(steperr)
				panic(err)
			}
		}
		totalSteps = totalSteps + stepCount
	}
	return
}

func ProcDaySix(filename string) {
	charMatrix := ReadDaySix(filename)
	var obstaclesInit []Obstacle
	var guardPathInit []GuardData
	var guardInit GuardData
	var totalObstructions int
	for _, obs := range Obstacles {
		obstaclesInit = append(obstaclesInit, obs)
		totalObstructions = totalObstructions + 1
	}
	for _, rp := range GuardPath {
		guardPathInit = append(guardPathInit, rp)
	}
	guardInit = Guard
	totalSteps, _ := gInZone(charMatrix, Obstacles)
	// totalUniquePositions := countUniquePositions(totalSteps)
	fmt.Println(totalSteps)
	var totalLoopPositions, totalOutPositions int
	for i := 1; i < len(charMatrix)-1; i++ {
		for j := 1; j < len(charMatrix[i])-1; j++ {
			Obstacles = []Obstacle{}
			GuardPath = []GuardData{}
			Guard = GuardData{}
			for _, obs := range obstaclesInit {
				Obstacles = append(Obstacles, obs)
			}
			for _, rp := range guardPathInit {
				GuardPath = append(GuardPath, rp)
			}
			Guard = guardInit
			localObs := genNewMap(j, i)
			if len(localObs) != 0 {
				_, steperr := gInZone(charMatrix, localObs)
				if steperr == "guard is looping" {
					// fmt.Println(localObs[len(localObs)-1])
					totalLoopPositions = totalLoopPositions + 1
				} else if steperr == "guard is out" {
					totalOutPositions = totalOutPositions + 1
				}
			}
		}
	}
	fmt.Println(totalLoopPositions)
	fmt.Println(totalOutPositions)
	fmt.Println(totalObstructions)
	fmt.Println(totalLoopPositions + totalOutPositions + totalObstructions + 1)
}
