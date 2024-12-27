package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/trsheph/2024AdventOfCode/pkg/dayeight"
	"github.com/trsheph/2024AdventOfCode/pkg/dayeleven"
	"github.com/trsheph/2024AdventOfCode/pkg/dayfive"
	"github.com/trsheph/2024AdventOfCode/pkg/dayfour"
	"github.com/trsheph/2024AdventOfCode/pkg/dayfourteen"
	"github.com/trsheph/2024AdventOfCode/pkg/daynine"
	"github.com/trsheph/2024AdventOfCode/pkg/dayone"
	"github.com/trsheph/2024AdventOfCode/pkg/dayseven"
	"github.com/trsheph/2024AdventOfCode/pkg/daysix"
	"github.com/trsheph/2024AdventOfCode/pkg/dayten"
	"github.com/trsheph/2024AdventOfCode/pkg/daythirteen"
	"github.com/trsheph/2024AdventOfCode/pkg/daythree"
	"github.com/trsheph/2024AdventOfCode/pkg/daytwentyfive"
	"github.com/trsheph/2024AdventOfCode/pkg/daytwentyfour"
	"github.com/trsheph/2024AdventOfCode/pkg/daytwentythree"
	"github.com/trsheph/2024AdventOfCode/pkg/daytwentytwo"
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
	} else if inDay == 4 {
		dayfour.ProcDayFour(inFile)
	} else if inDay == 5 {
		dayfive.ProcDayFive(inFile)
	} else if inDay == 6 {
		daysix.ProcDaySix(inFile)
	} else if inDay == 7 {
		dayseven.ProcDaySeven(inFile, procDos)
	} else if inDay == 8 {
		dayeight.ProcDayEight(inFile, procDos)
	} else if inDay == 9 {
		daynine.ProcDayNine(inFile)
	} else if inDay == 10 {
		dayten.ProcDayTen(inFile)
	} else if inDay == 11 {
		dayeleven.ProcDayEleven(inFile)
	} else if inDay == 13 {
		daythirteen.ProcDayThirteen(inFile)
	} else if inDay == 14 {
		dayfourteen.ProcDayFourteen(inFile)
	} else if inDay == 22 {
		daytwentytwo.ProcDayTwentyTwo(inFile)
	} else if inDay == 23 {
		daytwentythree.ProcDayTwentyThree(inFile)
	} else if inDay == 24 {
		daytwentyfour.ProcDayTwentyFour(inFile)
	} else if inDay == 25 {
		daytwentyfive.ProcDayTwentyFive(inFile)
	}
}
