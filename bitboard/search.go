package bitboard

func bitboardToSlice(board Bitboard) []int {
	bits := []int{}
	for i := 0; i < 64; i++ {
		if (1<<i)&board > 0 {
			bits = append(bits, i)
		}
	}
	return bits
}

type move struct {
	from  Bitboard
	to    Bitboard
	piece PieceType

	longCastle  bool
	shortCastle bool

	pawnPromotionPiece PieceType
	enPassant          bool
}

func psudoLegalMoves(board *ChessBoard) []move {

	psudomoves := []move{}
	if !board.BlacksTurn {
		for _, from := range bitboardToSlice(board.WhitePawns) {
			moves := whitePawnMoves(1<<from, board.AllPieces, board.AllBlackPieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, pawnMoves(1<<from, 1<<to, WhitePawn, false)...)
			}
			left, _, right, _ := whiteEnPassant(1<<from, board.AllPieces, board.BlackPawns, board.LastBlackPawns)
			if left > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<from, 1<<(from+9), WhitePawn, true)...)
			}
			if right > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<from, 1<<(from+7), WhitePawn, true)...)
			}
		}

		for _, from := range bitboardToSlice(board.WhiteRooks) {
			moves := rookMoves(1<<from, board.AllPieces, board.AllWhitePieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: WhiteRook})
			}
		}

		for _, from := range bitboardToSlice(board.WhiteKnights) {
			moves := knightMoves(1<<from, board.AllWhitePieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: WhiteKnight})
			}
		}

		for _, from := range bitboardToSlice(board.WhiteBishops) {
			moves := bishopMoves(1<<from, board.AllPieces, board.AllWhitePieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: WhiteBishop})
			}
		}

		for _, from := range bitboardToSlice(board.WhiteQueens) {
			moves := queenMoves(1<<from, board.AllPieces, board.AllWhitePieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: WhiteQueen})
			}
		}

		moves := kingMoves(board.WhiteKing, board.AllWhitePieces)
		for _, to := range bitboardToSlice(moves) {
			psudomoves = append(psudomoves, move{from: board.WhiteKing, to: 1 << to, piece: WhiteKing})
		}
		attacking := board.BlackAttacking()
		if whiteLongCastle(board.WhiteLongCastle, board.AllPieces, attacking) {
			psudomoves = append(psudomoves, move{from: board.WhiteKing, to: c1, piece: WhiteKing, longCastle: true})
		}
		if whiteShortCastle(board.WhiteShortCastle, board.AllPieces, attacking) {
			psudomoves = append(psudomoves, move{from: board.WhiteKing, to: g1, piece: WhiteKing, shortCastle: true})
		}
	} else {
		for _, from := range bitboardToSlice(board.BlackPawns) {
			moves := blackPawnMoves(1<<from, board.AllPieces, board.AllWhitePieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, pawnMoves(1<<from, 1<<to, BlackPawn, false)...)
			}
			left, _, right, _ := blackEnPassant(1<<from, board.AllPieces, board.WhitePawns, board.LastWhitePawns)
			if left > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<from, 1<<(from-7), BlackPawn, true)...)
			}
			if right > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<from, 1<<(from-9), BlackPawn, true)...)
			}
		}

		for _, from := range bitboardToSlice(board.BlackRooks) {
			moves := rookMoves(1<<from, board.AllPieces, board.AllBlackPieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: BlackRook})
			}
		}

		for _, from := range bitboardToSlice(board.BlackKnights) {
			moves := knightMoves(1<<from, board.AllBlackPieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: BlackKnight})
			}
		}

		for _, from := range bitboardToSlice(board.BlackBishops) {
			moves := bishopMoves(1<<from, board.AllPieces, board.AllBlackPieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: BlackBishop})
			}
		}

		for _, from := range bitboardToSlice(board.BlackQueens) {
			moves := queenMoves(1<<from, board.AllPieces, board.AllBlackPieces)
			for _, to := range bitboardToSlice(moves) {
				psudomoves = append(psudomoves, move{from: 1 << from, to: 1 << to, piece: BlackQueen})
			}
		}

		moves := kingMoves(board.BlackKing, board.AllBlackPieces)
		for _, to := range bitboardToSlice(moves) {
			psudomoves = append(psudomoves, move{from: board.BlackKing, to: 1 << to, piece: BlackKing})
		}
		attacking := board.WhiteAttacking()
		if blackLongCastle(board.BlackLongCastle, board.AllPieces, attacking) {
			psudomoves = append(psudomoves, move{from: board.BlackKing, to: c8, piece: BlackKing, longCastle: true})
		}
		if blackShortCastle(board.BlackShortCastle, board.AllPieces, attacking) {
			psudomoves = append(psudomoves, move{from: board.BlackKing, to: g8, piece: BlackKing, shortCastle: true})
		}
	}
	return psudomoves
}

func pawnMoves(from, to Bitboard, piece PieceType, enPassant bool) []move {
	if piece == WhitePawn && to&maskRank[rank8] > 0 {
		return []move{
			{from: from, to: to, piece: WhitePawn, enPassant: enPassant, pawnPromotionPiece: WhiteQueen},
			{from: from, to: to, piece: WhitePawn, enPassant: enPassant, pawnPromotionPiece: WhiteRook},
			{from: from, to: to, piece: WhitePawn, enPassant: enPassant, pawnPromotionPiece: WhiteBishop},
			{from: from, to: to, piece: WhitePawn, enPassant: enPassant, pawnPromotionPiece: WhiteKnight},
		}
	}
	if piece == BlackPawn && to&maskRank[rank1] > 0 {
		return []move{
			{from: from, to: to, piece: BlackPawn, enPassant: enPassant, pawnPromotionPiece: BlackQueen},
			{from: from, to: to, piece: BlackPawn, enPassant: enPassant, pawnPromotionPiece: BlackRook},
			{from: from, to: to, piece: BlackPawn, enPassant: enPassant, pawnPromotionPiece: BlackBishop},
			{from: from, to: to, piece: BlackPawn, enPassant: enPassant, pawnPromotionPiece: BlackKnight},
		}
	}
	return []move{{from: from, to: to, piece: piece, enPassant: enPassant}}
}

func DoMove(board *ChessBoard, m move) {
	board.LastWhitePawns = board.WhitePawns
	board.LastBlackPawns = board.BlackPawns
	if m.piece == WhitePawn {
		if m.enPassant {
			board.DeleteOnSquare(m.to >> 8)
		} else {
			board.DeleteOnSquare(m.to)
		}
		board.WhitePawns = (board.WhitePawns & ^m.from) | m.to
		if m.pawnPromotionPiece != 0 {
			board.PromotePawn(m.to, m.pawnPromotionPiece)
		}
	} else if m.piece == WhiteRook {
		board.DeleteOnSquare(m.to)
		board.WhiteRooks = (board.WhiteRooks & ^m.from) | m.to
	} else if m.piece == WhiteKnight {
		board.DeleteOnSquare(m.to)
		board.WhiteKnights = (board.WhiteKnights & ^m.from) | m.to
	} else if m.piece == WhiteBishop {
		board.DeleteOnSquare(m.to)
		board.WhiteBishops = (board.WhiteBishops & ^m.from) | m.to
	} else if m.piece == WhiteQueen {
		board.DeleteOnSquare(m.to)
		board.WhiteQueens = (board.WhiteQueens & ^m.from) | m.to
	} else if m.piece == WhiteKing {
		board.DeleteOnSquare(m.to)
		board.WhiteKing = (board.WhiteKing & ^m.from) | m.to
		if m.longCastle {
			board.WhiteRooks = (board.WhiteRooks & ^a1) | d1
		} else if m.shortCastle {
			board.WhiteRooks = (board.WhiteRooks & ^h1) | f1
		}
		board.WhiteLongCastle = false
		board.WhiteShortCastle = false
	}

	if m.piece == BlackPawn {
		if m.enPassant {
			board.DeleteOnSquare(m.to << 8)
		} else {
			board.DeleteOnSquare(m.to)
		}
		board.BlackPawns = (board.BlackPawns & ^m.from) | m.to
		if m.pawnPromotionPiece != 0 {
			board.PromotePawn(m.to, m.pawnPromotionPiece)
		}
	} else if m.piece == BlackRook {
		board.DeleteOnSquare(m.to)
		board.BlackRooks = (board.BlackRooks & ^m.from) | m.to
	} else if m.piece == BlackKnight {
		board.DeleteOnSquare(m.to)
		board.BlackKnights = (board.BlackKnights & ^m.from) | m.to
	} else if m.piece == BlackBishop {
		board.DeleteOnSquare(m.to)
		board.BlackBishops = (board.BlackBishops & ^m.from) | m.to
	} else if m.piece == BlackQueen {
		board.DeleteOnSquare(m.to)
		board.BlackQueens = (board.BlackQueens & ^m.from) | m.to
	} else if m.piece == BlackKing {
		board.DeleteOnSquare(m.to)
		board.BlackKing = (board.BlackKing & ^m.from) | m.to
		if m.longCastle {
			board.BlackRooks = (board.BlackRooks & ^a8) | d8
		} else if m.shortCastle {
			board.BlackRooks = (board.BlackRooks & ^h8) | f8
		}
		board.BlackLongCastle = false
		board.BlackShortCastle = false
	}

	if board.WhiteRooks&a1 == 0 {
		board.WhiteLongCastle = false
	}
	if board.WhiteRooks&h1 == 0 {
		board.WhiteShortCastle = false
	}
	if board.BlackRooks&a8 == 0 {
		board.BlackLongCastle = false
	}
	if board.BlackRooks&h8 == 0 {
		board.BlackShortCastle = false
	}

	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces
}

func combinations(board ChessBoard, depth int) int {
	if depth == 0 {
		return 1
	}

	moves := psudoLegalMoves(&board)
	num := 0
	for _, m := range moves {
		temp := board
		DoMove(&board, m)
		if !board.CheckForCheck(!board.BlacksTurn) {
			board.BlacksTurn = !board.BlacksTurn
			add := combinations(board, depth-1)
			num += add
		}
		board = temp
	}
	return num
}

func NegaMax(board ChessBoard, depth int, root bool) (int32, move) {
	if depth == 0 {
		return evaluate(board), move{}
	}
	max := int32(-2147483648)
	var bestMove move

	moves := psudoLegalMoves(&board)
	for _, m := range moves {
		temp := board
		DoMove(&board, m)
		if !board.CheckForCheck(!board.BlacksTurn) {
			board.BlacksTurn = !board.BlacksTurn
			score, _ := NegaMax(board, depth-1, false)
			if -score > max {
				max = -score
				bestMove = m
			}
		}
		board = temp
	}
	return max, bestMove
}

func evaluate(board ChessBoard) int32 {
	return 0
}
