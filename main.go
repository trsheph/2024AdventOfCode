package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"sort"
)

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Plugh(r rune) bool {
	return r=='\t' || r==' '
}

func parseTicker (textIn string) (leftValue int, rightValue int, err error) {
	i:=0
	colSlice := strings.FieldsFunc(textIn, Plugh)
	for _,colA := range colSlice {
		if len(colA) > 0 {
			tmpInt,err := strconv.Atoi(colA)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			if i==0 {
				leftValue=tmpInt
				i=i+1
			} else if i == 1 {
				rightValue=tmpInt
				i=i+1
			} else {
				err = fmt.Errorf("more values on line than expected")
				log.Fatal(err)
				panic(err)
			}
		}
	}
	if i != 2 && i != 0 {
		err = fmt.Errorf("fewer values on line than expected")
		log.Fatal(err)
		panic(err)
	}
	return
}

func secondCheck (vals []int) (isGood bool) {
	var err error
	isGood=false
	for i:=0; i<len(vals); i++ {
		var testVals []int
		if !(isGood) {
			for j,val := range vals {
				if i != j {
					testVals=append(testVals,val)
				}
			}
			if !(isGood) { 
				isGood,err = isGoodCheck(testVals)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return
}

func isGoodCheck (vals []int) (isGood bool, err error) {
	var valsSortUp,valsSortDown []int
	var goodFlag bool
	for _,val := range vals {
		valsSortUp = append(valsSortUp, val)
		valsSortDown = append(valsSortDown, val)
	}
	sort.Ints(valsSortUp)
	sort.Slice(valsSortDown, func(i,j int) bool {
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
		for i:=1; i<len(vals); i++ {
			if absInt(vals[i]-vals[i-1]) < 1 || absInt(vals[i]-vals[i-1]) > 3 {
				isGood=false
			}
		}
	} else {
		isGood=false
	}
	if len(vals) == 0 {
		isGood=false
	}
	return
}

func linesToSlice (textIn string) (vals []int, isGood bool, err error) {
	rowSlice := strings.FieldsFunc(textIn, Plugh)
	for _,rowA := range rowSlice {
		if len(rowA) > 0 {
			tmpInt,err := strconv.Atoi(rowA)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			vals=append(vals,tmpInt)
		}
	}
	isGood,err = isGoodCheck(vals)
	if err != nil {
		panic(err)
	}
	return
}

func dayOne (inFile string) {
	var rowSlice []string
	var leftSlice,rightSlice []int
	var leftValue,rightValue,sumOfDiffs,simScore int
	var fileString string
	rightCounts := make(map[int]int)
	dat, err := os.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fileString = string(dat)
	rowSlice = strings.Split(fileString, "\n")
	for _,rowA := range rowSlice {
		leftValue, rightValue, err = parseTicker(rowA)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		leftSlice=append(leftSlice, leftValue)
		rightSlice=append(rightSlice, rightValue)
	}
	sort.Ints(leftSlice)
	sort.Ints(rightSlice)
	if len(leftSlice) == len(rightSlice) {
		for i := range leftSlice {
			rightCounts[leftSlice[i]]=0
			sumOfDiffs = sumOfDiffs+absInt(rightSlice[i]-leftSlice[i])
		}
		for _,rVal := range rightSlice {
			rightCounts[rVal]=rightCounts[rVal]+1
		}
		for i := range leftSlice {
			simScore=simScore+leftSlice[i]*rightCounts[leftSlice[i]]
		}
	} else {
		err = fmt.Errorf("left and right columns have unequal number of entries")
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(sumOfDiffs)
	fmt.Println(simScore)
	return
}

func dayTwo (inFile string) {
	var rowSlice []string
	var rowVals []int
	var goodCount int
	var err error
	var rowBool,isAlmost bool
	var fileString string
	dat, err := os.ReadFile(inFile)
	if err != nil {
		panic(err)
	}
	fileString = string(dat)
	rowSlice = strings.Split(fileString, "\n")
	for _,rowA := range rowSlice {
		rowVals, rowBool, err = linesToSlice(rowA)
		if err != nil {
			panic(err)
		}
		if rowBool {
			goodCount = goodCount+1
		} else {
			isAlmost = secondCheck(rowVals)
			if isAlmost {
				goodCount = goodCount+1
			}
		}
	}
	fmt.Println(goodCount)
	return
}

func main () {
	var inFile string
	var err error
	inDay := 1
	if len(os.Args) == 1 {
		inFile = "tTtest.txt"
	} else if len(os.Args) == 2 {
		inFile = string(os.Args[1])
	} else if len(os.Args) == 3 {
		inDay,err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		inFile = string(os.Args[2])
	} else {
		err := fmt.Errorf("too many arguments")
		panic(err)
	}
	if inDay == 1 {
		dayOne(inFile)
	} else if inDay == 2 {
		dayTwo(inFile)
	}
}
