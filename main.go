package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var digMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	var total int
	b, err := os.ReadFile("pt1.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(b), "\n")
	for i := 0; i < len(s); i++ {
		total += intFinderV2(s[i])
	}

	fmt.Println("ANSWER: ", total)

}

func intFinder(str string) int {
	fmt.Println("----")
	fmt.Println("val: ", str)
	var x, y int = 0, len(str)
	var a, b int
	var matchA, matchB bool
	for x <= y {
		firstIndex := str[x : x+1]
		firstNum, e1 := strconv.Atoi(firstIndex)
		secondIndex := str[y-1 : y]
		secondNum, e2 := strconv.Atoi(secondIndex)
		if e1 == nil && !matchA {
			a = firstNum
			matchA = true
			fmt.Println("match found for first number:", a)
		}
		if e2 == nil && !matchB {
			b = secondNum
			matchB = true
			fmt.Println("match found for second number:", b)
		}
		if matchA && matchB {
			fmt.Printf("made %v iterations\n", x)
			strNum := strconv.Itoa(a) + strconv.Itoa(b)
			res, _ := strconv.Atoi(strNum)
			fmt.Println("result: ", res)
			return res
		}

		if !matchA {
			x++
		}
		if !matchB {
			y--
		}
	}

	if !matchB {
		strNum := strconv.Itoa(a) + strconv.Itoa(a)
		res, _ := strconv.Atoi(strNum)
		fmt.Println("result: ", res)
		return res
	} else {
		strNum := strconv.Itoa(b) + strconv.Itoa(b)
		res, _ := strconv.Atoi(strNum)
		fmt.Println("result: ", res)
		return res
	}
}

func intFinderV2(str string) int {
	fmt.Println("----")
	fmt.Println("val: ", str)
	dict := map[int]string{}
	for _, digit := range digits {
		firstIndex := strings.Index(str, digit)
		secondIndex := strings.LastIndex(str, digit)
		if firstIndex == -1 || secondIndex == -1 {
			continue
		}
		if secondIndex > firstIndex {
			dict[secondIndex] = digit
			dict[firstIndex] = digit
		} else {
			dict[firstIndex] = digit
		}

	}

	keys := make([]int, 0)
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	dig1 := dict[keys[0]]
	dig2 := dict[keys[len(keys)-1]]
	if digMap[dig1] != "" {
		dig1 = digMap[dig1]
	}
	if digMap[dig2] != "" {
		dig2 = digMap[dig2]
	}
	res, _ := strconv.Atoi(dig1 + dig2)
	fmt.Println("res: ", res)
	return res
}
