package bitboard

import (
	"fmt"
)

const (
	int32lowest  int32 = -2147483647
	int32highest int32 = 2147483647
)

var numberToBits [8][256][]uint8

func combinations(board ChessBoard, depth int) int {
	if depth == 0 {
		return 1
	}

	moves := board.PsudoLegalMoves(false)
	num := 0
	for _, m := range moves {
		temp := board
		board.DoMove(m)
		if !board.CheckForCheck(!board.BlacksTurn) {
			board.BlacksTurn = !board.BlacksTurn
			add := combinations(board, depth-1)
			num += add
		}
		board = temp
	}
	return num
}

func negaMax(board *ChessBoard, depth int, alpha, beta int32) int32 {
	if depth == 0 {
		return quiscence(board, alpha, beta)
	}

	bestScore := int32lowest
	moves := board.PsudoLegalMoves(false)
	moved := false
	for _, m := range moves {
		temp := *board
		board.DoMove(m)
		if !board.CheckForCheck(!board.BlacksTurn) {
			moved = true
			board.BlacksTurn = !board.BlacksTurn
			score := -negaMax(board, depth-1, -beta, -alpha)
			if score >= beta {
				return score
			}
			if score > bestScore {
				bestScore = score
				if score > alpha {
					alpha = score
				}
			}
		}
		*board = temp
	}

	if !moved {
		if !board.CheckForCheck(!board.BlacksTurn) {
			return 0
		}
	}

	return bestScore
}

func quiscence(board *ChessBoard, alpha, beta int32) int32 {
	standPat := evaluate(board)
	if standPat >= beta {
		return beta
	}
	if alpha < standPat {
		alpha = standPat
	}
	moves := board.PsudoLegalMoves(true)
	for _, m := range moves {
		temp := *board
		board.DoMove(m)
		if !board.CheckForCheck(!board.BlacksTurn) {
			board.BlacksTurn = !board.BlacksTurn
			score := -quiscence(board, -beta, -alpha)
			if score >= beta {
				return beta
			}
			if score > alpha {
				alpha = score
			}
		}
		*board = temp
	}
	return alpha
}

func Search(board *ChessBoard, depth int) Move {
	bestScore := int32lowest
	var bestMove Move

	moves := board.PsudoLegalMoves(false)
	for _, m := range moves {
		temp := *board
		board.DoMove(m)
		if !board.CheckForCheck(!board.BlacksTurn) {
			board.BlacksTurn = !board.BlacksTurn
			score := -negaMax(board, depth-1, int32lowest, int32highest)
			if score >= bestScore {
				bestScore = score
				bestMove = m
			}
		}
		*board = temp
	}
	fmt.Println(bestScore)
	return bestMove
}
