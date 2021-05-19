package bitboard

import (
	"fmt"
	"runtime"
	"time"
)

const (
	int32lowest  int32         = -2147483647
	int32highest int32         = 2147483647
	timeWait     time.Duration = 60_000
)

var (
	timeLeft bool
)

type response struct {
	move  Move
	score int32
}

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
	if !timeLeft {
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
	if tranNode != 0 && tranDepth >= depth && tranMatching && repetitions <= 1 {
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
	return bestMove
}

func SearchMultiProcessing(board *ChessBoard, depth uint8, age uint8) Move {
	runtime.GOMAXPROCS(runtime.NumCPU())

	bestScore := int32lowest
	var bestMove Move

	moves := board.PsudoLegalMoves(false)

	each := len(moves) / runtime.NumCPU()
	extra := len(moves) % runtime.NumCPU()
	add := 0
	channel := make(chan response)

	for i := 0; i < runtime.NumCPU(); i++ {
		var m []Move
		if extra == 0 {
			m = moves[i*each+add : i*each+add+each]
		} else {
			m = moves[i*each+add : i*each+add+each+1]
			extra--
			add++
		}
		boardCopy := *board
		boardCopy.Init()

		go SearchProcess(m, channel, &boardCopy, depth, age)
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		resp := <-channel
		if resp.score >= bestScore {
			bestScore = resp.score
			bestMove = resp.move
		}
	}

	return bestMove
}

func SearchProcess(moves []Move, channel chan response, board *ChessBoard, depth uint8, age uint8) {
	bestScore := int32lowest
	var bestMove Move

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
	channel <- response{bestMove, bestScore}
}

func IterativeDeepening(board *ChessBoard) Move {
	timeLeft = true
	go timer()
	var bestMove Move
	for depth := 1; timeLeft; depth++ {
		newMove := SearchMultiProcessing(board, uint8(depth), 0)
		if timeLeft {
			bestMove = newMove
		} else {
			fmt.Printf("Depth: %v\n", depth-1)
		}
	}
	return bestMove
}

func timer() {
	time.Sleep(timeWait * time.Millisecond)
	timeLeft = false
}
