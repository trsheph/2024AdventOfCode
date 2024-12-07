package daysixb

import (
	"fmt"
	"os"
	"strings"
)

type Obst struct {
	x int
	y int
}

type UPath struct {
	x int
	y int
}

type UVector struct {
	x  int
	y  int
	nx int
	ny int
}

type GData struct {
	x         int
	y         int
	nx        int
	ny        int
	dx        int
	dy        int
	isUniq    bool
	isObst    bool
	isLoop    bool
	outBounds bool
}

func (g *GData) ObstCheck(obs []Obst) {
	g.isObst = false
	for _, ob := range obs {
		if (g.nx == ob.x) && (g.ny == ob.y) {
			g.isObst = true
		}
	}
}

func (g *GData) GTurn() {
	if g.isObst {
		// the guard is obstructed
		if g.dy == -1 {
			g.dx = 1
			g.dy = 0
		} else if g.dy == 1 {
			g.dx = -1
			g.dy = 0
		} else if g.dx == 1 {
			g.dx = 0
			g.dy = 1
		} else if g.dx == -1 {
			g.dx = 0
			g.dy = -1
		} else {
			fmt.Println("Something went wrong with Turn")
		}
		g.nx = g.x + g.dx
		g.ny = g.y + g.dy
	}
}

func (g *GData) GTakeStep() {
	g.x = g.nx
	g.y = g.ny
	g.nx = g.x + g.dx
	g.ny = g.y + g.dy
}

func AddUniq(g *GData, gsteps []UPath, gvecs []UVector) (gsn []UPath, gvs []UVector) {
	g.isUniq = true
	for _, gp := range gsteps {
		if g.x == gp.x && g.y == gp.y {
			g.isUniq = false
		}
	}
	if g.isUniq {
		var ugp UPath
		ugp.x = g.x
		ugp.y = g.y
		// fmt.Println("NewStep here:")
		// fmt.Println(ugp)
		gsteps = append(gsteps, ugp)
	}
	for _, gv := range gvecs {
		if ((g.x == gv.x) && (g.y == gv.y)) && ((g.nx == gv.nx) && (g.ny == gv.ny)) {
			g.isLoop = true
		}
	}
	if !(g.isLoop) {
		var ugv UVector
		ugv.x = g.x
		ugv.y = g.y
		ugv.nx = g.nx
		ugv.ny = g.ny
		gvecs = append(gvecs, ugv)
	}
	gsn = gsteps
	gvs = gvecs
	return
}

func checkBounds(g *GData, xmax int, ymax int) (outBounds bool) {
	if g.x >= xmax || g.y >= ymax {
		outBounds = true
	} else if g.x < 0 || g.y < 0 {
		outBounds = true
	} else {
		outBounds = false
	}
	return
}

func readDaySix(filename string) (fileRows []string, xmax int, ymax int) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileRows = strings.Split(fileString, "\n")
	ymax = len(fileRows) - 1
	xmax = len(fileRows[0]) - 1
	return
}

func initDaySix(fileRows []string) (obstSet []Obst, g GData, gPath []UPath, gVecs []UVector) {
	gPath = []UPath{}
	gVecs = []UVector{}
	for i, row := range fileRows {
		for j, col := range row {
			if string(col) == "#" {
				var tmpObst Obst
				tmpObst.x = j
				tmpObst.y = i
				obstSet = append(obstSet, tmpObst)
			} else if string(col) == "^" {
				g.x = j
				g.y = i
				g.dx = 0
				g.dy = -1
			} else if string(col) == "<" {
				g.x = j
				g.y = i
				g.dx = -1
				g.dy = 0
			} else if string(col) == ">" {
				g.x = j
				g.y = i
				g.dx = 1
				g.dy = 0
			} else if string(col) == "v" {
				g.x = j
				g.y = i
				g.dx = 0
				g.dy = 1
			}
		}
	}
	g.nx = g.x + g.dx
	g.ny = g.y + g.dy
	g.isUniq = true
	g.isLoop = false
	g.outBounds = false
	gPath, gVecs = AddUniq(&g, gPath, gVecs)
	return
}

func followPath(obsts []Obst, g *GData, gp []UPath, gv []UVector, xmax int, ymax int) (gpout []UPath, gvout []UVector) {
	keepGoing := true
	for keepGoing {
		g.ObstCheck(obsts)
		if g.isObst {
			g.GTurn()
		} else {
			g.GTakeStep()
		}
		gp, gv = AddUniq(g, gp, gv)
		g.outBounds = checkBounds(g, xmax, ymax)
		if g.isLoop || g.outBounds {
			break
		}
	}
	gpout = gp
	gvout = gv
	return
}

func ProcDaySix(filename string) {
	var allLoops int
	fileRows, xmax, ymax := readDaySix(filename)
	obsts, g, ugp, ugv := initDaySix(fileRows)
	ginit := g
	ugp, ugv = followPath(obsts, &g, ugp, ugv, xmax, ymax)
	ugpinit := ugp
	fmt.Println(len(ugpinit) - 1)
	for i := 1; i < len(ugpinit); i++ {
		var newObst Obst
		newObst.x = ugpinit[i].x
		newObst.y = ugpinit[i].y
		obsts = append(obsts, newObst)
		ugp = []UPath{}
		ugv = []UVector{}
		g = ginit
		ugp, ugv = AddUniq(&g, ugp, ugv)
		ugp, ugv = followPath(obsts, &g, ugp, ugv, xmax, ymax)
		if g.isLoop {
			allLoops = allLoops + 1
		}
		obsts = obsts[:len(obsts)-1]
	}
	fmt.Println(allLoops)
}
