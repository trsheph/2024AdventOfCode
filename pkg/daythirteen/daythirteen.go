package daythirteen

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ax int64
	ay int64
	bx int64
	by int64
	px int64
	py int64
}

func calcAns(newGame Game) (a, b int64, checkGood bool) {
	var af, bf float64
	axf := float64(newGame.ax)
	ayf := float64(newGame.ay)
	bxf := float64(newGame.bx)
	byf := float64(newGame.by)
	pxf := float64(newGame.px)
	pyf := float64(newGame.py)
	bf = (pyf - ayf/axf*pxf) / (-ayf/axf*bxf + byf)
	af = (pxf - bxf*bf) / axf
	a = int64(math.Round(af))
	b = int64(math.Round(bf))
	checkGood = (newGame.py == newGame.ay*a+newGame.by*b)
	if checkGood {
		checkGood = (newGame.px == newGame.ax*a+newGame.bx*b)
	}
	return
}

func readDayThirteen(filename string) (games []Game) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileSlice := strings.Split(fileString, "\n")
	var newGame Game
	for _, row := range fileSlice {
		if len(row) > 0 {
			if string(strings.Split(row, ":")[0]) == "Button A" {
				newGame.ax, err = strconv.ParseInt(string(strings.Split(string(strings.Split(row, ",")[0]), "+")[1]), 10, 64)
				if err != nil {
					panic(err)
				}
				newGame.ay, err = strconv.ParseInt(string(strings.Split(row, "+")[2]), 10, 64)
				if err != nil {
					panic(err)
				}
			} else if string(strings.Split(row, ":")[0]) == "Button B" {
				newGame.bx, err = strconv.ParseInt(string(strings.Split(string(strings.Split(row, ",")[0]), "+")[1]), 10, 64)
				if err != nil {
					panic(err)
				}
				newGame.by, err = strconv.ParseInt(string(strings.Split(row, "+")[2]), 10, 64)
				if err != nil {
					panic(err)
				}
			} else if string(strings.Split(row, ":")[0]) == "Prize" {
				newGame.px, err = strconv.ParseInt(string(strings.Split(string(strings.Split(row, ",")[0]), "=")[1]), 10, 64)
				if err != nil {
					panic(err)
				}
				newGame.py, err = strconv.ParseInt(string(strings.Split(row, "=")[2]), 10, 64)
				if err != nil {
					panic(err)
				}
			}
		} else {
			games = append(games, newGame)
			newGame = Game{}
		}
	}
	return
}

func ProcDayThirteen(filename string) {
	var a, b, atot, btot int64
	var isGood bool
	games := readDayThirteen(filename)
	for _, newGame := range games {
		a, b, isGood = calcAns(newGame)
		if isGood {
			atot = atot + a
			btot = btot + b
		}
	}
	outAns := 3*atot + btot
	fmt.Println(outAns)
	for i := 0; i < len(games); i++ {
		games[i].px = games[i].px + 10000000000000
		games[i].py = games[i].py + 10000000000000
	}
	atot = 0
	btot = 0
	for _, newGame := range games {
		a, b, isGood = calcAns(newGame)
		if isGood {
			atot = atot + a
			btot = btot + b
		}
	}
	outAns = 3*atot + btot
	fmt.Println(outAns)
}
