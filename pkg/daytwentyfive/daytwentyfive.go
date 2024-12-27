package daytwentyfive

import (
	"bytes"
	"fmt"
	"os"
)

type LockAndKey struct {
	LockID int
	KeyID  int
}

func createMatrix(r int, c int) ([][]uint8, [][]bool, [][]byte) {
	uM := make([][]uint8, r)
	boolM := make([][]bool, r)
	byteM := make([][]byte, r)
	for i := range uM {
		uM[i] = make([]uint8, c)
		boolM[i] = make([]bool, c)
		byteM[i] = make([]byte, c)
	}
	return uM, boolM, byteM
}

func transpose(inMat [][]byte) (outMat [][]byte) {
	_, _, outMat = createMatrix(len(inMat[0]), len(inMat))
	for i := range inMat {
		for j := range inMat[i] {
			outMat[j][i] = inMat[i][j]
		}
	}
	return
}

func partOne(imgs map[int][]uint8) (lak []LockAndKey) { // bug here
	for i := 0; i < len(imgs); i++ {
		for j := i + 1; j < len(imgs); j++ {
			addFlag := true
			for k := range imgs[i] {
				result := imgs[i][k] & imgs[j][k]
				if result > uint8(0x00) {
					// fmt.Println(i, j, result)
					addFlag = false
				}
			}
			if addFlag {
				lak = append(lak, LockAndKey{j, i})
			}
		}
	}
	return
}

func imgToBinary(imgs map[int][][]byte) (ibin map[int][]uint8) {
	sharp := byte('#')
	flat := byte('.')
	ibin = make(map[int][]uint8)
	for imgind, img := range imgs {
		resultSlice := []uint8{}
		for _, row := range img {
			var result uint8
			for _, val := range row {
				result <<= 1
				if val == sharp {
					result |= 1
				} else if val == flat {
				} else {
					fmt.Println(imgind)
					fmt.Println(result)
					fmt.Println(row)
					fmt.Println(val)
					err := fmt.Errorf("unrecognized character in imgToBinary")
					panic(err)
				}
			}
			resultSlice = append(resultSlice, result)
		}
		ibin[imgind] = resultSlice
	}
	return
}

func readDTwFi(filename string) (imgs map[int][][]byte) {
	imgs = make(map[int][][]byte)
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var mat [][]byte
	_, _, mat = createMatrix(7, 5)
	var rowCount, imgCount int
	rowCount = 0
	imgCount = 0
	rows := bytes.Split(fileBytes, []byte{0x0A})
	for i := range rows {
		if len(rows[i]) < 2 {
			// fmt.Println("New matrix:", mat)
			tmpImg := transpose(mat)
			// fmt.Println("new transpose:", tmpImg)
			imgs[imgCount] = tmpImg
			_, _, mat = createMatrix(7, 5)
			rowCount = 0
			imgCount++
		} else {
			for j, val := range rows[i] {
				mat[rowCount][j] = val
			}
			rowCount++
		}
	}
	delete(imgs, 500)
	return
}

func ProcDayTwentyFive(filename string) {
	// fmt.Println("hi")
	imgs := readDTwFi(filename)
	// fmt.Println(imgs)
	imgbins := imgToBinary(imgs)
	// fmt.Println(imgbins)
	lak := partOne(imgbins)
	fmt.Println(len(lak))
}
