package daythree

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	x int
	y int
}

func (prod Product) Mult() int {
	return prod.x * prod.y
}

func readFile(filename string) (outTxt string, err error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(fileBytes)
	fileRows := strings.Split(fileContents, "\n")
	for _, rowA := range fileRows {
		outTxt = outTxt + rowA
	}
	return
}

func procMatchDos(inTxt string) (outMuls []Product, err error) {
	var editedLine string
	dontBreaks := strings.Split(inTxt, "don't()")
	editedLine = string(dontBreaks[0])
	for i, dontbreak := range dontBreaks {
		if i > 0 {
			doBreaks := strings.Split(dontbreak, "do()")
			if len(doBreaks) > 1 {
				for j, dobreak := range doBreaks {
					if j > 0 {
						editedLine = editedLine + string(dobreak)
					}
				}
			}
		}
	}
	outMuls, err = procMatch(editedLine)
	return
}

func procMatch(inTxt string) (outMuls []Product, err error) {
	var outMul Product
	mulBreaks := strings.Split(inTxt, "mul(")
	for _, mulBreak := range mulBreaks {
		orderPair := strings.Split(mulBreak, ")")
		orderPairB := string(orderPair[0])
		orderPairC := strings.Split(orderPairB, ",")
		if len(orderPairC) == 2 {
			outMul.x, err = strconv.Atoi(orderPairC[0])
			if err != nil {
				outMul.x = 0
				err = nil
			}
			outMul.y, err = strconv.Atoi(orderPairC[1])
			if err != nil {
				outMul.y = 0
				err = nil
			}
			outMuls = append(outMuls, outMul)
		}
	}
	return
}

func ProcDayThree(filename string, procDos string) {
	var sumOfProducts int
	var products []Product
	var inTxt string
	var err error
	inTxt, err = readFile(filename)
	if err != nil {
		panic(err)
	}
	if ((procDos == "true" || procDos == "True") || (procDos == "T" || procDos == "t")) || procDos == "TRUE" {
		products, err = procMatchDos(inTxt)
		if err != nil {
			panic(err)
		}
	} else {
		products, err = procMatch(inTxt)
		if err != nil {
			panic(err)
		}
	}
	for _, product := range products {
		sumOfProducts = sumOfProducts + product.Mult()
	}
	fmt.Println(sumOfProducts)
	return
}
