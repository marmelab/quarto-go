package tests

import (
	"quarto/grid"
	"reflect"
	"testing"
)

func TestCopyGridShouldReturnANewGridEqualToSource(t *testing.T) {
	var sourceGrid = grid.GetNewGrid(5)
	var newGrid = grid.CopyGrid(sourceGrid)
	if !reflect.DeepEqual(newGrid, sourceGrid) {
		t.Errorf("Source grid should be equal to new grid")
	}
}

func TestCopyGridShouldReturnANewGridNotEqualToSourceAfterChanges(t *testing.T) {
	var sourceGrid = grid.GetNewGrid(4)
	var newGrid = grid.CopyGrid(sourceGrid)
	newGrid[2][0] = 3
	if reflect.DeepEqual(newGrid, sourceGrid) {
		t.Errorf("Source grid shouldn't be equal to new grid after a change was made")
	}
}

func TestGetNewGridShouldReturnAnEmptyGrid(t *testing.T) {
	var newGrid = grid.GetNewGrid(4)
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{0, 0, 0, 0}
	referenceGrid[1] = []int{0, 0, 0, 0}
	referenceGrid[2] = []int{0, 0, 0, 0}
	referenceGrid[3] = []int{0, 0, 0, 0}
	if !reflect.DeepEqual(newGrid, referenceGrid) {
		t.Errorf("Grid should be empty at first move")
	}
}

func TestGetPiecesRawShouldReturnAnArrayEqualToThirdRaw(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	var list = grid.GetPiecesRaw(3, 2, referenceGrid)
	var referenceList = []int{8, 7, 6, 5}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Returned list should be equal to grid thrid raw")
	}
}

func TestGetPiecesColumnShouldReturnAnArrayEqualToFourthColumn(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	var list = grid.GetPiecesColumn(3, 2, referenceGrid)
	var referenceList = []int{4, 12, 5, 13}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Returned list should be equal to grid fourth column")
	}
}

func TestGetPiecesSlashDiagShouldReturnAnArrayEqualToDiag(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	var list = grid.GetPiecesSlashDiag(3, 3, referenceGrid)
	var referenceList = []int{1, 10, 6, 13}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Returned list should be equal to grid diag")
	}
}

func TestGetPiecesBackSlashDiagShouldReturnAnEmptyArrayEqualToDiag(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	var list = grid.GetPiecesBackSlashDiag(3, 0, referenceGrid)
	var referenceList = []int{4, 11, 7, 16}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Returned list should be equal to grid diag")
	}
}

func TestGetPiecesSlashDiagShouldReturnAnEmptyArray(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	var list = grid.GetPiecesSlashDiag(3, 1, referenceGrid)
	var referenceList = []int{}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Returned list should be empty")
	}
}

func TestGetPiecesBackSlashDiagShouldReturnAnEmptyArray(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	var list = grid.GetPiecesBackSlashDiag(3, 1, referenceGrid)
	var referenceList = []int{}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Returned list should be empty")
	}
}

func TestIsWinningLineShouldReturnFalseWithPieces1And4And5(t *testing.T) {
	if grid.IsWinningLine([]int{1, 4, 5}) {
		t.Errorf("List of pieces shouldn't be a winning line")
	}
}

func TestIsWinningLineShouldReturnTrueWithPieces1And3And5(t *testing.T) {
	if !grid.IsWinningLine([]int{1, 3, 5}) {
		t.Errorf("List of pieces should be a winning line")
	}
}

func TestIsWinningLineShouldReturnFalseWithPieces1And4And5And10(t *testing.T) {
	if grid.IsWinningLine([]int{1, 4, 5, 10}) {
		t.Errorf("List of pieces shouldn't be a winning line")
	}
}

func TestIsWinningLineShouldReturnTrueWithPieces1And4And5And15(t *testing.T) {
	if !grid.IsWinningLine([]int{1, 3, 5, 15}) {
		t.Errorf("List of pieces should be a winning line")
	}
}

func TestGetEmptyBoxesShouldReturnOnlyOneCoord(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 0}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	list := grid.GetEmptyBoxes(referenceGrid)
	if list[0] != [2]int{1, 3} {
		t.Errorf("List of empty boxes should contain [1,3]")
	}
}

func TestGetEmptyBoxesShouldReturnEmptyList(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 12}
	referenceGrid[2] = []int{8, 7, 6, 5}
	referenceGrid[3] = []int{16, 15, 14, 13}
	list := grid.GetEmptyBoxes(referenceGrid)
	if len(list) != 0 {
		t.Errorf("List of empty boxes should be empty")
	}
}
