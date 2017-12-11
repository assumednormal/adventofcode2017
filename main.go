package main

import (
	"fmt"
	"io/ioutil"

	"adventofcode2017/advent2017"
)

func readFile(f string) string {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	// day 1
	day01Input := readFile("day01_input.txt")
	fmt.Println(advent2017.SumSkipDigits(day01Input, 1))
	fmt.Println(advent2017.SumSkipDigits(day01Input, len(day01Input)/2))

	// day 2
	day02Input := readFile("day02_input.txt")
	fmt.Println(advent2017.RowMaxDiffChecksum(day02Input))
	fmt.Println(advent2017.RowDivChecksum(day02Input))

	// day 3
	fmt.Println(advent2017.MovesToCenter(325489))
	fmt.Println(advent2017.LeastCumulativeSpiral(325489))

	// day 4
	day04Input := readFile("day04_input.txt")
	fmt.Println(advent2017.ValidPassphrases(day04Input))
	fmt.Println(advent2017.ValidAnagramPassphrases(day04Input))

	// day 5
	day05Input := readFile("day05_input.txt")
	fmt.Println(advent2017.StepsToExit(day05Input))
	fmt.Println(advent2017.StepsToExit2(day05Input))

	// day 6
	day06Input := readFile("day06_input.txt")
	fmt.Println(advent2017.StepsToPriorPattern(day06Input))
	fmt.Println(advent2017.StepsToPriorPattern2(day06Input))

	// day 7
	day07Input := readFile("day07_input.txt")
	fmt.Println(advent2017.FindBottom(day07Input))
	fmt.Println(advent2017.BalanceTower(day07Input))

	// day 8
	day08Input := readFile("day08_input.txt")
	fmt.Println(advent2017.LargestRegister(day08Input))
	fmt.Println(advent2017.LargestRegisterEver(day08Input))

	// day 9
	day09Input := readFile("day09_input.txt")
	fmt.Println(advent2017.StreamScore(day09Input))
	fmt.Println(advent2017.GarbageCounter(day09Input))

	// day 10
	fmt.Println(advent2017.KnotHash(256, []int{14, 58, 0, 116, 179, 16, 1, 104, 2, 254, 167, 86, 255, 55, 122, 244}))
	fmt.Println(advent2017.KnotHash2([]byte("14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244")))

	// day 11
	day11Input := readFile("day11_input.txt")
	fmt.Println(advent2017.HexDistance(day11Input))
	fmt.Println(advent2017.HexDistance2(day11Input))
}
