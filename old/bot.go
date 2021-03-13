package main

import (
	"math/rand"
)

func randomMove(board chessBoard) chessBoard {
	var piece int
	var moves []int = []int{}
	for len(moves) == 0 {
		piece = rand.Intn(len(board.blackPieces))
		moves = legalMoves(board, board.blackPieces[piece], true)
	}
	move := moves[rand.Intn(len(moves))]
	board, _ = movePiece(board, piece, black, move, false)

	return board
}
