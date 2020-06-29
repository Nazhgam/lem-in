package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type room struct {
	Index int
	Name  string
	Links []string
	Start bool
	End   bool
}
type ants struct {
	AntIndex       int
	AntsPath       []string
	CounterForRoom int
	PathIndex      int
}

var ant []ants
var arrLenOptimalPath []int

func Compare(s1, s2 []string) bool {
	for _, i := range s1 {
		for _, j := range s2 {
			if i != j {
				return false
			}
		}
	}
	return true
}

func collectAntsStruct() {
	ant = make([]ants, numAnt)
	for _, k := range optimalPath {
		arrLenOptimalPath = append(arrLenOptimalPath, len(k))
	}
	for i := 0; i < numAnt; i++ {
		ant[i].distributionAnt(i)

	}
}
func minLen() int {
	t := arrLenOptimalPath[0]
	ret := 0
	for i, k := range arrLenOptimalPath {
		if t > k {
			t = k
			ret = i
		}
	}
	return ret
}
func (ants) distributionAnt(j int) {
	i := minLen()
	ant[j].AntIndex = j + 1
	ant[j].AntsPath = optimalPath[i]
	ant[j].PathIndex = i
	arrLenOptimalPath[i] += 1

}
func prepareForPrint() {
	t := 0
	for i := 0; i < arrLenOptimalPath[0]-1; i++ {
		t = 0
		for j := 0; j < len(ant); j++ {
			if ant[j].CounterForRoom > 0 && ant[j].CounterForRoom < len(ant[j].AntsPath) {
				fmt.Print("L", ant[j].AntIndex, "-", ant[j].AntsPath[ant[j].CounterForRoom], " ")
				ant[j].CounterForRoom++
			}
			if ant[j].PathIndex == t && ant[j].CounterForRoom == 0 {
				fmt.Print("L", ant[j].AntIndex, "-", ant[j].AntsPath[ant[j].CounterForRoom], " ")
				t++
				ant[j].CounterForRoom++
			}

		}
		fmt.Println()
	}
}

var a []room

var paths []string
var differPath [][]string
var numAnt int
var optimalPath [][]string

func antNum(arr []string) {
	for _, k := range arr {
		n, err := strconv.Atoi(k)
		if err == nil {
			numAnt = n
		}
	}

}
func isLink(b string) bool {
	for _, k := range b {
		if k == '-' {
			return true
		}
	}
	return false
}
func onlyRoom(r string) string {
	l := ""
	for _, k := range r {
		if k == ' ' {
			break
		} else {
			l = l + string(k)
		}
	}

	return l
}
func split(b string) (string, string) {
	var b1, b2 string
	t := 0
	for _, k := range b {
		if k != '-' {
			if t == 0 {
				b1 = b1 + string(k)
			} else {
				b2 = b2 + string(k)
			}
		} else {
			t++
		}
	}
	return b1, b2
}
func (room) link(i int, arr []string) {
	for j := 0; j < len(arr); j++ {
		if isLink(arr[j]) {
			link1, link2 := split(arr[j])
			if a[i].Name == link1 {
				a[i].Links = append(a[i].Links, link2)
			} else if a[i].Name == link2 {
				a[i].Links = append(a[i].Links, link1)
			}

		}
	}
}
func (room) start(i int, arr []string) {
	for j := 0; j < len(arr); j++ {
		if isRoom(arr[j]) {
			a[i].Name = onlyRoom(arr[j])
			if arr[j-1] == "##start" {
				a[i].Start = true
			} else if arr[j-1] == "##end" {
				a[i].End = true
			}
			arr[j] = ""
			a[i].link(i, arr)
			break
		}
	}
}
func isRoom(b string) bool {
	for _, k := range b {
		if k == ' ' {
			return true
		}
	}
	return false
}
func readfile(a string) ([]string, int) {
	var arr []string
	countRoom := 0
	file, err := os.Open(a)
	if err != nil {
		fmt.Println("no file like this")
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isRoom(scanner.Text()) {
			countRoom++
		}
		arr = append(arr, scanner.Text())
	}
	return arr, countRoom
}
func sortStruct(a []room) {
	for i := 0; i < len(a); i++ {
		if a[i].Start == true && i != 0 {
			a[0].Index, a[i].Index = a[i].Index, a[0].Index
			a[0], a[i] = a[i], a[0]

		}
		if a[i].End == true && i != len(a)-1 {
			a[len(a)-1].Index, a[i].Index = a[i].Index, a[len(a)-1].Index
			a[len(a)-1], a[i] = a[i], a[len(a)-1]
		}
	}
}

func main() {

	collectRoomStruct()
	collectPath(a[0], "")
	sortArrPath()
	// for _, k := range paths {
	// 	fmt.Println(k)
	// 	fmt.Println()
	// }
	// fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
	splitDifferPath()
	// for _, k := range differPath {
	// 	for _, j := range k {
	// 		fmt.Println(j)
	// 	}
	// 	fmt.Println()
	// }

	optimalPathToTraval()
	// fmt.Println(optimalPath)
	collectAntsStruct()
	// for _, k := range ant {
	// 	fmt.Println(k)
	// }
	for _, k := range ant {
		fmt.Println(k)
	}
	prepareForPrint()
	// printingResult()
}
func collectRoomStruct() {

	fil := os.Args[1:]
	if len(fil) != 1 {
		// fmt.Println("tut")
		os.Exit(0)
	}
	file := fil[0]
	arr, t := readfile(file)
	a = make([]room, t)
	for i := 0; i < t; i++ {
		a[i].Index = i
		a[i].start(i, arr)
	}
	sortStruct(a)
	antNum(arr)

}

func lencounter(a string) int {
	t := -1
	split := strings.Split(a, " -> ")
	for _, k := range split {
		if k != "" {
			t++
		}
	}
	return t
}
func commonLen(a []string) (int, int) {
	t := 0

	max := lencounter(a[len(a)-1])

	for _, k := range a {
		t = t + (max - lencounter(k))
	}

	return t, max
}
func optimalPathToTraval() {
	a := numAnt
	opPath := 0
	t := 0
	l := 0
	var arr []string
	for i, k := range differPath {
		if lencounter(k[0]) == 1 {
			split := strings.Split(k[0], " -> ")
			for i, k := range split {
				if k != "" && i != 0 {
					arr = append(arr, k)
				}
			}
			optimalPath = append(optimalPath, arr)
			return
		}
		ant, max := commonLen(k)

		a = a - ant

		ostatok := a / len(k)
		ost := a % len(k)
		// fmt.Println(ost, ostatok, ant, max)
		if ost != 0 {
			ostatok = ostatok + 1
		}
		t = max + ostatok - 1
		if i == 0 {
			opPath = t
			l = i
		}
		if i > 0 && t < opPath {
			opPath = t
			l = i
		}

		a = numAnt
		t = 0
	}

	for i := 0; i < len(differPath[l]); i++ {
		split := strings.Split(differPath[l][i], " -> ")
		for i, k := range split {
			if k != "" && i != 0 {
				arr = append(arr, k)
			}
		}
		optimalPath = append(optimalPath, arr)
		arr = nil
	}
}
func exist(l, word string) bool {
	split := strings.Split(word, " -> ")
	for _, k := range split {
		if k == l {
			return true
		}
	}
	return false
}
func collectPath(b room, word string) {

	if b.End == true {
		word = word + " -> " + b.Name
		paths = append(paths, word)
		return
	}
	if word == "" {
		word = b.Name
	} else {
		word = word + " -> " + b.Name
	}
	for i := 0; i < len(b.Links); i++ {
		l := findStruct(b.Links[i])
		if !exist(l.Name, word) {
			collectPath(l, word)
		}
	}
}
func findStruct(b string) room {
	var l room
	for i := 0; i < len(a); i++ {
		if a[i].Name == b {
			l = a[i]
		}
	}

	return l
}
func sortArrPath() {
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
}
func isThereSame(s1, s2 string) bool {
	b2 := strings.Split(s1, " -> ")
	b1 := strings.Split(s2, " -> ")
	for i := 1; i < len(b1)-1; i++ {
		for j := 1; j < len(b2)-1; j++ {
			if b1[i] == b2[j] {
				return false
			}
		}

	}
	return true
}
func splitDifferPath() {
	t := 0
	var arr []string
	if len(paths) == 1 {
		arr = append(arr, paths[0])
		differPath = append(differPath, arr)
		return
	}
	for i := 0; i < len(paths)-1; i++ {

		arr = append(arr, paths[i])
		for j := i + 1; j < len(paths); j++ {
			if isThereSame(paths[i], paths[j]) {
				for _, k := range arr {
					if isThereSame(k, paths[j]) {
						t++
					}
				}
				if t == len(arr) {
					arr = append(arr, paths[j])
				}
				t = 0

			}
		}
		differPath = append(differPath, arr)
		arr = nil
	}
}
