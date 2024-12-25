package daytwentyone

import (
	"bytes"
	"fmt"
	"os"
)

type Numpad struct {
	From  byte
	To    byte
	Move  [][]byte
	Score []int
}

func buildNumpad() map[byte]map[byte][]string {
	numpad := make(map[byte]map[byte][]string)
	numpad[0x41] = make(map[byte][]string)
	for i := 0; i < 10; i++ {
		numpad[byte(uint8(i))] = make(map[byte][]string)
	}
	// numpad[0x41][0x00] = ["<"]
	// numpad[0x41][]

	return numpad
}

func readDayTwentyOne(filename string) (codes [][]byte) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	codes = bytes.Split(fileByte, []byte{0x0A})
	return
}

func ProcDayTwentyOne(filename string) {
	codes := readDayTwentyOne(filename)
	for _, row := range codes { // This is the loop over all the codes
		for _, inKey := range row { // This is the loop over each of the numbers of the code
			fmt.Println(inKey)
		}
	}
}
