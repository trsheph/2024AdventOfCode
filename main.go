package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/trsheph/2024AdventOfCode/pkg/dayone"
	"github.com/trsheph/2024AdventOfCode/pkg/daythree"
	"github.com/trsheph/2024AdventOfCode/pkg/daytwo"
)

func main() {
	var inFile string
	var err error
	inDay := 1
	var procDos string
	if len(os.Args) == 1 {
		inFile = "DayOneTest.txt"
	} else if len(os.Args) == 2 {
		inFile = string(os.Args[1])
	} else if len(os.Args) == 3 {
		inDay, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		inFile = string(os.Args[2])
	} else if len(os.Args) == 4 {
		inDay, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		inFile = string(os.Args[2])
		procDos = string(os.Args[3])
	} else {
		err := fmt.Errorf("too many arguments")
		panic(err)
	}
	if inDay == 1 {
		dayone.ProcDayOne(inFile)
	} else if inDay == 2 {
		daytwo.ProcDayTwo(inFile)
	} else if inDay == 3 {
		daythree.ProcDayThree(inFile, procDos)
	}
}
