package dayfive

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type PageOrder struct {
	IndBefore int
	IndAfter int
	Before int
	After int
}

// var rules []string
var RuleSets []PageOrder

func readRules (filename string) (goodSum int, badSum int) {
	fileBytes,err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileString, "\n")
	rulesFlag := true
	for _,row := range fileRows {
		if string(row) == "" {
			rulesFlag = false
		} else {
			if rulesFlag {
				twoPages := strings.Split(string(row), "|")
				if len(twoPages) != 2 {
					errA := fmt.Errorf("something went wrong with more than 2 pages on a rule")
					panic(errA)
				}
				pageA,err := strconv.Atoi(twoPages[0])
				if err != nil {
					panic(err)
				}
				pageB,err := strconv.Atoi(twoPages[1])
				if err != nil {
					panic(err)
				}
				var ruleset PageOrder
				ruleset.Before = pageA
				ruleset.After = pageB
				RuleSets = append(RuleSets, ruleset)
			} else {
				goodMid,badMid := procPages(row)
				goodSum = goodSum + goodMid
				badSum = badSum + badMid
			}
		}
	}
	return
}

func orderPages(pageSlice []int) []int {
	var pagesBad bool
	var badBefore, badAfter int
	pagesBad = true
	newPageSlice := pageSlice
	for pagesBad==true {
		var pageRules []PageOrder
		pagesBad=false
		pageRules = genPageRules(newPageSlice)
		pagesBad,badBefore,badAfter = arePagesBad(pageRules)
		if pagesBad {
			var orderedPages []int
			for i := range newPageSlice {
				if i == badBefore {
					orderedPages = append(orderedPages, newPageSlice[badAfter])
				} else if i == badAfter {
					orderedPages = append(orderedPages, newPageSlice[badBefore])
				} else {
					orderedPages = append(orderedPages, newPageSlice[i])
				}
			}
			newPageSlice = orderedPages
		}
	}
	return newPageSlice
}

func calcMid(pageSlice []int) (midVal int) {
	midVal = pageSlice[((len(pageSlice)-1)/2)]
	return
}

func genPageRules(pageSlice []int) (pageRules []PageOrder) {
	for i := 0; i < len(pageSlice); i++ {
		pageA := pageSlice[i]
		for j := i+1; j < len(pageSlice); j++ {
			pageB := pageSlice[j]
			var pageRule PageOrder
			pageRule.Before = pageA
			pageRule.After = pageB
			pageRule.IndBefore = i
			pageRule.IndAfter = j
			pageRules = append(pageRules, pageRule)
		}
	}
	return
}

func arePagesBad(pageRules []PageOrder) (pagesBad bool, badBefore int, badAfter int) {
	var isRuleFound bool
	pagesBad=false
	for _,pR := range pageRules {
		isRuleFound = false
		for _,rule := range RuleSets {
			if !(isRuleFound) || pagesBad {
				if (pR.Before == rule.Before) && (pR.After == rule.After) {
					isRuleFound = true
				} else if (pR.Before == rule.After) && (pR.After == rule.Before) {
					pagesBad = true
					badBefore = pR.IndBefore
					badAfter = pR.IndAfter
				}
			}
		}
	}
	return
}

func procPages(pages string) (goodMid int, badMid int) {
	goodMid=0
	badMid=0
	var pagesBad bool
	var pageRules []PageOrder
	var pageSlice []int
	pageSliceStr := strings.Split(pages, ",")
	for _,pgSS := range pageSliceStr {
		pgVal,err := strconv.Atoi(pgSS)
		if err != nil {
			panic(err)
		}
		pageSlice = append(pageSlice, pgVal)
	}
	pageRules = genPageRules(pageSlice)
	pagesBad,_,_=arePagesBad(pageRules)
	if !(pagesBad) {
		goodMid=calcMid(pageSlice)
	} else {
		orderedPageSlice := orderPages(pageSlice)
		badMid = calcMid(orderedPageSlice)
	}
	return
}

func ProcDayFive (filename string) {
	goodSum, badSum := readRules(filename)
	fmt.Println("Good sum")
	fmt.Println(goodSum)
	fmt.Println("Reordered group sum")
	fmt.Println(badSum)
}
