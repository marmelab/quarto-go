package grid

import (
	"github.com/ahl5esoft/golang-underscore"
	"math"
	"strconv"
	"strings"
)

// Point define data for a grid coordinate
type Point struct {
	X int
	Y int
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
	testGrid[y][x] = piece
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

	return CountIdenticalCaracteristics(piecesLine, len(piecesLine)) > 0
}

// GetSafestBoxesIncludingPieceChoice return list of boxes in the grid where grid is the less filled with less common caracteristics
func GetSafestBoxesIncludingPieceChoice(grid [][]int, piece int) []Point {
	boxList := GetEmptyBoxes(grid)
	minValue := MinPositionScoreForPiece(grid, boxList, piece)

	safestBoxList := underscore.Select(boxList, func(n Point, _ int) bool {
		return minValue == GetPositionScoreForPiece(grid, n.Y, n.X, piece)
	})
	if safestBoxList == nil {
		return []Point{}
	}
	return safestBoxList.([]Point)
}

// GetBlockingBoxesIncludingPieceChoice return list of boxes in the grid where player can block a nearly winning line
func GetBlockingBoxesIncludingPieceChoice(grid [][]int, piece int) []Point {
	boxList := GetEmptyBoxes(grid)

	blockingBoxList := underscore.Select(boxList, func(n Point, _ int) bool {
		return IsBlockingPositionForPiece(grid, n.Y, n.X, piece)
	})
	if blockingBoxList == nil {
		return []Point{}
	}
	return blockingBoxList.([]Point)
}

// GetEmptyBoxes return list of empty boxes in the grid
func GetEmptyBoxes(grid [][]int) []Point {
	boxList := []Point{}
	size := len(grid)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				boxList = append(boxList, Point{j, i})
			}
		}
	}
	return boxList
}

// GetAlignedPieceNumber return an evaluation of grid occupation of lines for a given coordinate
func GetAlignedPieceNumber(grid [][]int, i int, j int) int {
	piecesRawNumber := BoxFilledNumber(GetPiecesRaw(j, i, grid))
	piecesColumnNumber := BoxFilledNumber(GetPiecesColumn(j, i, grid))
	piecesSlashDiagNumber := BoxFilledNumber(GetPiecesSlashDiag(j, i, grid))
	piecesBackSlashDiagNumber := BoxFilledNumber(GetPiecesBackSlashDiag(j, i, grid))

	return piecesRawNumber + piecesColumnNumber + piecesSlashDiagNumber + piecesBackSlashDiagNumber
}

// GetPositionScoreForPiece return an evaluation of grid occupation of lines for a given coordinate and a given piece
func GetPositionScoreForPiece(grid [][]int, i int, j int, piece int) int {
	piecesRawNumber := BoxFilledNumber(GetPiecesRaw(j, i, grid))
	piecesColumnNumber := BoxFilledNumber(GetPiecesColumn(j, i, grid))
	piecesSlashDiagNumber := BoxFilledNumber(GetPiecesSlashDiag(j, i, grid))
	piecesBackSlashDiagNumber := BoxFilledNumber(GetPiecesBackSlashDiag(j, i, grid))
	pieceNumber := piecesRawNumber + piecesColumnNumber + piecesSlashDiagNumber + piecesBackSlashDiagNumber

	newGrid := CopyGrid(grid)
	newGrid[i][j] = piece
	piecesRawCommonCaracteristics := CountIdenticalCaracteristics(GetPiecesRaw(j, i, newGrid), len(newGrid))
	piecesColumnCommonCaracteristics := CountIdenticalCaracteristics(GetPiecesColumn(j, i, newGrid), len(newGrid))
	piecesSlashDiagCommonCaracteristics := CountIdenticalCaracteristics(GetPiecesSlashDiag(j, i, newGrid), len(newGrid))
	piecesBackSlashDiagCommonCaracteristics := CountIdenticalCaracteristics(GetPiecesBackSlashDiag(j, i, newGrid), len(newGrid))
	commonCaracteristics := piecesRawCommonCaracteristics + piecesColumnCommonCaracteristics + piecesSlashDiagCommonCaracteristics + piecesBackSlashDiagCommonCaracteristics

	return pieceNumber + commonCaracteristics
}

// IsBlockingPositionForPiece return an evaluation of grid occupation of lines for a given coordinate and a given piece
func IsBlockingPositionForPiece(grid [][]int, i int, j int, piece int) bool {
	isBlocking := false
	piecesRaw := GetPiecesRaw(j, i, grid)
	// IsBlockingLineForPiece(GetPiecesRaw(j, i, grid), len(grid), piece)
	if BoxFilledNumber(piecesRaw) == 3 && CountIdenticalCaracteristics(piecesRaw, len(grid)) > 0 {
		isBlocking = true
		piecesRaw = append(piecesRaw, piece)
		if CountIdenticalCaracteristics(piecesRaw, len(grid)) > 0 {
			return false
		}
	}
	piecesColumn := GetPiecesRaw(j, i, grid)
	if BoxFilledNumber(piecesColumn) == 3 && CountIdenticalCaracteristics(piecesColumn, len(grid)) > 0 {
		isBlocking = true
		piecesColumn = append(piecesColumn, piece)
		if CountIdenticalCaracteristics(piecesColumn, len(grid)) > 0 {
			return false
		}
	}
	piecesSlashDiag := GetPiecesRaw(j, i, grid)
	if BoxFilledNumber(piecesSlashDiag) == 3 && CountIdenticalCaracteristics(piecesSlashDiag, len(grid)) > 0 {
		isBlocking = true
		piecesSlashDiag = append(piecesSlashDiag, piece)
		if CountIdenticalCaracteristics(piecesSlashDiag, len(grid)) > 0 {
			return false
		}
	}
	piecesBackSlashDiag := GetPiecesRaw(j, i, grid)
	if BoxFilledNumber(piecesBackSlashDiag) == 3 && CountIdenticalCaracteristics(piecesBackSlashDiag, len(grid)) > 0 {
		isBlocking = true
		piecesBackSlashDiag = append(piecesBackSlashDiag, piece)
		if CountIdenticalCaracteristics(piecesBackSlashDiag, len(grid)) > 0 {
			return false
		}
	}
	return isBlocking
}

func IsBlockingLineForPiece(line []int, gridSize int, piece int) bool {
	isBlocking := false
	if BoxFilledNumber(line) == 3 && CountIdenticalCaracteristics(line, gridSize) > 0 {
		isBlocking = true
		line = append(line, piece)
		if CountIdenticalCaracteristics(line, gridSize) > 0 {
			return false
		}
	}
	return isBlocking
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

// MinPositionScoreForPiece return the value of the min position score for the given list for the given piece
func MinPositionScoreForPiece(grid [][]int, pointList []Point, piece int) int {
	number := underscore.Reduce(pointList, func(prev int, curr Point, _ int) int {
		positionScore := GetPositionScoreForPiece(grid, curr.Y, curr.X, piece)
		if positionScore < prev {
			return positionScore
		}
		return prev
	}, 9999)
	return number.(int)
}

// CountIdenticalCaracteristics return the number of identicals caracteristics in a piece list for a given grid size
func CountIdenticalCaracteristics(pieceList []int, gridSize int) int {
	pieceListWithoutEmptyBox := GetListPieceMinusPiece(pieceList, 0)

	if len(pieceListWithoutEmptyBox) < 2 {
		return 0
	}

	bitInverser := int(math.Pow(2, float64(gridSize)) - 1)
	propertiesAtTrue := underscore.Reduce(pieceListWithoutEmptyBox, func(prev int, curr, _ int) int {
		return prev & (curr - 1)
	}, bitInverser)

	propertiesAtFalse := underscore.Reduce(pieceListWithoutEmptyBox, func(prev int, curr, _ int) int {
		return prev & ((curr - 1) ^ bitInverser)
	}, bitInverser)
	stringReprensentationForPropertiesAtTrue := strconv.FormatInt(int64(propertiesAtTrue.(int)), 2)
	stringReprensentationForPropertiesAtFalse := strconv.FormatInt(int64(propertiesAtFalse.(int)), 2)

	numberOfCommonPropertiesAtTrue := strings.Count(stringReprensentationForPropertiesAtTrue, "1")
	numberOfCommonPropertiesAtFalse := strings.Count(stringReprensentationForPropertiesAtFalse, "1")

	return numberOfCommonPropertiesAtTrue + numberOfCommonPropertiesAtFalse
}
