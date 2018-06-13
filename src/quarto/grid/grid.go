package grid

import (
	"github.com/ahl5esoft/golang-underscore"
	"math"
)

// Point define data for a grid coordinate
type Point struct {
	X int
	Y int
}

// Box define data for a grid position with evaluations among its grid
type Box struct {
	Position           Point
	AlignedPieceNumber int
}

// GetNewGrid return a blank grid of defined size
func GetNewGrid(size int) [][]int {
	newGrid := [][]int{}
	for i := 0; i < size; i++ {
		newGrid = append(newGrid, []int{})
		for j := 0; j < size; j++ {
			newGrid[i] = append(newGrid[i], 0)
		}
	}
	return newGrid
}

// CopyGrid return a copy of the source grid
func CopyGrid(grid [][]int) [][]int {
	newGrid := GetNewGrid(len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			newGrid[i][j] = grid[i][j]
		}
	}
	return newGrid
}

// IsWinningPosition return true if the piece placed at these coordinates make a winning situation
func IsWinningPosition(x int, y int, grid [][]int, piece int) bool {
	testGrid := CopyGrid(grid)
	testGrid[x][y] = piece
	if IsWinningLine(GetPiecesRaw(x, y, testGrid)) {
		return true
	}
	if IsWinningLine(GetPiecesColumn(x, y, testGrid)) {
		return true
	}
	if IsWinningLine(GetPiecesSlashDiag(x, y, testGrid)) {
		return true
	}
	if IsWinningLine(GetPiecesBackSlashDiag(x, y, testGrid)) {
		return true
	}
	return false
}

// GetPiecesRaw return an array of the raw of pieces aligned in [x,y] coordinates
func GetPiecesRaw(x int, y int, grid [][]int) []int {
	return grid[y]
}

// GetPiecesColumn return an array of the column of pieces aligned in [x,y] coordinates
func GetPiecesColumn(x int, y int, grid [][]int) []int {
	piecesLine := []int{}
	for i := 0; i < len(grid); i++ {
		piecesLine = append(piecesLine, grid[i][x])
	}
	return piecesLine
}

// GetPiecesSlashDiag return an array of the diag (slash oriented) of pieces aligned in [x,y] coordinates
func GetPiecesSlashDiag(x int, y int, grid [][]int) []int {
	piecesLine := []int{}
	if x == y {
		for i := 0; i < len(grid); i++ {
			piecesLine = append(piecesLine, grid[i][i])
		}
	}
	return piecesLine
}

// GetPiecesBackSlashDiag return an array of the diag (backslash oriented) of pieces aligned in [x,y] coordinates
func GetPiecesBackSlashDiag(x int, y int, grid [][]int) []int {
	piecesLine := []int{}
	if x == len(grid)-y-1 {
		for i := 0; i < len(grid); i++ {
			piecesLine = append(piecesLine, grid[i][len(grid)-i-1])
		}
	}
	return piecesLine
}

// IsWinningLine return true if all piece in array makes a winning situation when aligned
func IsWinningLine(piecesLine []int) bool {
	if len(piecesLine) == 0 {
		return false
	}
	var indexEmptyPiece = underscore.FindIndex(piecesLine, func(n, _ int) bool {
		return n == 0
	})
	if indexEmptyPiece >= 0 {
		return false
	}

	bitInverser := int(math.Pow(2, float64(len(piecesLine))) - 1)
	propertiesAtTrue := underscore.Reduce(piecesLine, func(prev int, curr, _ int) int {
		return prev & (curr - 1)
	}, 1)

	propertiesAtFalse := underscore.Reduce(piecesLine, func(prev int, curr, _ int) int {
		return prev & ((curr - 1) ^ bitInverser)
	}, 1)
	return propertiesAtTrue != 0 || propertiesAtFalse != 0
}

// GetSafestBoxes return list of boxes in the grid where grid is the less filled
func GetSafestBoxes(grid [][]int) []Box {
	boxList := GetEmptyBoxes(grid)
	minValue := MinOppucationValue(boxList)
	safestBoxList := underscore.Select(boxList, func(n Box, _ int) bool {
		return minValue == n.AlignedPieceNumber
	})
	return safestBoxList.([]Box)
}

// GetEmptyBoxes return list of empty boxes in the grid
func GetEmptyBoxes(grid [][]int) []Box {
	boxList := []Box{}
	size := len(grid)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				boxList = append(boxList, Box{Point{j, i}, GetOppupationValue(grid, i, j)})
			}
		}
	}
	return boxList
}

// GetOppupationValue return an evaluation of grid occupation non lines for a given coordinate
func GetOppupationValue(grid [][]int, i int, j int) int {
	piecesRawNumber := BoxFilledNumber(GetPiecesRaw(j, i, grid))
	piecesColumnNumber := BoxFilledNumber(GetPiecesColumn(j, i, grid))
	piecesSlashDiagNumber := BoxFilledNumber(GetPiecesSlashDiag(j, i, grid))
	piecesBackSlashDiagNumber := BoxFilledNumber(GetPiecesBackSlashDiag(j, i, grid))
	return piecesRawNumber + piecesColumnNumber + piecesSlashDiagNumber + piecesBackSlashDiagNumber
}

// BoxFilledNumber count number of filled boxes in a list
func BoxFilledNumber(piecesLine []int) int {
	number := underscore.Reduce(piecesLine, func(prev int, curr, _ int) int {
		if curr > 0 {
			return prev + 1
		}
		return prev
	}, 0)
	return number.(int)
}

// MinOppucationValue return the value of the min LinesOccupationValue in the list
func MinOppucationValue(pointList []Box) int {
	number := underscore.Reduce(pointList, func(prev int, curr Box, _ int) int {
		if curr.AlignedPieceNumber < prev {
			return curr.AlignedPieceNumber
		}
		return prev
	}, 9999)
	return number.(int)
}
