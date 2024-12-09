// Package implements the 2024 - Day 6 Advent of Code Solution
package daysix

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// Obst contains the x and y position of an obstacle
type Obst struct {
	x int
	y int
}

// UPath contains the x and y positions of the unique 
// path explored by the guard
type UPath struct {
	x int
	y int
}

// UVector contains the x and y positions and the 
// next x and next y positions of the path explored 
// by the guard
type UVector struct {
	x  int
	y  int
	nx int
	ny int
}

// GData contains the guard data. x and y positions, 
// next x, next y, dx, dy directions, isUniq step, 
// isObst is the guard obstructed on the next step, 
// isLoop is the guard looping
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

// Function addUniq identifies if the guard is on a unique tile and if the guard is on a 
// unique tile and vector. If the latter, then the guard is looping, so g.isLoop returns
// true. If the guard is on a unique tile, then the tile is added to the UPath slice
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

// Function checkBounds returns a boolean true if the guard is out of the map, 
// or else returns false if the guard is in the map
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

// Function readDaySix reads in the file at filename and produces the slice of rows
// of the file.
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

// Function initDaySix initializes the obstacle position and guard position
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
	gPath, gVecs = addUniq(&g, gPath, gVecs)
	return
}

// Function followPath uses methods to move the guard along the map, given the obstacle positions obsts
func followPath(obsts []Obst, g *GData, gp []UPath, gv []UVector, xmax int, ymax int) (gpout []UPath, gvout []UVector) {
	keepGoing := true
	for keepGoing {
		g.ObstCheck(obsts)
		if g.isObst {
			g.GTurn()
		} else {
			g.GTakeStep()
			gp, gv = addUniq(g, gp, gv)
		}
		g.outBounds = checkBounds(g, xmax, ymax)
		if g.isLoop || g.outBounds {
			break // TO FIX - setting keepGoing=false did not work here.
		}
	}
	gpout = gp
	gvout = gv
	return
}

func ProcDaySix(filename string) {
	var allLoops int // allLoops is the total count of times when the object causes the guard to loop
	var ugpinit []UPath // ugp is the unique guard path, the only path worth placing obstacles in part 2
	maxThreads := 1  // Wacky behavior on multithreading from passing pointers to functions maybe. TO FIX
	testUGPCh := make(chan UPath) // making a channel for the go routine
	var wg sync.WaitGroup // make the wg WaitGroup
	fileRows, xmax, ymax := readDaySix(filename) // read in the file into a slice of rows
	obsts, g, ugp, ugv := initDaySix(fileRows) // read in the guard and obstacles
	ginit := g // set an initial guard position for part 2
	ugp, ugv = followPath(obsts, &g, ugp, ugv, xmax, ymax) // Let the guard follow the path
	for _,ugstep := range ugp {
		if ugstep.x >= 0 && ugstep.y >= 0 {
			ugpinit = append(ugpinit, ugstep) // Removing the final step that the guard takes out of bounds
		}
	}
	fmt.Println(len(ugpinit))
	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for testUGP := range testUGPCh {
				var newObst Obst
				obstsA := obsts // set initial obstructions
				newObst.x = testUGP.x // for each unique plate, try an obstacle
				newObst.y = testUGP.y
				obstsA = append(obstsA, newObst)
				ugpN := []UPath{} // set blank path
				ugvN := []UVector{} // set blank vector
				gN := ginit // set initialized guard
				ugpN, ugvN = addUniq(&gN, ugpN, ugvN)
				ugpN, ugvN = followPath(obstsA, &gN, ugpN, ugvN, xmax, ymax)
				if gN.isLoop {
					allLoops = allLoops + 1
				}
			}
		}()
	}
	for j := 1; j < len(ugpinit); j++ {
		testUGPCh <- ugpinit[j]
	}
	close(testUGPCh)
	wg.Wait()
	fmt.Println(allLoops)
}
