package tests

import (
	"quarto/grid"
	"testing"
	"reflect"
	"strconv"
)

func TestGetListBoxAMinusListBoxBShouldReturnReferenceList(t *testing.T) {
	var listA = []grid.Point{}
	listA = append(listA, grid.Point{1, 0})
	listA = append(listA, grid.Point{1, 1})
	listA = append(listA, grid.Point{1, 2})
	listA = append(listA, grid.Point{3, 0})
	listA = append(listA, grid.Point{3, 1})
	listA = append(listA, grid.Point{3, 2})

	var listB = []grid.Point{}
	listB = append(listB, grid.Point{1, 0})
	listB = append(listB, grid.Point{3, 0})
	listB = append(listB, grid.Point{4, 0})

	var listResult = []grid.Point{}
	listResult = append(listResult, grid.Point{1, 1})
	listResult = append(listResult, grid.Point{1, 2})
	listResult = append(listResult, grid.Point{3, 1})
	listResult = append(listResult, grid.Point{3, 2})

	var testList = grid.GetListBoxAMinusListBoxB(listA, listB)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return box list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListBoxAMinusEmptyListBoxBShouldReturnSameList(t *testing.T) {
	var listA = []grid.Point{}
	listA = append(listA, grid.Point{1, 0})
	listA = append(listA, grid.Point{1, 1})
	listA = append(listA, grid.Point{1, 2})
	listA = append(listA, grid.Point{3, 0})
	listA = append(listA, grid.Point{3, 1})
	listA = append(listA, grid.Point{3, 2})

	var listB = []grid.Point{}

	var listResult = []grid.Point{}
	listResult = append(listResult, grid.Point{1, 0})
	listResult = append(listResult, grid.Point{1, 1})
	listResult = append(listResult, grid.Point{1, 2})
	listResult = append(listResult, grid.Point{3, 0})
	listResult = append(listResult, grid.Point{3, 1})
	listResult = append(listResult, grid.Point{3, 2})

	var testList = grid.GetListBoxAMinusListBoxB(listA, listB)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return box list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetEmptyListBoxAMinusListBoxBShouldEmptyList(t *testing.T) {
	var listA = []grid.Point{}
	listA = append(listA, grid.Point{1, 0})
	listA = append(listA, grid.Point{1, 1})
	listA = append(listA, grid.Point{1, 2})
	listA = append(listA, grid.Point{3, 0})
	listA = append(listA, grid.Point{3, 1})
	listA = append(listA, grid.Point{3, 2})

	var listB = []grid.Point{}

	var listResult = []grid.Point{}

	var testList = grid.GetListBoxAMinusListBoxB(listB, listA)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return box list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListPieceAMinusListPieceBShouldReturnReferenceList(t *testing.T) {
	var listA = []int{}
	listA = append(listA, 1)
	listA = append(listA, 2)
	listA = append(listA, 3)
	listA = append(listA, 4)
	listA = append(listA, 5)
	listA = append(listA, 6)

	var listB = []int{}
	listB = append(listB, 3)
	listB = append(listB, 5)
	listB = append(listB, 10)

	var listResult = []int{}
	listResult = append(listResult, 1)
	listResult = append(listResult, 2)
	listResult = append(listResult, 4)
	listResult = append(listResult, 6)

	var testList = grid.GetListPieceAMinusListPieceB(listA, listB)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return piece list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListPieceAMinusEmptyListPieceBShouldReturnSameList(t *testing.T) {
	var listA = []int{}
	listA = append(listA, 1)
	listA = append(listA, 2)
	listA = append(listA, 3)
	listA = append(listA, 4)
	listA = append(listA, 5)
	listA = append(listA, 6)

	var listB = []int{}

	var listResult = []int{}
	listResult = append(listResult, 1)
	listResult = append(listResult, 2)
	listResult = append(listResult, 3)
	listResult = append(listResult, 4)
	listResult = append(listResult, 5)
	listResult = append(listResult, 6)

	var testList = grid.GetListPieceAMinusListPieceB(listA, listB)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return piece list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetEmptyListPieceAMinusListPieceBShouldEmptyList(t *testing.T) {
	var listA = []int{}
	listA = append(listA, 1)
	listA = append(listA, 2)
	listA = append(listA, 3)
	listA = append(listA, 4)
	listA = append(listA, 5)
	listA = append(listA, 6)

	var listB = []int{}

	var listResult = []int{}

	var testList = grid.GetListPieceAMinusListPieceB(listB, listA)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return piece list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListBoxMinusPointShouldReturnReferenceList(t *testing.T) {
	var listA = []grid.Point{}
	listA = append(listA, grid.Point{1, 0})
	listA = append(listA, grid.Point{1, 1})
	listA = append(listA, grid.Point{1, 2})
	listA = append(listA, grid.Point{3, 0})
	listA = append(listA, grid.Point{3, 1})
	listA = append(listA, grid.Point{3, 2})

	var point = grid.Point{1, 0}

	var listResult = []grid.Point{}
	listResult = append(listResult, grid.Point{1, 1})
	listResult = append(listResult, grid.Point{1, 2})
	listResult = append(listResult, grid.Point{3, 0})
	listResult = append(listResult, grid.Point{3, 1})
	listResult = append(listResult, grid.Point{3, 2})

	var testList = grid.GetListBoxMinusPoint(listA, point)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return box list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListBoxMinusPointShouldReturnSameList(t *testing.T) {
	var listA = []grid.Point{}
	listA = append(listA, grid.Point{1, 0})
	listA = append(listA, grid.Point{1, 1})
	listA = append(listA, grid.Point{1, 2})
	listA = append(listA, grid.Point{3, 0})
	listA = append(listA, grid.Point{3, 1})
	listA = append(listA, grid.Point{3, 2})

	var point = grid.Point{4, 0}

	var listResult = []grid.Point{}
	listResult = append(listResult, grid.Point{1, 0})
	listResult = append(listResult, grid.Point{1, 1})
	listResult = append(listResult, grid.Point{1, 2})
	listResult = append(listResult, grid.Point{3, 0})
	listResult = append(listResult, grid.Point{3, 1})
	listResult = append(listResult, grid.Point{3, 2})

	var testList = grid.GetListBoxMinusPoint(listA, point)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return box list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListPieceMinusPieceShouldReturnReferenceList(t *testing.T) {
	var listA = []int{}
	listA = append(listA, 1)
	listA = append(listA, 2)
	listA = append(listA, 3)
	listA = append(listA, 4)
	listA = append(listA, 5)
	listA = append(listA, 6)

	var piece = 4

	var listResult = []int{}
	listResult = append(listResult, 1)
	listResult = append(listResult, 2)
	listResult = append(listResult, 3)
	listResult = append(listResult, 5)
	listResult = append(listResult, 6)

	var testList = grid.GetListPieceMinusPiece(listA, piece)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return piece list (" + strconv.Itoa(len(testList)) + ")")
	}
}

func TestGetListPieceMinusPieceShouldReturnSameList(t *testing.T) {
	var listA = []int{}
	listA = append(listA, 1)
	listA = append(listA, 2)
	listA = append(listA, 3)
	listA = append(listA, 4)
	listA = append(listA, 5)
	listA = append(listA, 6)

	var piece = 11

	var listResult = []int{}
	listResult = append(listResult, 1)
	listResult = append(listResult, 2)
	listResult = append(listResult, 3)
	listResult = append(listResult, 4)
	listResult = append(listResult, 5)
	listResult = append(listResult, 6)

	var testList = grid.GetListPieceMinusPiece(listA, piece)
	
	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Bad return piece list (" + strconv.Itoa(len(testList)) + ")")
	}
}