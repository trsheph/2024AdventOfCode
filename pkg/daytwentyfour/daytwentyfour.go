package daytwentyfour

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	OpA  string
	OpB  string
	Oper string
	NewR string
}

func procRegs(inRegs []string) map[string]bool {
	regMap := make(map[string]bool)
	var strS []string
	for _, inReg := range inRegs {
		strS = strings.Split(inReg, ": ")
		if strS[1] == "0" {
			regMap[strS[0]] = false
		} else if strS[1] == "1" {
			regMap[strS[0]] = true
		} else {
			err := fmt.Errorf("problem with procRegs")
			panic(err)
		}
	}
	return regMap
}

func procOpB(opCache []Operation, registers map[string]bool) (newOpCache []Operation) {
	for _, op := range opCache {
		_, okA := registers[op.OpA]
		_, okB := registers[op.OpB]
		if okA && okB {
			if op.Oper == "AND" {
				registers[op.NewR] = registers[op.OpA] && registers[op.OpB]
			} else if op.Oper == "OR" {
				registers[op.NewR] = registers[op.OpA] || registers[op.OpB]
			} else if op.Oper == "XOR" {
				registers[op.NewR] = registers[op.OpA] != registers[op.OpB]
			} else {
				err := fmt.Errorf("error reading operations in procOps")
				panic(err)
			}
		} else {
			newOpCache = append(opCache, Operation{op.OpA, op.OpB, op.Oper, op.NewR})
		}
	}
	return
}

func procOps(inOps []string, registers map[string]bool) {
	var strO, leftSl []string
	var newR, leftS, opA, opB, oper string
	var opCache []Operation
	for _, inOp := range inOps {
		strO = strings.Split(inOp, " -> ")
		newR = strO[1]
		leftS = strO[0]
		leftSl = strings.Split(leftS, " ")
		opA = leftSl[0]
		opB = leftSl[2]
		oper = leftSl[1]
		_, okA := registers[opA]
		_, okB := registers[opB]
		if okA && okB {
			// fmt.Println("registers exist: ", opA, opB, oper)
			if oper == "AND" {
				registers[newR] = registers[opA] && registers[opB]
			} else if oper == "OR" {
				registers[newR] = registers[opA] || registers[opB]
			} else if oper == "XOR" {
				registers[newR] = registers[opA] != registers[opB]
			} else {
				err := fmt.Errorf("error reading operations in procOps")
				panic(err)
			}
		} else {
			opCache = append(opCache, Operation{opA, opB, oper, newR})
		}
		// fmt.Println(registers)
	}
	for len(opCache) > 0 {
		opCache = procOpB(opCache, registers)
	}
}

func printZregs(registers map[string]bool) {
	var outString string
	orderedMap := make(map[int]string)
	for reg, val := range registers {
		if reg[:1] == "z" {
			regPos, err := strconv.Atoi(reg[1:])
			if err != nil {
				panic(err)
			}
			if val {
				orderedMap[regPos] = "1"
			} else {
				orderedMap[regPos] = "0"
			}
		}
	}
	// tln(orderedMap)
	for i := 0; i < len(orderedMap); i++ {
		outString = orderedMap[i] + outString
	}
	fmt.Println(outString)
	if i, err := strconv.ParseInt(outString, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func readDTwFo(filename string) (registers map[string]bool) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileString, "\n")
	var opFlag, regFlag bool
	opFlag = false
	regFlag = true
	var inRegs, inOps []string
	for _, row := range fileRows {
		if len(row) < 2 && regFlag {
			opFlag = true
			regFlag = false
		}
		if (regFlag && !(opFlag)) && len(row) > 2 {
			inRegs = append(inRegs, row)
		} else if (opFlag && !(regFlag)) && len(row) > 2 {
			inOps = append(inOps, row)
		}
	}
	registers = procRegs(inRegs)
	// fmt.Println(registers)
	procOps(inOps, registers)
	return
}

func ProcDayTwentyFour(filename string) {
	registers := readDTwFo(filename)
	// fmt.Println(registers)
	printZregs(registers)
}
