package dayfourteen

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RData struct {
	x  int
	y  int
	nx int
	ny int
	dx int
	dy int
}

var (
	FieldLength = 101
	FieldHeight = 103
	StepNumber  = 10000
)

func average(inVals []int) (ave float64) {
	var total int
	for _, val := range inVals {
		total = total + val
	}
	ave = float64(total) / float64(len(inVals))
	return
}

func avefloat(inVals []float64) (ave float64) {
	var total float64
	for _, val := range inVals {
		total = total + val
	}
	ave = total / float64(len(inVals))
	return
}

func variancefloat(inVals []float64) (outVariance float64) {
	var totalDifSq float64
	aveVals := avefloat(inVals)
	for _, val := range inVals {
		diffSq := (val - aveVals) * (val - aveVals)
		totalDifSq = totalDifSq + diffSq
	}
	outVariance = totalDifSq / float64(len(inVals)-1)
	return
}

func stdevfloat(inVals []float64) (outStdev float64) {
	outStdev = math.Sqrt(variancefloat(inVals))
	return
}

func printPicture(xS, yS []int) (outString []string) {
	for i := 0; i < FieldHeight; i++ {
		var outRow string
		for j := 0; j < FieldLength; j++ {
			outRow = outRow + "."
			for k := 0; k < len(xS); k++ {
				if xS[k] == j && yS[k] == i {
					outRow = outRow[:len(outRow)-1] + "#"
				}
			}
		}
		outRow = outRow + "\n"
		outString = append(outString, outRow)
	}
	return
}

func variance(inVals []int) (outVariance float64) {
	var totalDifSq float64
	aveVals := average(inVals)
	for _, val := range inVals {
		diffSq := (float64(val) - aveVals) * (float64(val) - aveVals)
		totalDifSq = totalDifSq + diffSq
	}
	outVariance = totalDifSq / float64(len(inVals)-1)
	return
}

func robotJustXY(robots []RData) (xS, yS []int) {
	for _, newRobot := range robots {
		xS = append(xS, newRobot.x)
		yS = append(yS, newRobot.y)
	}
	return
}

func takeStep(robots []RData) (outRobots []RData) {
	for i := 0; i < len(robots); i++ {
		if robots[i].nx < 0 {
			robots[i].nx = FieldLength + (robots[i].nx % FieldLength)
		}
		if robots[i].nx >= FieldLength {
			robots[i].nx = ((robots[i].nx) % FieldLength)
		}
		if robots[i].ny < 0 {
			robots[i].ny = FieldHeight + ((robots[i].ny) % FieldHeight)
		}
		if robots[i].ny >= FieldHeight {
			robots[i].ny = ((robots[i].ny) % FieldHeight)
		}
	}
	for i := 0; i < len(robots); i++ {
		robots[i].x = robots[i].nx
		robots[i].y = robots[i].ny
		robots[i].nx = robots[i].x + robots[i].dx
		robots[i].ny = robots[i].y + robots[i].dy
	}
	outRobots = robots
	return
}

func readDayFourteen(filename string) (robots []RData) {
	var newRobot RData
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileString, "\n")
	fileRows = fileRows[:len(fileRows)-1]
	for _, row := range fileRows {
		vecs := strings.Split(row, " ")
		pVec := strings.Split(vecs[0], "=")[1]
		vVec := strings.Split(vecs[1], "=")[1]
		newRobot.x, err = strconv.Atoi(strings.Split(pVec, ",")[0])
		if err != nil {
			panic(err)
		}
		newRobot.y, err = strconv.Atoi(strings.Split(pVec, ",")[1])
		if err != nil {
			panic(err)
		}
		newRobot.dx, err = strconv.Atoi(strings.Split(vVec, ",")[0])
		if err != nil {
			panic(err)
		}
		newRobot.dy, err = strconv.Atoi(strings.Split(vVec, ",")[1])
		if err != nil {
			panic(err)
		}
		newRobot.nx = newRobot.x + newRobot.dx
		newRobot.ny = newRobot.y + newRobot.dy
		robots = append(robots, newRobot)
	}
	return
}

func ProcDayFourteen(filename string) {
	var tl, tr, bl, br int
	var disp []float64
	robots := readDayFourteen(filename)
	for i := 0; i < StepNumber; i++ {
		robots = takeStep(robots)
		xS, yS := robotJustXY(robots)
		// fmt.Println(printPicture(xS, yS))
		xVari := variance(xS)
		yVari := variance(yS)
		robotGM := 1 / math.Sqrt(math.Pow(xVari, 2)+(math.Pow(yVari, 2)))
		disp = append(disp, robotGM)
	}
	framesAve := avefloat(disp)
	framesStdev := stdevfloat(disp)
	var susFrames []int
	for i, frame := range disp {
		if math.Abs((frame-framesAve))/framesStdev > 10 {
			fmt.Println("Suspicious Frame: ", i+1)
			susFrames = append(susFrames, i+1)
		}
	}
	for _, newRobot := range robots {
		if newRobot.x < ((FieldLength-1)/2) && newRobot.y < ((FieldHeight-1)/2) {
			tl = tl + 1
		} else if newRobot.x > ((FieldLength-1)/2) && newRobot.y < ((FieldHeight-1)/2) {
			tr = tr + 1
		} else if newRobot.x < ((FieldLength-1)/2) && newRobot.y > ((FieldHeight-1)/2) {
			bl = bl + 1
		} else if newRobot.x > ((FieldLength-1)/2) && newRobot.y > ((FieldHeight-1)/2) {
			br = br + 1
		}
	}
	fmt.Println("Part 1: ", tl*tr*bl*br)
	robots = readDayFourteen(filename)
	for i := 0; i <= StepNumber; i++ {
		robots = takeStep(robots)
		for _, susFrame := range susFrames {
			if i == susFrame-1 {
				xS, yS := robotJustXY(robots)
				fmt.Println(printPicture(xS, yS))
			}
		}
	}
}
