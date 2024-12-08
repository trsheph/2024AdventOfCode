package dayeight

import (
	"fmt"
	"strings"
	"os"
)

type AntPos struct {
	X int
	Y int
}

type AntLine struct {
	M int
	B int
}

func readDayEight (filename string) (outMatrix [][]rune) {
	fileByte, err := os.ReadFile(filename)
	if err != nil { panic(err) }
	fileString := string(fileByte)
	fileRows := strings.Split(fileString, "\n")
	a := make([][]rune, len(fileRows)-1)
	for i := range a {
		a[i] = make([]rune, len(fileRows[0]))
	}
	for i,row := range fileRows {
		for j, chara := range row {
			a[i][j]=rune(chara)
		}
	}
	outMatrix=a
	return
}

func findAllAnt (a rune, inMat [][]rune) (allAntXY []AntPos) {
	for y := range inMat {
		for x := range inMat[y] {
			if a == inMat[y][x] {
				var aP AntPos
				aP.X = x
				aP.Y = y
				allAntXY = append(allAntXY, aP)
			}
		}
	}
	return
}

func ProcDayEight (filename string, ptBS string) {
	var ptB bool
	if ptBS == "t" || ptBS == "true" {
		ptB = true
	}
	inMat := readDayEight(filename)
	var totalAns int
	// var allLines []AntLine
	ymax := len(inMat)
	xmax := len(inMat[0])
	anMat := make([][]rune, len(inMat))
	for i := range anMat {
		anMat[i] = make([]rune, len(inMat[0]))
	}
	// fmt.Println(inMat)
	// fmt.Println(anMat)
	allChars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allSymb := []rune(allChars)
	for _,node := range allSymb {
		tmpXY := findAllAnt(node, inMat)
		if len(tmpXY) != 0 {
			for i:=0; i < len(tmpXY); i++ {
				for j:=i+1; j < len(tmpXY); j++ {
					if !(ptB) {
						var posA, posB AntPos
						dx := tmpXY[j].X-tmpXY[i].X
						dy := tmpXY[j].Y-tmpXY[i].Y
						posA.X = tmpXY[i].X - dx
						posA.Y = tmpXY[i].Y - dy
						posB.X = tmpXY[j].X + dx
						posB.Y = tmpXY[j].Y + dy
						if (posA.X >= 0 && posA.X < xmax) && (posA.Y >= 0 && posA.Y < ymax) { 
							anMat[posA.Y][posA.X] = '#'
						}
						if (posB.X >= 0 && posB.X < xmax) && (posB.Y >= 0 && posB.Y < ymax) {
							anMat[posB.Y][posB.X] = '#'
						}
					} else {
						dx := tmpXY[j].X-tmpXY[i].X
						dy := tmpXY[j].Y-tmpXY[i].Y
						var posA AntPos
						posA.X = tmpXY[i].X
						posA.Y = tmpXY[i].Y
						for (((posA.X >= 0) && (posA.X < xmax)) && ((posA.Y >= 0) && (posA.Y < ymax))) {
							anMat[posA.Y][posA.X] = '#'
							posA.X = posA.X - dx
							posA.Y = posA.Y - dy
						}
						posA.X = tmpXY[i].X
						posA.Y = tmpXY[i].Y
						for (((posA.X >= 0) && (posA.X < xmax)) && ((posA.Y >= 0) && (posA.Y < ymax))) {
							anMat[posA.Y][posA.X] = '#'
							posA.X = posA.X + dx
							posA.Y = posA.Y + dy
						}
					}
				}
			}
		}
	}
	for i := range anMat {
		for j := range anMat[i] {
			if anMat[i][j]=='#' {
				totalAns = totalAns+1
			}
		}
	}
	fmt.Println(totalAns)
}
