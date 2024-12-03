package daytwo

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/trsheph/2024AdventOfCode/pkg/tickersort"
)

func secondCheck(vals []int) (isGood bool) {
	var err error
	isGood = false
	for i := 0; i < len(vals); i++ {
		var testVals []int
		if !(isGood) {
			for j, val := range vals {
				if i != j {
					testVals = append(testVals, val)
				}
			}
			if !(isGood) {
				isGood, err = isGoodCheck(testVals)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return
}

func isGoodCheck(vals []int) (isGood bool, err error) {
	var valsSortUp, valsSortDown []int
	var goodFlag bool
	for _, val := range vals {
		valsSortUp = append(valsSortUp, val)
		valsSortDown = append(valsSortDown, val)
	}
	sort.Ints(valsSortUp)
	sort.Slice(valsSortDown, func(i, j int) bool {
		return valsSortDown[i] > valsSortDown[j]
	})
	if len(valsSortUp) == len(vals) {
		goodFlag = true
		for i := range vals {
			if vals[i] != valsSortUp[i] {
				goodFlag = false
			}
		}
	}
	if len(valsSortDown) == len(vals) && !(goodFlag) {
		goodFlag = true
		for i := range vals {
			if vals[i] != valsSortDown[i] {
				goodFlag = false
			}
		}
	}
	if goodFlag {
		isGood = true
		for i := 1; i < len(vals); i++ {
			if tickersort.AbsInt(vals[i]-vals[i-1]) < 1 || tickersort.AbsInt(vals[i]-vals[i-1]) > 3 {
				isGood = false
			}
		}
	} else {
		isGood = false
	}
	if len(vals) == 0 {
		isGood = false
	}
	return
}

func linesToSlice(textIn string) (vals []int, isGood bool, err error) {
	rowSlice := strings.FieldsFunc(textIn, tickersort.TabOrSpace)
	for _, rowA := range rowSlice {
		if len(rowA) > 0 {
			tmpInt, err := strconv.Atoi(rowA)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			vals = append(vals, tmpInt)
		}
	}
	isGood, err = isGoodCheck(vals)
	if err != nil {
		panic(err)
	}
	return
}

func ProcDayTwo(inFile string) {
	var rowSlice []string
	var rowVals []int
	var goodCount int
	var err error
	var rowBool, isAlmost bool
	var fileString string
	dat, err := os.ReadFile(inFile)
	if err != nil {
		panic(err)
	}
	fileString = string(dat)
	rowSlice = strings.Split(fileString, "\n")
	for _, rowA := range rowSlice {
		rowVals, rowBool, err = linesToSlice(rowA)
		if err != nil {
			panic(err)
		}
		if rowBool {
			goodCount = goodCount + 1
		} else {
			isAlmost = secondCheck(rowVals)
			if isAlmost {
				goodCount = goodCount + 1
			}
		}
	}
	fmt.Println(goodCount)
	return
}
