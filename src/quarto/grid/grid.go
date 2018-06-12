package grid

var emptyCoord = [2]int{-1, -1}

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
	if (IsWinningLine(GetPiecesRaw(x,y,testGrid))) {
		return true
	}
	if (IsWinningLine(GetPiecesColumn(x,y,testGrid))) {
		return true
	}
	if (IsWinningLine(GetPiecesSlashDiag(x,y,testGrid))) {
		return true
	}
	if (IsWinningLine(GetPiecesBackSlashDiag(x,y,testGrid))) {
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
	if (x == y) {
		for i := 0; i < len(grid); i++ {
			piecesLine = append(piecesLine, grid[i][i])
		}
	}
	return piecesLine
}

// GetPiecesBackSlashDiag return an array of the diag (backslash oriented) of pieces aligned in [x,y] coordinates
func GetPiecesBackSlashDiag(x int, y int, grid [][]int) []int {
	piecesLine := []int{}
	if (x == len(grid) - y - 1) {
		for i := 0; i < len(grid); i++ {
			piecesLine = append(piecesLine, grid[i][len(grid) - i - 1])
		}
	}
	return piecesLine
}

// IsWinningLine return true if all piece in array makes a winning situation when aligned
func IsWinningLine(piecesLine []int) bool {
	if (len(piecesLine) == 0) {
		return false
	}
	return false
}
