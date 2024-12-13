package dayeleven

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fillCache(stone string, blinks int) (outStone []string) {
	inStone := []string{stone}
	for i := 0; i < blinks; i++ {
		inStone = blink(inStone)
	}
	outStone = inStone
	return
}

func blink(stones []string) (newStones []string) {
	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if (len(stone) % 2) == 0 {
			newStoneAInt, err := strconv.ParseInt(stone[:len(stone)/2], 10, 64)
			if err != nil {
				panic(err)
			}
			newStoneStr := strconv.FormatInt(newStoneAInt, 10)
			newStones = append(newStones, newStoneStr)
			newStoneAInt, err = strconv.ParseInt(stone[len(stone)/2:], 10, 64)
			if err != nil {
				panic(err)
			}
			newStoneStr = strconv.FormatInt(newStoneAInt, 10)
			newStones = append(newStones, newStoneStr)
		} else {
			tmpIntStone, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			newStoneInt := int64(tmpIntStone) * int64(2024)
			newStoneStr := strconv.FormatInt(newStoneInt, 10)
			newStones = append(newStones, newStoneStr)
		}
	}
	return
}

func readDayEleven(filename string) (stones []string) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fileString = strings.Trim(fileString, "\n")
	fileArray := strings.Split(fileString, " ")
	for _, stone := range fileArray {
		stones = append(stones, string(stone))
	}
	return
}

func ProcDayEleven(filename string) {
	var stones []string
	numBlinks := 25
	stones = readDayEleven(filename)
	for i := 0; i < numBlinks; i++ {
		stones = blink(stones)
		// fmt.Println(i)
	}
	fmt.Println(len(stones))
	numBlinks = 75
	cycles := 15
	numBlinkAheads := numBlinks / cycles
	stones = readDayEleven(filename)
	reference := make(map[string][]string)
	for i := 0; i < cycles; i++ {
		var newStones []string
		for _, stone := range stones {
			addStones, ok := reference[stone]
			if ok {
				newStones = append(newStones, addStones...)
			} else {
				reference[stone] = fillCache(stone, numBlinkAheads)
				newStones = append(newStones, reference[stone]...)
			}
		}
		stones = newStones
		fmt.Println(len(reference))
	}
	fmt.Println(len(stones))
}
