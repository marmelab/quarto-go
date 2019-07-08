package grid

import (
	"github.com/ahl5esoft/golang-underscore"
)

// GetListBoxAMinusListBoxB generate a new list of elements of listBoxA less elements of listBoxB
func GetListBoxAMinusListBoxB(listBoxA []Point, listBoxB []Point) []Point {
	minusList := underscore.Select(listBoxA, func(ni Point, _ int) bool {
		var indexBox = underscore.FindIndex(listBoxB, func(nw Point, _ int) bool {
			return ni.X == nw.X && ni.Y == nw.Y
		})
		return indexBox < 0
	})

	if minusList == nil {
		return []Point{}
	}

	return minusList.([]Point)
}

// GetListBoxMinusPoint generate a new list of elements of listBox less without point
func GetListBoxMinusPoint(listBox []Point, point Point) []Point {
	minusList := underscore.Select(listBox, func(ni Point, _ int) bool {
		return ni.X != point.X || ni.Y != point.Y
	})

	if minusList == nil {
		return []Point{}
	}

	return minusList.([]Point)
}


// GetListBoxAMinusListBoxB generate a new list of elements of listPieceA less elements of listPieceB
func GetListPieceAMinusListPieceB(listPieceA []int, listPieceB []int) []int {
	minusList := underscore.Select(listPieceA, func(ni int, _ int) bool {
		var indexBox = underscore.FindIndex(listPieceB, func(nw int, _ int) bool {
			return ni == nw
		})
		return indexBox < 0
	})

	if minusList == nil {
		return []int{}
	}

	return minusList.([]int)
}

// GetListPieceMinusPiece generate a new list of elements of listPiece less without piece
func GetListPieceMinusPiece(listPiece []int, piece int) []int {
	minusList := underscore.Select(listPiece, func(ni int, _ int) bool {
		return ni != piece
	})

	if minusList == nil {
		return []int{}
	}

	return minusList.([]int)
}
