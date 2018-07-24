package tests

import (
	"quarto/grid"
	"reflect"
	"strconv"
	"testing"
)

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
	if (list[0] != grid.Point{3, 1}) {
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

func TestGetAlignedPieceNumberShouldReturn5(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 2, 3, 4}
	referenceGrid[1] = []int{9, 10, 11, 0}
	referenceGrid[2] = []int{8, 0, 6, 5}
	referenceGrid[3] = []int{16, 0, 14, 13}
	value := grid.GetAlignedPieceNumber(referenceGrid, 3, 1)
	if value != 5 {
		t.Errorf("Occupation value should be 5 for coordinate [3, 1] (" + strconv.Itoa(value) + ")")
	}
}

func TestGetAlignedPieceNumberShouldReturn0(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	value := grid.GetAlignedPieceNumber(referenceGrid, 3, 1)
	if value != 0 {
		t.Errorf("Occupation value should be 0 for coordinate [3, 1] (" + strconv.Itoa(value) + ")")
	}
}

func TestBoxFilledNumberShouldReturn3(t *testing.T) {
	piecesList := []int{1, 2, 0, 6}
	value := grid.BoxFilledNumber(piecesList)
	if value != 3 {
		t.Errorf("Number of occupied value should be 3 for this list of Point")
	}
}

func TestBoxFilledNumberShouldReturn0(t *testing.T) {
	piecesList := []int{0, 0, 0, 0}
	value := grid.BoxFilledNumber(piecesList)
	if value != 0 {
		t.Errorf("Number of occupied value should be 0 for this list of Point")
	}
}

func TestBoxFilledNumberShouldReturn4(t *testing.T) {
	piecesList := []int{1, 2, 3, 6}
	value := grid.BoxFilledNumber(piecesList)
	if value != 4 {
		t.Errorf("Number of occupied value should be 4 for this list of Point")
	}
}

func TestGetPositionScoreForPieceShouldReturn3(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 0, 0, 0}
	referenceGrid[1] = []int{0, 0, 0, 0}
	referenceGrid[2] = []int{0, 9, 0, 0}
	referenceGrid[3] = []int{0, 0, 0, 16}
	value := grid.GetPositionScoreForPiece(referenceGrid, 0, 1, 8)
	if value != 3 {
		t.Errorf("position score should be 3 (" + strconv.Itoa(value) + ")")
	}
}

func TestGetPositionScoreForPieceShouldReturn1(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 0, 0, 0}
	referenceGrid[1] = []int{0, 0, 0, 0}
	referenceGrid[2] = []int{0, 9, 0, 0}
	referenceGrid[3] = []int{0, 0, 0, 16}
	value := grid.GetPositionScoreForPiece(referenceGrid, 1, 2, 8)
	if value != 1 {
		t.Errorf("position score should be 1 (" + strconv.Itoa(value) + ")")
	}
}

func TestGetSafestBoxesIncludingPieceChoiceShouldReturnPositionX2Y1(t *testing.T) {
	var referenceGrid = grid.GetNewGrid(4)
	referenceGrid[0] = []int{1, 0, 0, 0}
	referenceGrid[1] = []int{0, 0, 0, 0}
	referenceGrid[2] = []int{0, 9, 0, 0}
	referenceGrid[3] = []int{0, 0, 0, 16}
	boxList := grid.GetSafestBoxesIncludingPieceChoice(referenceGrid, 8)
	referencesBoxList := []grid.Point{grid.Point{2, 1}}
	if !reflect.DeepEqual(boxList, referencesBoxList) {
		t.Errorf("Safest boxes should be [1, 2]")
	}
}
