package bitboard

import (
	"fmt"
	"time"
)

const (
	int32lowest  int32 = -2147483647
	int32highest int32 = 2147483647
)

var (
	timeLeft int
)

var numberToBits [8][256][]uint8

func combinations(board *ChessBoard, depth int) int {
	if depth == 0 {
		return 1
	}

	moves := board.PsudoLegalMoves(false)
	num := 0
	for _, m := range moves {
		temp := *board
		board.DoMove(m)
		if !board.CheckForCheck(board.BlacksTurn) {
			add := combinations(board, depth-1)
			num += add
		}
		*board = temp
	}
	return num
}

func negaMax(board *ChessBoard, depth uint8, alpha, beta int32, age uint8) int32 {
	if timeLeft <= 0 {
		return 0
	}
	if depth == 0 {
		return quiscence(board, alpha, beta)
	}

	repetitions := 0
	for _, hash := range board.LastHashes {
		if hash == board.Zobrist {
			repetitions++
			if repetitions == 3 {
				return 0
			}
		}
	}

	tranBestMoveIndex, tranDepth, tranScore, tranNode, tranAge, tranMatching := GetEntry(board.Zobrist)
	if tranNode != 0 && tranDepth >= depth && tranMatching {
		if tranNode == ExactNode {
			return tranScore
		}
		if tranNode == UpperBoundNode && tranScore <= alpha {
			return tranScore
		}
		if tranNode == LowerBoundNode && tranScore >= beta {
			return tranScore
		}
	}

	var bestMoveIndex uint8
	bestScore := int32lowest
	moves := board.PsudoLegalMoves(false)
	moved := false

	node := UpperBoundNode
	save := false
	if tranDepth > depth || tranNode == 0 || tranAge < age {
		save = true
	}

	usedTran := false
	if tranMatching && int(tranBestMoveIndex) < len(moves) {
		usedTran = true
		temp := *board
		board.DoMove(moves[tranBestMoveIndex])
		if !board.CheckForCheck(board.BlacksTurn) {
			moved = true
			score := -negaMax(board, depth-1, -beta, -alpha, age)
			if score >= beta {
				if save {
					StoreEntry(board.Zobrist, bestMoveIndex, depth, bestScore, LowerBoundNode, age)
				}
				return score
			}
			if score > bestScore {
				bestScore = score
				bestMoveIndex = tranBestMoveIndex
				if score > alpha {
					alpha = score
					node = ExactNode
				}
			}
		}
		*board = temp
	}

	for i, m := range moves {
		if i != int(tranBestMoveIndex) || !usedTran {
			temp := *board
			board.DoMove(m)
			if !board.CheckForCheck(board.BlacksTurn) {
				moved = true
				score := -negaMax(board, depth-1, -beta, -alpha, age)
				if score >= beta {
					if save {
						StoreEntry(board.Zobrist, bestMoveIndex, depth, bestScore, LowerBoundNode, age)
					}
					return score
				}
				if score > bestScore {
					bestScore = score
					bestMoveIndex = uint8(i)
					if score > alpha {
						alpha = score
						node = ExactNode
					}
				}
			}
			*board = temp
		}
	}
	if save {
		StoreEntry(board.Zobrist, uint8(bestMoveIndex), depth, bestScore, node, age)
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
		if !board.CheckForCheck(board.BlacksTurn) {
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

func Search(board *ChessBoard, depth uint8, age uint8) Move {
	bestScore := int32lowest
	var bestMove Move

	moves := board.PsudoLegalMoves(false)
	for _, m := range moves {
		temp := *board
		board.DoMove(m)
		if !board.CheckForCheck(board.BlacksTurn) {
			score := -negaMax(board, depth-1, int32lowest, int32highest, age)
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

func IterativeDeepening(board *ChessBoard) Move {
	timeLeft = 1000
	go timer()
	var bestMove Move
	for depth := 1; timeLeft > 0; depth++ {
		newMove := Search(board, uint8(depth), 0)
		if timeLeft > 0 {
			bestMove = newMove
		}
	}
	return bestMove
}

func timer() {
	for timeLeft > 0 {
		time.Sleep(time.Millisecond * 10)
		timeLeft -= 10
	}
}
