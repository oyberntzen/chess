package main

import (
	"unicode"
)

var letterToPiece map[rune]uint8 = map[rune]uint8{
	'K': king | white,
	'Q': queen | white,
	'B': bishop | white,
	'N': knight | white,
	'R': rook | white,
	'P': pawn | white,
	'k': king | black,
	'q': queen | black,
	'b': bishop | black,
	'n': knight | black,
	'r': rook | black,
	'p': pawn | black,
}

func fenToBoard(fen string) chessBoard {
	x := 0
	y := 0
	var fenBoard chessBoard
	fenBoard.whitePieces = []int{0}
	fenBoard.blackPieces = []int{0}
	for _, char := range fen {
		if char == '/' {
			x = 0
			y++
		} else if unicode.IsDigit(char) {
			x += int(char)
		} else {
			piece := letterToPiece[char]
			fenBoard.board[y*8+x] = letterToPiece[char]
			if piece&white > 0 {
				if piece == king|white {
					fenBoard.whitePieces[0] = y*8 + x
				} else {
					fenBoard.whitePieces = append(fenBoard.whitePieces, y*8+x)
				}
			} else if piece&black > 0 {
				if piece == king|black {
					fenBoard.blackPieces[0] = y*8 + x
				} else {
					fenBoard.blackPieces = append(fenBoard.blackPieces, y*8+x)
				}
			}
			x++
		}
	}
	return fenBoard
}
