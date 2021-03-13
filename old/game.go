package main

type indexAndY struct {
	index int
	y     int
}

type chessBoard struct {
	board       [64]uint8
	blackPieces []int
	whitePieces []int
}

var kingPositions [8]indexAndY = [8]indexAndY{{-9, -1}, {-8, -1}, {-7, -1}, {-1, 0}, {1, 0}, {7, 1}, {8, 1}, {9, 1}}
var knightPositions [8]indexAndY = [8]indexAndY{{-17, -2}, {-15, -2}, {-10, -1}, {-6, -1}, {6, 1}, {10, 1}, {15, 2}, {17, 2}}
var diagonalPositions [4]indexAndY = [4]indexAndY{{-9, -1}, {-7, -1}, {7, 1}, {9, 1}}
var straightPositions [4]indexAndY = [4]indexAndY{{-8, -1}, {-1, 0}, {1, 0}, {8, 1}}

func legalMoves(board chessBoard, pos int, check bool) []int {
	piece := board.board[pos]
	var onlyPiece uint8
	var color uint8
	if piece&white > 0 {
		onlyPiece = piece ^ white
		color = white
	} else if piece&black > 0 {
		onlyPiece = piece ^ black
		color = black
	} else {
		return []int{}
	}
	moves := []int{}
	if onlyPiece == king {
		for _, kingPos := range kingPositions {
			newPos := pos + kingPos.index
			if newPos < 64 && newPos >= 0 && pos/8+kingPos.y == newPos/8 {
				if board.board[newPos]&color == 0 {
					if check {
						if checkForCheck(board, color, pos, newPos) {
							moves = append(moves, newPos)
						}
					} else {
						moves = append(moves, newPos)
					}
				}
			}
		}
	}
	if onlyPiece == knight {
		for _, knightPos := range knightPositions {
			newPos := pos + knightPos.index
			if newPos < 64 && newPos >= 0 && pos/8+knightPos.y == newPos/8 {
				if board.board[newPos]&color == 0 {
					if check {
						if checkForCheck(board, color, pos, newPos) {
							moves = append(moves, newPos)
						}
					} else {
						moves = append(moves, newPos)
					}
				}
			}
		}
	}
	if onlyPiece == queen || onlyPiece == bishop {
		for _, dir := range diagonalPositions {
			oldPos := pos
			oldY := pos / 8
			for {
				newPos := oldPos + dir.index
				if newPos < 64 && newPos >= 0 && oldY+dir.y == newPos/8 {
					if board.board[newPos]&color == 0 {
						if check {
							if checkForCheck(board, color, pos, newPos) {
								moves = append(moves, newPos)
							}
						} else {
							moves = append(moves, newPos)
						}
					}
					if board.board[newPos] != 0 {
						break
					}
				} else {
					break
				}
				oldPos = newPos
				oldY = newPos / 8
			}
		}
	}
	if onlyPiece == queen || onlyPiece == rook {
		for _, dir := range straightPositions {
			oldPos := pos
			oldY := pos / 8
			for {
				newPos := oldPos + dir.index
				if newPos < 64 && newPos >= 0 && oldY+dir.y == newPos/8 {
					if board.board[newPos]&color == 0 {
						if check {
							if checkForCheck(board, color, pos, newPos) {
								moves = append(moves, newPos)
							}
						} else {
							moves = append(moves, newPos)
						}
					}
					if board.board[newPos] != 0 {
						break
					}
				} else {
					break
				}
				oldPos = newPos
				oldY = newPos / 8
			}
		}
	}
	if onlyPiece == pawn {
		if color == white {
			if pos/8 == 6 {
				if board.board[pos-8] == 0 && board.board[pos-16] == 0 {

					if check {
						if checkForCheck(board, color, pos, pos-16) {
							moves = append(moves, pos-16)
						}
					} else {
						moves = append(moves, pos-16)
					}
				}
			}
			if pos/8 != 0 {
				if board.board[pos-8] == 0 {
					if check {
						if checkForCheck(board, color, pos, pos-8) {
							moves = append(moves, pos-8)
						}
					} else {
						moves = append(moves, pos-8)
					}
				}
				if pos-9 >= 0 && (pos-9)/8 == pos/8-1 {
					if board.board[pos-9]&black > 0 {
						if check {
							if checkForCheck(board, color, pos, pos-9) {
								moves = append(moves, pos-9)
							}
						} else {
							moves = append(moves, pos-9)
						}
					}
				}
				if board.board[pos-7]&black > 0 && (pos-7)/8 == pos/8-1 {
					if check {
						if checkForCheck(board, color, pos, pos-7) {
							moves = append(moves, pos-7)
						}
					} else {
						moves = append(moves, pos-7)
					}
				}
			}
		} else {
			if pos/8 == 1 {
				if board.board[pos+8] == 0 && board.board[pos+16] == 0 {
					if check {
						if checkForCheck(board, color, pos, pos+16) {
							moves = append(moves, pos+16)
						}
					} else {
						moves = append(moves, pos+16)
					}
				}
			}
			if pos/8 != 7 {
				if board.board[pos+8] == 0 {
					if check {
						if checkForCheck(board, color, pos, pos+8) {
							moves = append(moves, pos+8)
						}
					} else {
						moves = append(moves, pos+8)
					}
				}
				if pos+9 < 64 && (pos+9)/8 == pos/8+1 {
					if board.board[pos+9]&white > 0 {
						if check {
							if checkForCheck(board, color, pos, pos+9) {
								moves = append(moves, pos+9)
							}
						} else {
							moves = append(moves, pos+9)
						}
					}
				}
				if board.board[pos+7]&white > 0 && (pos+7)/8 == pos/8+1 {
					if check {
						if checkForCheck(board, color, pos, pos+7) {
							moves = append(moves, pos+7)
						}
					} else {
						moves = append(moves, pos+7)
					}
				}
			}
		}
	}
	return moves
}

func movePiece(board chessBoard, index int, color uint8, to int, checkLegal bool) (chessBoard, bool) {
	var legal bool
	var from int
	if color == white {
		from = board.whitePieces[index]
	} else {
		from = board.blackPieces[index]
	}
	if checkLegal {
		moves := legalMoves(board, from, true)
		for _, move := range moves {
			if move == to {
				legal = true
				break
			}
		}
	}

	if legal || !checkLegal {
		if color == white {
			board.whitePieces[index] = to
			if board.board[to]&black > 0 {
				for i, j := range board.blackPieces {
					if j == to {
						board.blackPieces[i] = board.blackPieces[len(board.blackPieces)-1]
						board.blackPieces = board.blackPieces[:len(board.blackPieces)-1]
						break
					}
				}
			}
		} else {
			board.blackPieces[index] = to
			if board.board[to]&white > 0 {
				for i, j := range board.whitePieces {
					if j == to {
						board.whitePieces[i] = board.whitePieces[len(board.whitePieces)-1]
						board.whitePieces = board.whitePieces[:len(board.whitePieces)-1]
						break
					}
				}
			}
		}

		board.board[to] = board.board[from]
		board.board[from] = 0
	}

	return board, legal || !checkLegal
}

func attackedSquares(board chessBoard, color uint8) [64]bool {
	var pieces []int
	if color|white > 0 {
		pieces = board.whitePieces
	} else {
		pieces = board.blackPieces
	}

	var attacked [64]bool
	for _, piece := range pieces {
		moves := legalMoves(board, piece, false)
		for _, move := range moves {
			attacked[move] = true
		}
	}
	return attacked
}

func checkForCheck(board chessBoard, color uint8, from int, to int) bool {
	board.board[to] = board.board[from]
	board.board[from] = 0

	var pieces []int
	var king int

	if color == white {
		pieces = board.blackPieces
		king = board.whitePieces[0]
	} else {
		pieces = board.whitePieces
		king = board.blackPieces[0]
	}

	for _, piece := range pieces {
		moves := legalMoves(board, piece, false)
		for _, move := range moves {
			if move == king {
				return false
			}
		}
	}
	return true
}
