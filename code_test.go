package advent20201211_test

import (
	"fmt"
	"testing"

	advent "github.com/bwarren2/advent20201211"
)

func TestRecordsFromFile(t *testing.T) {
	// fmt.Println(advent.DataFromFile("sample.txt"))
}

func TestPart1(t *testing.T) {
	fmt.Println("part 1 input:", advent.Part1("input.txt"))
}
func TestPart1Sample(t *testing.T) {
	fmt.Println("part 1 sample:", advent.Part1("sample.txt"))
}

func TestPart2Input(t *testing.T) {
	fmt.Println("part 2 input:", advent.Part2("input.txt"))
}
func TestPart2Sample(t *testing.T) {
	fmt.Println("part 2 sample:", advent.Part2("sample.txt"))
}
