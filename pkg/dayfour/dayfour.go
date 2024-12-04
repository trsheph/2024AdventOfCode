package dayfour

import (
	"fmt"
	"os"
	"strings"
)

type Xcoord struct {
	x int
	y int
}

type Mcoord struct {
	x int
	y int
}

type Acoord struct {
	x int
	y int
}

type Scoord struct {
	x int
	y int
}

type Coord struct {
	x int
	y int
}

type Dist struct {
	Xx int
	Xy int
	Dxxm int
	Dyxm int
	Dxma int
	Dyma int
	Dxas int
	Dyas int
}

var Xs []Xcoord
var Ms []Mcoord
var As []Acoord
var Ss []Scoord
var Ns []Coord
var dists []Dist

func calcDists () (goodCnt int) {
	for _,xf := range Xs {
		for _,mf := range Ms {
			dx := mf.x-xf.x
			dy := mf.y-xf.y
			if (dx > -2 && dx < 2) && (dy > -2 && dy < 2) {
				var xdist Dist
				xdist.Xx = xf.x
				xdist.Xy = xf.y
				xdist.Dxxm = mf.x-xf.x
				xdist.Dyxm = mf.y-xf.y
				for _,af := range As {
					if (mf.x+xdist.Dxxm == af.x) && (mf.y+xdist.Dyxm == af.y) {
						for _,sf := range Ss {
							if (af.x+xdist.Dxxm == sf.x) && (af.y+xdist.Dyxm == sf.y) {
								goodCnt = goodCnt+1
							}
						}
					}
				}
			}
		}
	}
	return
}

func calcMASs () (goodCnt int) {
	for _,af := range As {
		goodA := 0
		for _,mf := range Ms {
			dx := af.x - mf.x
			dy := af.y - mf.y
			if (dx == -1 || dx == 1) && (dy == -1 || dy == 1) {
				for _,sf := range Ss {
					ddx := sf.x - af.x
					ddy := sf.y - af.y
					if dx == ddx && dy == ddy {
						goodA = goodA + 1
						if goodA == 2 {
							goodCnt = goodCnt+1
						}
					}
				}
			}
		}
	}
	return
}

func loadMatrix (filename string) (outRows [][]rune, goodCnt int) {
	fileBytes,err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(fileBytes)
	fileSlice := strings.Split(fileContents, "\n")
	colsNum := len(fileSlice[0]) + 4
	rowsNum := len(fileSlice) + 3
	outRows = make([][]rune, rowsNum)
	for i := range outRows {
		outRows[i] = make([]rune, colsNum)
	}
	for i := 0; i < colsNum; i++ {
		for j := range outRows[0] {
			outRows[0][j] = '.'
			outRows[1][j] = '.'
			outRows[len(fileSlice)+1][j] = '.'
			outRows[len(fileSlice)+2][j] = '.'
		}
		var tmp Coord
		tmp.x=0
		tmp.y=i
		Ns = append(Ns, tmp)
		tmp.x=1
		Ns = append(Ns, tmp)
		tmp.x=len(fileSlice)+1
		Ns = append(Ns, tmp)
		tmp.x=len(fileSlice)+2
		Ns = append(Ns, tmp)
	}
	for i := 0; i < len(fileSlice)+2; i++ {
		outRows[i][0] = '.'
		outRows[i][1] = '.'
		outRows[i][colsNum-1] = '.'
		outRows[i][colsNum-2] = '.'
		var tmp Coord
		tmp.x=i
		tmp.y=0
		Ns = append(Ns, tmp)
		tmp.y=1
		Ns = append(Ns, tmp)
		tmp.y=colsNum
		Ns = append(Ns, tmp)
		tmp.y=colsNum-1
		Ns = append(Ns, tmp)
	}
	for i := 0; i < len(fileSlice)-1; i++ {
		for j := 0; j < len(fileSlice[0]); j++ {
			if string(fileSlice[i][j])=="X" {
				outRows[i+2][j+2]='X'
				var Xtmp Xcoord
				Xtmp.x=i+2
				Xtmp.y=j+2
				Xs=append(Xs, Xtmp)
			} else if string(fileSlice[i][j])=="M" {
				outRows[i+2][j+2]='M'
				var tmp Mcoord
				tmp.x=i+2
				tmp.y=j+2
				Ms=append(Ms, tmp)
			} else if string(fileSlice[i][j])=="A" {
				outRows[i+2][j+2]='A'
				var tmp Acoord
				tmp.x=i+2
				tmp.y=j+2
				As=append(As, tmp)
			} else if string(fileSlice[i][j])=="S" {
				outRows[i+2][j+2]='S'
				var tmp Scoord
				tmp.x=i+2
				tmp.y=j+2
				Ss=append(Ss, tmp)
			} else {
				outRows[i+2][j+2]='.'
				var tmp Coord
				tmp.x=i+2
				tmp.y=j+2
				Ns=append(Ns, tmp)
			}
		}
	}
	goodCnt = calcDists()
	goodMASs := calcMASs()
	fmt.Println("Good XMASs:")
	fmt.Println(goodCnt)
	fmt.Println("Good X-MASs:")
	fmt.Println(goodMASs)
	return
}

func ProcDayFour (filename string) {
	loadMatrix(filename)
}
