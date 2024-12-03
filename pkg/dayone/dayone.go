package dayone

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"sort"

	"github.com/trsheph/2024AdventOfCode/pkg/tickersort"
)

func parseTicker (textIn string) (leftValue int, rightValue int, err error) {
	i:=0
	colSlice := strings.FieldsFunc(textIn, tickersort.TabOrSpace)
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

func ProcDayOne (inFile string) {
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
			sumOfDiffs = sumOfDiffs+tickersort.AbsInt(rightSlice[i]-leftSlice[i])
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
