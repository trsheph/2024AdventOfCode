// Package implements the 2024 - Day 15 Advent of Code Solution
package dayfifteen

import (
	"fmt"
	"os"
	"strings"
)

// Obst contains the x and y position of an obstacle
type Obst struct {
	x int
	y int
}

// UPath contains the x and y positions of the unique
// path explored by the player
type UPath struct {
	x int
	y int
}

// UVector contains the x and y positions and the
// next x and next y positions of the path explored
// by the player
type UVector struct {
	x  int
	y  int
	nx int
	ny int
}

// GData contains the player data. x and y positions,
// next x, next y, dx, dy directions, isUniq step,
// isObst is the player obstructed on the next step,
// isLoop is the player looping
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

// Method applied to guard to check if they are
// about to hit an obstacle
func (g *GData) ObstCheck(obs []Obst) {
	g.isObst = false
	for _, ob := range obs {
		if (g.nx == ob.x) && (g.ny == ob.y) {
			g.isObst = true
		}
	}
}

// Method applied to guard to turn
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

// Method applied to guard to take a step
func (g *GData) GTakeStep() {
	g.x = g.nx
	g.y = g.ny
	g.nx = g.x + g.dx
	g.ny = g.y + g.dy
}

// Function addUniq identifies if the player is on a unique tile and if the player is on a
// unique tile and vector. If not unique, then the player on a path alread explored.
// If the guard is on a unique tile, then the tile is added to the UPath slice
func addUniq(g *GData, gsteps []UPath, gvecs []UVector) (gsn []UPath, gvs []UVector) {
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

// Function readDaySix reads in the file at filename and produces the slice of rows
// of the file.
func readDayFifteen(filename string) (fileRows []string, xmax int, ymax int) {
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

// Function initDayFifteen initializes the obstacle positions and player position
func initDayFifteen(fileRows []string) (obstSet []Obst, g GData, gPath []UPath, gVecs []UVector) {
	gPath = []UPath{}
	gVecs = []UVector{}
	for i, row := range fileRows {
		for j, col := range row {
			if string(col) == "#" {
				var tmpObst Obst
				tmpObst.x = j
				tmpObst.y = i
				obstSet = append(obstSet, tmpObst)
			} else if string(col) == "S" {
				g.x = j
				g.y = i
				g.dx = 0
				g.dy = -1
			} else if string(col) == "E" {
				g.x = j
				g.y = i
				g.dx = -1
				g.dy = 0
			}
		}
	}
	g.nx = g.x + g.dx
	g.ny = g.y + g.dy
	g.isUniq = true
	g.isLoop = false
	g.outBounds = false
	// gPath, gVecs = addUniq(&g, gPath, gVecs)
	return
}
/*
func ProcDayFifteen(filename string) {
	fileRows, xmax, ymax := readDayFifteen(filename)
	obsts, p, ppath, pvecs := initDayFifteen(fileRows)
}
*/