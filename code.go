package advent20201211

import (
	"bufio"
	"fmt"
	"os"
)

// DataFromFile gets the nums
func DataFromFile(filename string) (returnlist [][]rune) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		sublist := make([]rune, 0)
		text := scanner.Text()
		for _, rune := range text {
			sublist = append(sublist, rune)
		}

		returnlist = append(returnlist, sublist)
	}
	return
}

var outerIndices = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

// FillSeats does work
func FillSeats(input [][]rune) [][]rune {
	newState := make([][]rune, len(input))
	for i := 0; i < len(input); i++ {
		newStateRow := make([]rune, 0)
		for j := 0; j < len(input[i]); j++ {
			newStateRow = append(newStateRow, ' ')
		}
		newState[i] = newStateRow
	}
	for rowIdx, row := range input {
		for colIdx, value := range row {
			if value == '.' {
				newState[rowIdx][colIdx] = '.'
			} else {
				totalSeated := 0
				for _, Adjs := range outerIndices {
					rowAdj := Adjs[0]
					colAdj := Adjs[1]
					lookupRow := rowIdx + rowAdj
					lookupCol := colIdx + colAdj
					// fmt.Printf("%v, %v %v %v", lookupRow, lookupCol, len(row), row)
					if lookupRow >= 0 && lookupRow < len(input) &&
						lookupCol >= 0 && lookupCol < len(row) {
						// fmt.Println('.')
						if input[lookupRow][lookupCol] == '#' {
							totalSeated++
						}
					}
				}
				if totalSeated == 0 {
					newState[rowIdx][colIdx] = '#'
				} else if totalSeated >= 4 {
					newState[rowIdx][colIdx] = 'L'
				} else {
					newState[rowIdx][colIdx] = input[rowIdx][colIdx]
				}
			}
		}

	}
	return newState
}

// FillSeatsV2 does work
func FillSeatsV2(input [][]rune) [][]rune {
	newState := make([][]rune, len(input))
	for i := 0; i < len(input); i++ {
		newStateRow := make([]rune, 0)
		for j := 0; j < len(input[i]); j++ {
			newStateRow = append(newStateRow, ' ')
		}
		newState[i] = newStateRow
	}

	for rowIdx, row := range input {
		for colIdx, value := range row {
			if value == '.' {
				newState[rowIdx][colIdx] = '.'
			} else {
				totalSeated := 0
				for _, Adjs := range outerIndices {
					rowAdj := Adjs[0]
					colAdj := Adjs[1]
					lookupRow := rowIdx + rowAdj
					lookupCol := colIdx + colAdj
					// fmt.Printf("%v, %v %v %v", lookupRow, lookupCol, len(row), row)
					for lookupRow >= 0 && lookupRow < len(input) &&
						lookupCol >= 0 && lookupCol < len(row) {
						if input[lookupRow][lookupCol] == '#' {
							totalSeated++
							break
						} else if input[lookupRow][lookupCol] == 'L' {
							break
						}
						lookupRow += rowAdj
						lookupCol += colAdj
					}
				}
				if totalSeated == 0 {
					newState[rowIdx][colIdx] = '#'
				} else if totalSeated >= 5 {
					newState[rowIdx][colIdx] = 'L'
				} else {
					newState[rowIdx][colIdx] = input[rowIdx][colIdx]
				}
			}
		}
	}
	// fmt.Println(newState)
	return newState
}

// NumSeated does work
func NumSeated(input [][]rune) (count int) {
	for _, row := range input {
		for _, value := range row {
			if value == '#' {
				count++
			}
		}
	}
	return
}

// Part1 answers part 1
func Part1(filename string) (count int) {
	firstState := DataFromFile(filename)
	newState := FillSeats(firstState)
	oldSeated := NumSeated(firstState)
	newSeated := NumSeated(newState)
	for oldSeated != newSeated {
		fmt.Print(".")
		oldSeated = newSeated
		newState = FillSeats(newState)
		newSeated = NumSeated(newState)
	}

	// spew.Dump(NumSeated(newState))
	return newSeated
}

// Part2 answers part 2
func Part2(filename string) (count int) {
	firstState := DataFromFile(filename)
	newState := FillSeatsV2(firstState)
	oldSeated := NumSeated(firstState)
	newSeated := NumSeated(newState)
	for oldSeated != newSeated {
		fmt.Print(".")

		oldSeated = newSeated
		newState = FillSeatsV2(newState)
		newSeated = NumSeated(newState)
	}

	// spew.Dump(NumSeated(newState))
	return newSeated

}
