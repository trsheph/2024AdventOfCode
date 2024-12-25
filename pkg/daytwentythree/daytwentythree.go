package daytwentythree

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Edge struct {
	A string
	B string
}

func deepCopy(slice []uint16) []uint16 {
	newSlice := make([]uint16, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(inSlice []uint16, element uint16) []uint16 {
	for i, val := range inSlice {
		if val == element {
			return append(inSlice[:i], inSlice[i+1:]...)
		}
	}
	return inSlice
}

func intersection(a, b []uint16) []uint16 {
	result := []uint16{}
	bucket := make(map[uint16]bool)
	for _, val := range a {
		if _, ok := bucket[val]; !ok {
			bucket[val] = true
		}
	}
	for _, val := range b {
		if _, ok := bucket[val]; ok {
			result = append(result, val)
			delete(bucket, val) // remove from bucket to avoid duplicates
		}
	}
	return result
}

func removeDupUint16(a []uint16) []uint16 {
	result := []uint16{}
	unique := map[uint16]uint16{}
	for _, val := range a {
		if _, ok := unique[val]; !ok {
			result = append(result, val)
			unique[val] = val
		}
	}
	return result
}

func removeDup3uint16(a [][]uint16) [][]uint16 {
	result := [][]uint16{}
	for _, tmpSlice := range a {
		addFlag := true
		for _, row := range result {
			if (tmpSlice[0] == row[0] && tmpSlice[1] == row[1]) && (tmpSlice[2] == row[2]) {
				addFlag = false
			}
		}
		if addFlag {
			result = append(result, tmpSlice)
		}
	}
	return result
}

func fillCliques(graphMap map[uint16][]uint16, R, P, X []uint16, cliques *[][]uint16) {
	if len(P) == 0 && len(X) == 0 {
		*cliques = append(*cliques, deepCopy(R))
		return
	}
	for i := 0; i < len(P); i++ {
		newNode := P[i]
		newR := deepCopy(R)
		newR = append(newR, newNode)
		newP := intersection(graphMap[newNode], P)
		newX := intersection(graphMap[newNode], X)
		fillCliques(graphMap, newR, newP, newX, cliques)
		P = removeElement(P, newNode)
		X = append(X, newNode)
		i--
	}
}

func dfs(allNodes map[uint16][]uint16, node uint16, visited map[uint16]bool) [][]uint16 {
	// func dfs(graphMap map[uint16][]uint16) [][]uint16 {
	var cliques [][]uint16
	var R, P, X []uint16
	visited[node] = true
	for _, neighbor := range allNodes[node] {
		if !visited[neighbor] {
			dfs(allNodes, neighbor, visited)
		}
	}
	var graphMap map[uint16][]uint16
	graphMap = make(map[uint16][]uint16)
	for nd := range visited {
		graphMap[nd] = allNodes[nd]
	}
	// fmt.Println("Length of graphMap being tested: ", len(graphMap))
	for snode := range graphMap {
		P = append(P, snode)
	}
	fillCliques(graphMap, R, P, X, &cliques)
	return cliques
}

func seedClique(nodes map[uint16][]uint16) (tnodes map[uint16][]uint16) {
	tnodes = make(map[uint16][]uint16)
	for nd, subset := range nodes {
		if nd > 493 && nd < 520 {
			tnodes[nd] = subset
		}
	}
	return
}

func cliqueB(nodes map[uint16][]uint16, tnode []uint16) (snodes map[uint16][]uint16) {
	snodes = make(map[uint16][]uint16)
	for _, nd := range tnode {
		snodes[nd] = nodes[nd]
	}
	return
}

func deconv(a uint16) (sA string) {
	k := uint16(0)
	rs := "abcdefghijklmnopqrstuvwxyz"
	dconv := make(map[uint16]string)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			dconv[k] = rs[i:i+1] + rs[j:j+1]
			k = k + 1
		}
	}
	sA = dconv[a]
	return
}

func getCliques(nodes map[uint16][]uint16, tnodes map[uint16][]uint16) (maxClique []uint16) {
	maxClique = []uint16{}
	allCliques := [][]uint16{}
	counter := 0
	fmt.Println("Number of tnodes:", len(tnodes))
	for tnode := range tnodes {
		doFlag := true
	outerLoop:
		for _, rw := range allCliques {
			for _, cl := range rw {
				if tnode == cl {
					doFlag = false
					break outerLoop
				}
			}
		}
		if doFlag {
			visited := make(map[uint16]bool)
			cliques := dfs(nodes, tnode, visited)
			fmt.Println("Clique complete")
			for _, cliqueE := range cliques {
				allCliques = append(allCliques, cliqueE)
				if len(cliqueE) > len(maxClique) {
					maxClique = cliqueE
				}
			}
		}
		counter += 1
		fmt.Println(counter)
	}
	return
}

func initDTT(elfnet []Edge) (nodes map[uint16][]uint16) {
	var k int
	k = 0
	rs := "abcdefghijklmnopqrstuvwxyz"
	conv := make(map[string]int)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			conv[rs[i:i+1]+rs[j:j+1]] = k
			k = k + 1
		}
	}
	nodes = make(map[uint16][]uint16)
	for _, cmp := range elfnet {
		Au := uint16(conv[cmp.A])
		Bu := uint16(conv[cmp.B])
		nodes[Au] = append(nodes[Au], Bu)
		nodes[Bu] = append(nodes[Bu], Au)
	}
	for nd := range nodes {
		nodes[nd] = removeDupUint16(nodes[nd])
		slices.Sort(nodes[nd])
	}
	return
}

func readDTT(filename string) (elfnet []Edge) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileByte)
	fileRows := strings.Split(fileString, "\n")
	for _, row := range fileRows {
		computers := strings.Split(row, "-")
		if len(computers) == 2 {
			var tmpNet Edge
			tmpNet.A = computers[0]
			tmpNet.B = computers[1]
			elfnet = append(elfnet, tmpNet)
		}
	}
	return
}

func partOne(nodes map[uint16][]uint16, tnodes map[uint16][]uint16) (outSlice [][]uint16) {
	for tnd := range tnodes {
		snodes := cliqueB(nodes, tnodes[tnd]) // builds the second node on the networks second:[thirds]
		for snd := range snodes {
			unodes := cliqueB(nodes, snodes[snd]) // builds the third node on the networks third:[fourths]
			for und := range unodes {
				vnodes := cliqueB(nodes, unodes[und])
				for vnd := range vnodes {
					if vnd == tnd {
						tmpSlice := []uint16{tnd, snd, und}
						slices.Sort(tmpSlice)
						outSlice = append(outSlice, tmpSlice)
					}
				}
			}
		}
	}
	outSlice = removeDup3uint16(outSlice)
	return
}

func partTwo(nodes map[uint16][]uint16, tnodes map[uint16][]uint16) (outString string) {
	maxClique := getCliques(nodes, tnodes)
	slices.Sort(maxClique)
	fmt.Println(maxClique)
	outString = ""
	for _, node := range maxClique {
		outString = outString + "," + deconv(node)
	}
	return
}

func ProcDayTwentyThree(filename string) {
	elfnet := readDTT(filename)
	nodes := initDTT(elfnet)
	tnodes := seedClique(nodes) // seeds with those beginning with t
	oneSlice := partOne(nodes, tnodes)
	fmt.Println(len(oneSlice))
	outString := partTwo(nodes, tnodes)
	fmt.Println(outString[1:])
}
