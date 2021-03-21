package bitboard

import (
	"strings"
	"unicode"
)

type Bitboard uint64

type ChessBoard struct {
	WhitePawns   Bitboard
	WhiteRooks   Bitboard
	WhiteKnights Bitboard
	WhiteBishops Bitboard
	WhiteQueens  Bitboard
	WhiteKing    Bitboard

	BlackPawns   Bitboard
	BlackRooks   Bitboard
	BlackKnights Bitboard
	BlackBishops Bitboard
	BlackQueens  Bitboard
	BlackKing    Bitboard

	AllBitboards [12]*Bitboard

	AllWhitePieces Bitboard
	AllBlackPieces Bitboard
	AllPieces      Bitboard

	WhiteShortCastle bool
	WhiteLongCastle  bool
	BlackShortCastle bool
	BlackLongCastle  bool

	LastWhitePawns Bitboard
	LastBlackPawns Bitboard

	BlacksTurn bool
}

type PieceType uint8

const (
	WhitePawn   PieceType = 1
	WhiteRook   PieceType = 2
	WhiteKnight PieceType = 3
	WhiteBishop PieceType = 4
	WhiteQueen  PieceType = 5
	WhiteKing   PieceType = 6
	BlackPawn   PieceType = 7
	BlackRook   PieceType = 8
	BlackKnight PieceType = 9
	BlackBishop PieceType = 10
	BlackQueen  PieceType = 11
	BlackKing   PieceType = 12
)

type Move struct {
	From  Bitboard
	To    Bitboard
	Piece PieceType

	LongCastle  bool
	ShortCastle bool

	PawnPromotionPiece PieceType
	EnPassant          bool
}

var (
	clearRank [8]Bitboard = [8]Bitboard{
		0xffffffffffffff00,
		0xffffffffffff00ff,
		0xffffffffff00ffff,
		0xffffffff00ffffff,
		0xffffff00ffffffff,
		0xffff00ffffffffff,
		0xff00ffffffffffff,
		0x00ffffffffffffff,
	}
	maskRank [8]Bitboard = [8]Bitboard{
		0x00000000000000ff,
		0x000000000000ff00,
		0x0000000000ff0000,
		0x00000000ff000000,
		0x000000ff00000000,
		0x0000ff0000000000,
		0x00ff000000000000,
		0xff00000000000000,
	}
	clearFile [8]Bitboard = [8]Bitboard{
		0x7f7f7f7f7f7f7f7f,
		0xbfbfbfbfbfbfbfbf,
		0xdfdfdfdfdfdfdfdf,
		0xefefefefefefefef,
		0xf7f7f7f7f7f7f7f7,
		0xfbfbfbfbfbfbfbfb,
		0xfdfdfdfdfdfdfdfd,
		0xfefefefefefefefe,
	}
	maskFile [8]Bitboard = [8]Bitboard{
		0x8080808080808080,
		0x4040404040404040,
		0x2020202020202020,
		0x1010101010101010,
		0x0808080808080808,
		0x0404040404040404,
		0x0202020202020202,
		0x0101010101010101,
	}
)

const (
	rank1 int = 0
	rank2 int = 1
	rank3 int = 2
	rank4 int = 3
	rank5 int = 4
	rank6 int = 5
	rank7 int = 6
	rank8 int = 7

	fileA int = 0
	fileB int = 1
	fileC int = 2
	fileD int = 3
	fileE int = 4
	fileF int = 5
	fileG int = 6
	fileH int = 7

	a1 Bitboard = 0x0000000000000080
	b1 Bitboard = 0x0000000000000040
	c1 Bitboard = 0x0000000000000020
	d1 Bitboard = 0x0000000000000010
	e1 Bitboard = 0x0000000000000008
	f1 Bitboard = 0x0000000000000004
	g1 Bitboard = 0x0000000000000002
	h1 Bitboard = 0x0000000000000001
	a2 Bitboard = 0x0000000000008000
	b2 Bitboard = 0x0000000000004000
	c2 Bitboard = 0x0000000000002000
	d2 Bitboard = 0x0000000000001000
	e2 Bitboard = 0x0000000000000800
	f2 Bitboard = 0x0000000000000400
	g2 Bitboard = 0x0000000000000200
	h2 Bitboard = 0x0000000000000100
	a3 Bitboard = 0x0000000000800000
	b3 Bitboard = 0x0000000000400000
	c3 Bitboard = 0x0000000000200000
	d3 Bitboard = 0x0000000000100000
	e3 Bitboard = 0x0000000000080000
	f3 Bitboard = 0x0000000000040000
	g3 Bitboard = 0x0000000000020000
	h3 Bitboard = 0x0000000000010000
	a4 Bitboard = 0x0000000080000000
	b4 Bitboard = 0x0000000040000000
	c4 Bitboard = 0x0000000020000000
	d4 Bitboard = 0x0000000010000000
	e4 Bitboard = 0x0000000008000000
	f4 Bitboard = 0x0000000004000000
	g4 Bitboard = 0x0000000002000000
	h4 Bitboard = 0x0000000001000000
	a5 Bitboard = 0x0000008000000000
	b5 Bitboard = 0x0000004000000000
	c5 Bitboard = 0x0000002000000000
	d5 Bitboard = 0x0000001000000000
	e5 Bitboard = 0x0000000800000000
	f5 Bitboard = 0x0000000400000000
	g5 Bitboard = 0x0000000200000000
	h5 Bitboard = 0x0000000100000000
	a6 Bitboard = 0x0000800000000000
	b6 Bitboard = 0x0000400000000000
	c6 Bitboard = 0x0000200000000000
	d6 Bitboard = 0x0000100000000000
	e6 Bitboard = 0x0000080000000000
	f6 Bitboard = 0x0000040000000000
	g6 Bitboard = 0x0000020000000000
	h6 Bitboard = 0x0000010000000000
	a7 Bitboard = 0x0080000000000000
	b7 Bitboard = 0x0040000000000000
	c7 Bitboard = 0x0020000000000000
	d7 Bitboard = 0x0010000000000000
	e7 Bitboard = 0x0008000000000000
	f7 Bitboard = 0x0004000000000000
	g7 Bitboard = 0x0002000000000000
	h7 Bitboard = 0x0001000000000000
	a8 Bitboard = 0x8000000000000000
	b8 Bitboard = 0x4000000000000000
	c8 Bitboard = 0x2000000000000000
	d8 Bitboard = 0x1000000000000000
	e8 Bitboard = 0x0800000000000000
	f8 Bitboard = 0x0400000000000000
	g8 Bitboard = 0x0200000000000000
	h8 Bitboard = 0x0100000000000000
)

var posToString map[Bitboard]string = map[Bitboard]string{
	a1: "a1",
	a2: "a2",
	a3: "a3",
	a4: "a4",
	a5: "a5",
	a6: "a6",
	a7: "a7",
	a8: "a8",
	b1: "b1",
	b2: "b2",
	b3: "b3",
	b4: "b4",
	b5: "b5",
	b6: "b6",
	b7: "b7",
	b8: "b8",
	c1: "c1",
	c2: "c2",
	c3: "c3",
	c4: "c4",
	c5: "c5",
	c6: "c6",
	c7: "c7",
	c8: "c8",
	d1: "d1",
	d2: "d2",
	d3: "d3",
	d4: "d4",
	d5: "d5",
	d6: "d6",
	d7: "d7",
	d8: "d8",
	e1: "e1",
	e2: "e2",
	e3: "e3",
	e4: "e4",
	e5: "e5",
	e6: "e6",
	e7: "e7",
	e8: "e8",
	f1: "f1",
	f2: "f2",
	f3: "f3",
	f4: "f4",
	f5: "f5",
	f6: "f6",
	f7: "f7",
	f8: "f8",
	g1: "g1",
	g2: "g2",
	g3: "g3",
	g4: "g4",
	g5: "g5",
	g6: "g6",
	g7: "g7",
	g8: "g8",
	h1: "h1",
	h2: "h2",
	h3: "h3",
	h4: "h4",
	h5: "h5",
	h6: "h6",
	h7: "h7",
	h8: "h8",
}

func (board *ChessBoard) PsudoLegalMoves(onlyCaptures bool) []Move {
	andWith := ^Bitboard(0)
	if onlyCaptures {
		if board.BlacksTurn {
			andWith = board.AllWhitePieces
		} else {
			andWith = board.AllBlackPieces
		}
	}

	psudomoves := []Move{}
	if !board.BlacksTurn {
		for _, From := range BitboardToSlice(board.WhitePawns) {
			moves := whitePawnMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<To, WhitePawn, false)...)
			}
			left, _, right, _ := whiteEnPassant(1<<From, board.AllPieces, board.BlackPawns, board.LastBlackPawns)
			if left > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From+9), WhitePawn, true)...)
			}
			if right > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From+7), WhitePawn, true)...)
			}
		}

		for _, From := range BitboardToSlice(board.WhiteRooks) {
			moves := rookMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: WhiteRook})
			}
		}

		for _, From := range BitboardToSlice(board.WhiteKnights) {
			moves := knightMoves(1<<From, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: WhiteKnight})
			}
		}

		for _, From := range BitboardToSlice(board.WhiteBishops) {
			moves := bishopMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: WhiteBishop})
			}
		}

		for _, From := range BitboardToSlice(board.WhiteQueens) {
			moves := queenMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: WhiteQueen})
			}
		}

		moves := kingMoves(board.WhiteKing, board.AllWhitePieces) & andWith
		for _, To := range BitboardToSlice(moves) {
			psudomoves = append(psudomoves, Move{From: board.WhiteKing, To: 1 << To, Piece: WhiteKing})
		}
		attacking := board.BlackAttacking()
		if whiteLongCastle(board.WhiteLongCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.WhiteKing, To: c1, Piece: WhiteKing, LongCastle: true})
		}
		if whiteShortCastle(board.WhiteShortCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.WhiteKing, To: g1, Piece: WhiteKing, ShortCastle: true})
		}
	} else {
		for _, From := range BitboardToSlice(board.BlackPawns) {
			moves := blackPawnMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<To, BlackPawn, false)...)
			}
			left, _, right, _ := blackEnPassant(1<<From, board.AllPieces, board.WhitePawns, board.LastWhitePawns)
			if left > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From-7), BlackPawn, true)...)
			}
			if right > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From-9), BlackPawn, true)...)
			}
		}

		for _, From := range BitboardToSlice(board.BlackRooks) {
			moves := rookMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: BlackRook})
			}
		}

		for _, From := range BitboardToSlice(board.BlackKnights) {
			moves := knightMoves(1<<From, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: BlackKnight})
			}
		}

		for _, From := range BitboardToSlice(board.BlackBishops) {
			moves := bishopMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: BlackBishop})
			}
		}

		for _, From := range BitboardToSlice(board.BlackQueens) {
			moves := queenMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, Piece: BlackQueen})
			}
		}

		moves := kingMoves(board.BlackKing, board.AllBlackPieces) & andWith
		for _, To := range BitboardToSlice(moves) {
			psudomoves = append(psudomoves, Move{From: board.BlackKing, To: 1 << To, Piece: BlackKing})
		}
		attacking := board.WhiteAttacking()
		if blackLongCastle(board.BlackLongCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.BlackKing, To: c8, Piece: BlackKing, LongCastle: true})
		}
		if blackShortCastle(board.BlackShortCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.BlackKing, To: g8, Piece: BlackKing, ShortCastle: true})
		}
	}
	return psudomoves
}

func (board *ChessBoard) DoMove(m Move) {
	board.LastWhitePawns = board.WhitePawns
	board.LastBlackPawns = board.BlackPawns
	if m.Piece == WhitePawn {
		if m.EnPassant {
			board.DeleteOnSquare(m.To >> 8)
		} else {
			board.DeleteOnSquare(m.To)
		}
		board.WhitePawns = (board.WhitePawns & ^m.From) | m.To
		if m.PawnPromotionPiece != 0 {
			board.PromotePawn(m.To, m.PawnPromotionPiece)
		}
	} else if m.Piece == WhiteRook {
		board.DeleteOnSquare(m.To)
		board.WhiteRooks = (board.WhiteRooks & ^m.From) | m.To
	} else if m.Piece == WhiteKnight {
		board.DeleteOnSquare(m.To)
		board.WhiteKnights = (board.WhiteKnights & ^m.From) | m.To
	} else if m.Piece == WhiteBishop {
		board.DeleteOnSquare(m.To)
		board.WhiteBishops = (board.WhiteBishops & ^m.From) | m.To
	} else if m.Piece == WhiteQueen {
		board.DeleteOnSquare(m.To)
		board.WhiteQueens = (board.WhiteQueens & ^m.From) | m.To
	} else if m.Piece == WhiteKing {
		board.DeleteOnSquare(m.To)
		board.WhiteKing = (board.WhiteKing & ^m.From) | m.To
		if m.LongCastle {
			board.WhiteRooks = (board.WhiteRooks & ^a1) | d1
		} else if m.ShortCastle {
			board.WhiteRooks = (board.WhiteRooks & ^h1) | f1
		}
		board.WhiteLongCastle = false
		board.WhiteShortCastle = false
	}

	if m.Piece == BlackPawn {
		if m.EnPassant {
			board.DeleteOnSquare(m.To << 8)
		} else {
			board.DeleteOnSquare(m.To)
		}
		board.BlackPawns = (board.BlackPawns & ^m.From) | m.To
		if m.PawnPromotionPiece != 0 {
			board.PromotePawn(m.To, m.PawnPromotionPiece)
		}
	} else if m.Piece == BlackRook {
		board.DeleteOnSquare(m.To)
		board.BlackRooks = (board.BlackRooks & ^m.From) | m.To
	} else if m.Piece == BlackKnight {
		board.DeleteOnSquare(m.To)
		board.BlackKnights = (board.BlackKnights & ^m.From) | m.To
	} else if m.Piece == BlackBishop {
		board.DeleteOnSquare(m.To)
		board.BlackBishops = (board.BlackBishops & ^m.From) | m.To
	} else if m.Piece == BlackQueen {
		board.DeleteOnSquare(m.To)
		board.BlackQueens = (board.BlackQueens & ^m.From) | m.To
	} else if m.Piece == BlackKing {
		board.DeleteOnSquare(m.To)
		board.BlackKing = (board.BlackKing & ^m.From) | m.To
		if m.LongCastle {
			board.BlackRooks = (board.BlackRooks & ^a8) | d8
		} else if m.ShortCastle {
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

func (board *ChessBoard) MovesOnSquare(square Bitboard) Bitboard {
	if board.WhitePawns&square > 0 {
		moves := whitePawnMoves(square, board.AllPieces, board.AllBlackPieces)
		left, _, right, _ := whiteEnPassant(square, board.AllPieces, board.BlackPawns, board.LastBlackPawns)
		moves |= left | right
		return moves
	} else if board.WhiteRooks&square > 0 {
		return rookMoves(square, board.AllPieces, board.AllWhitePieces)
	} else if board.WhiteKnights&square > 0 {
		return knightMoves(square, board.AllWhitePieces)
	} else if board.WhiteBishops&square > 0 {
		return bishopMoves(square, board.AllPieces, board.AllWhitePieces)
	} else if board.WhiteQueens&square > 0 {
		return queenMoves(square, board.AllPieces, board.AllWhitePieces)
	} else if board.WhiteKing&square > 0 {
		moves := kingMoves(square, board.AllWhitePieces)
		attacking := board.BlackAttacking()
		if whiteLongCastle(board.WhiteLongCastle, board.AllWhitePieces, attacking) {
			moves |= c1
		}
		if whiteShortCastle(board.WhiteShortCastle, board.AllWhitePieces, attacking) {
			moves |= g1
		}
		return moves
	} else if board.BlackPawns&square > 0 {
		moves := blackPawnMoves(square, board.AllPieces, board.AllWhitePieces)
		left, _, right, _ := blackEnPassant(square, board.AllPieces, board.WhitePawns, board.LastWhitePawns)
		moves |= left | right
		return moves
	} else if board.BlackRooks&square > 0 {
		return rookMoves(square, board.AllPieces, board.AllBlackPieces)
	} else if board.BlackKnights&square > 0 {
		return knightMoves(square, board.AllBlackPieces)
	} else if board.BlackBishops&square > 0 {
		return bishopMoves(square, board.AllPieces, board.AllBlackPieces)
	} else if board.BlackQueens&square > 0 {
		return queenMoves(square, board.AllPieces, board.AllBlackPieces)
	} else if board.BlackKing&square > 0 {
		moves := kingMoves(square, board.AllBlackPieces)
		attacking := board.WhiteAttacking()
		if blackLongCastle(board.BlackLongCastle, board.AllBlackPieces, attacking) {
			moves |= c8
		}
		if blackShortCastle(board.BlackShortCastle, board.AllBlackPieces, attacking) {
			moves |= g8
		}
		return moves
	}
	return 0
}

func (board *ChessBoard) MovePiece(From Bitboard, To Bitboard) (bool, bool, Bitboard) {
	promotions := Bitboard(0)
	whitePawn := false
	newLastWhitePawns := Bitboard(0)
	newLastBlackPawns := Bitboard(0)
	if board.WhitePawns&From > 0 && !board.BlacksTurn {
		left, leftTaken, right, rightTaken := whiteEnPassant(From, board.AllPieces, board.BlackPawns, board.LastBlackPawns)
		if left == To {
			board.WhitePawns = (board.WhitePawns & ^From) | To
			board.BlackPawns &= ^leftTaken
			if board.CheckForCheck(true) {
				board.WhitePawns = (board.WhitePawns & ^To) | From
				board.BlackPawns |= leftTaken
			}
		} else if right == To {
			board.WhitePawns = (board.WhitePawns & ^From) | To
			board.BlackPawns &= ^rightTaken
			if board.CheckForCheck(true) {
				board.WhitePawns = (board.WhitePawns & ^To) | From
				board.BlackPawns |= rightTaken
			}
		} else if whitePawnMoves(From, board.AllPieces, board.AllBlackPieces)&To > 0 {
			before := board.WhitePawns
			piece := board.DeleteOnSquare(To)
			board.WhitePawns = (board.WhitePawns & ^From) | To
			if board.CheckForCheck(true) {
				board.WhitePawns = (board.WhitePawns & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
			promotions = board.PawnPromotions(true)
			whitePawn = true
			newLastWhitePawns = before
		} else {
			return false, false, 0
		}
	} else if board.WhiteRooks&From > 0 && !board.BlacksTurn {
		if rookMoves(From, board.AllPieces, board.AllWhitePieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.WhiteRooks = (board.WhiteRooks & ^From) | To
			if board.CheckForCheck(true) {
				board.WhiteRooks = (board.WhiteRooks & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
			if board.WhiteRooks&a1 == 0 {
				board.WhiteLongCastle = false
			}
			if board.WhiteRooks&h1 == 0 {
				board.WhiteShortCastle = false
			}
		} else {
			return false, false, 0
		}
	} else if board.WhiteKnights&From > 0 && !board.BlacksTurn {
		if knightMoves(From, board.AllWhitePieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.WhiteKnights = (board.WhiteKnights & ^From) | To
			if board.CheckForCheck(true) {
				board.WhiteKnights = (board.WhiteKnights & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
		} else {
			return false, false, 0
		}
	} else if board.WhiteBishops&From > 0 && !board.BlacksTurn {
		if bishopMoves(From, board.AllPieces, board.AllWhitePieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.WhiteBishops = (board.WhiteBishops & ^From) | To
			if board.CheckForCheck(true) {
				board.WhiteBishops = (board.WhiteBishops & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
		} else {
			return false, false, 0
		}
	} else if board.WhiteQueens&From > 0 && !board.BlacksTurn {
		if queenMoves(From, board.AllPieces, board.AllWhitePieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.WhiteQueens = (board.WhiteQueens & ^From) | To
			if board.CheckForCheck(true) {
				board.WhiteQueens = (board.WhiteQueens & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
		} else {
			return false, false, 0
		}
	} else if board.WhiteKing&From > 0 && !board.BlacksTurn {
		attacking := board.BlackAttacking()
		if whiteLongCastle(board.WhiteLongCastle, board.AllPieces, attacking) && To == c1 {
			board.WhiteKing = (board.WhiteKing & ^From) | To
			board.WhiteRooks = (board.WhiteRooks & ^a1) | d1
			board.WhiteLongCastle = false
			board.WhiteShortCastle = false
		} else if whiteShortCastle(board.WhiteShortCastle, board.AllPieces, attacking) && To == g1 {
			board.WhiteKing = (board.WhiteKing & ^From) | To
			board.WhiteRooks = (board.WhiteRooks & ^h1) | f1
			board.WhiteLongCastle = false
			board.WhiteShortCastle = false
		} else if kingMoves(From, board.AllWhitePieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.WhiteKing = (board.WhiteKing & ^From) | To
			if board.CheckForCheck(true) {
				board.WhiteKing = (board.WhiteKing & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
			if board.WhiteKing&e1 == 0 {
				board.WhiteLongCastle = false
				board.WhiteShortCastle = false
			}
		} else {
			return false, false, 0
		}
	} else if board.BlackPawns&From > 0 && board.BlacksTurn {
		left, leftTaken, right, rightTaken := blackEnPassant(From, board.AllPieces, board.WhitePawns, board.LastWhitePawns)
		if left == To {
			board.BlackPawns = (board.BlackPawns & ^From) | To
			board.WhitePawns &= ^leftTaken
			if board.CheckForCheck(false) {
				board.BlackPawns = (board.BlackPawns & ^To) | From
				board.WhitePawns |= leftTaken
			}
		} else if right == To {
			board.BlackPawns = (board.BlackPawns & ^From) | To
			board.WhitePawns &= ^rightTaken
			if board.CheckForCheck(false) {
				board.BlackPawns = (board.BlackPawns & ^To) | From
				board.WhitePawns |= rightTaken
			}
		} else if blackPawnMoves(From, board.AllPieces, board.AllWhitePieces)&To > 0 {
			before := board.BlackPawns
			piece := board.DeleteOnSquare(To)
			board.BlackPawns = (board.BlackPawns & ^From) | To
			if board.CheckForCheck(false) {
				board.BlackPawns = (board.BlackPawns & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
			promotions = board.PawnPromotions(false)
			newLastBlackPawns = before
		} else {
			return false, false, 0
		}
	} else if board.BlackRooks&From > 0 && board.BlacksTurn {
		if rookMoves(From, board.AllPieces, board.AllBlackPieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.BlackRooks = (board.BlackRooks & ^From) | To
			if board.CheckForCheck(false) {
				board.BlackRooks = (board.BlackRooks & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
			if board.BlackRooks&a8 == 0 {
				board.WhiteLongCastle = false
			}
			if board.BlackRooks&h8 == 0 {
				board.WhiteShortCastle = false
			}
		} else {
			return false, false, 0
		}
	} else if board.BlackKnights&From > 0 && board.BlacksTurn {
		if knightMoves(From, board.AllBlackPieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.BlackKnights = (board.BlackKnights & ^From) | To
			if board.CheckForCheck(false) {
				board.BlackKnights = (board.BlackKnights & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
		} else {
			return false, false, 0
		}
	} else if board.BlackBishops&From > 0 && board.BlacksTurn {
		if bishopMoves(From, board.AllPieces, board.AllBlackPieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.BlackBishops = (board.BlackBishops & ^From) | To
			if board.CheckForCheck(false) {
				board.BlackBishops = (board.BlackBishops & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
		} else {
			return false, false, 0
		}
	} else if board.BlackQueens&From > 0 && board.BlacksTurn {
		if queenMoves(From, board.AllPieces, board.AllBlackPieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.BlackQueens = (board.BlackQueens & ^From) | To
			if board.CheckForCheck(false) {
				board.BlackQueens = (board.BlackQueens & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
		} else {
			return false, false, 0
		}
	} else if board.BlackKing&From > 0 && board.BlacksTurn {
		attacking := board.WhiteAttacking()
		if blackLongCastle(board.BlackLongCastle, board.AllPieces, attacking) && To == c8 {
			board.BlackKing = (board.BlackKing & ^From) | To
			board.BlackRooks = (board.BlackRooks & ^a8) | d8
			board.BlackLongCastle = false
			board.BlackShortCastle = false
		} else if blackShortCastle(board.BlackShortCastle, board.AllPieces, attacking) && To == g8 {
			board.BlackKing = (board.BlackKing & ^From) | To
			board.BlackRooks = (board.BlackRooks & ^h8) | f8
			board.BlackLongCastle = false
			board.BlackShortCastle = false
		} else if kingMoves(From, board.AllBlackPieces)&To > 0 {
			piece := board.DeleteOnSquare(To)
			board.BlackKing = (board.BlackKing & ^From) | To
			if board.CheckForCheck(false) {
				board.BlackKing = (board.BlackKing & ^To) | From
				board.PlaceOnSquare(To, piece)
				return false, false, 0
			}
			if board.BlackKing&e8 == 0 {
				board.BlackLongCastle = false
				board.BlackShortCastle = false
			}
		} else {
			return false, false, 0
		}
	} else {
		return false, false, 0
	}
	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces

	board.LastWhitePawns = newLastWhitePawns
	board.LastBlackPawns = newLastBlackPawns

	board.BlacksTurn = !board.BlacksTurn

	if To == a1 {
		board.WhiteLongCastle = false
	} else if To == h1 {
		board.WhiteShortCastle = false
	} else if To == a8 {
		board.BlackLongCastle = false
	} else if To == h8 {
		board.BlackShortCastle = false
	}

	return true, whitePawn, promotions
}

func (board *ChessBoard) DeleteOnSquare(square Bitboard) PieceType {
	if board.WhitePawns&square > 0 {
		board.WhitePawns &= ^square
		return WhitePawn
	} else if board.WhiteRooks&square > 0 {
		board.WhiteRooks &= ^square
		return WhiteRook
	} else if board.WhiteKnights&square > 0 {
		board.WhiteKnights &= ^square
		return WhiteKnight
	} else if board.WhiteBishops&square > 0 {
		board.WhiteBishops &= ^square
		return WhiteBishop
	} else if board.WhiteQueens&square > 0 {
		board.WhiteQueens &= ^square
		return WhiteQueen
	} else if board.WhiteKing&square > 0 {
		board.WhiteKing &= ^square
		return WhiteKing
	} else if board.BlackPawns&square > 0 {
		board.BlackPawns &= ^square
		return BlackPawn
	} else if board.BlackRooks&square > 0 {
		board.BlackRooks &= ^square
		return BlackRook
	} else if board.BlackKnights&square > 0 {
		board.BlackKnights &= ^square
		return BlackKnight
	} else if board.BlackBishops&square > 0 {
		board.BlackBishops &= ^square
		return BlackBishop
	} else if board.BlackQueens&square > 0 {
		board.BlackQueens &= ^square
		return BlackQueen
	} else if board.BlackKing&square > 0 {
		board.BlackKing &= ^square
		return BlackKing
	}
	return 0
}

func (board *ChessBoard) PlaceOnSquare(square Bitboard, piece PieceType) {
	switch piece {
	case WhitePawn:
		board.WhitePawns |= square
	case WhiteRook:
		board.WhiteRooks |= square
	case WhiteKnight:
		board.WhiteKnights |= square
	case WhiteBishop:
		board.WhiteBishops |= square
	case WhiteQueen:
		board.WhiteQueens |= square
	case WhiteKing:
		board.WhiteKing |= square
	case BlackPawn:
		board.BlackPawns |= square
	case BlackRook:
		board.BlackRooks |= square
	case BlackKnight:
		board.BlackKnights |= square
	case BlackBishop:
		board.BlackBishops |= square
	case BlackQueen:
		board.BlackQueens |= square
	case BlackKing:
		board.BlackKing |= square
	}
}

func (board *ChessBoard) MovePieceSimple(From Bitboard, To Bitboard, pieces Bitboard, whiteSide bool) (bool, Bitboard) {
	board.DeleteOnSquare(To)
	newPieces := (pieces & ^From) | To
	if board.CheckForCheck(whiteSide) {
		return false, pieces
	}
	return true, newPieces
}

func (board *ChessBoard) CheckForCheck(whiteSide bool) bool {
	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces
	if whiteSide {
		attacking := board.BlackAttacking()
		if attacking&board.WhiteKing > 0 {
			return true
		}
	} else {
		attacking := board.WhiteAttacking()
		if attacking&board.BlackKing > 0 {
			return true
		}
	}
	return false
}

func (board *ChessBoard) WhiteAttacking() Bitboard {
	attacking := whitePawnAttacks(board.WhitePawns)
	attacking |= rookMoves(board.WhiteRooks, board.AllPieces, board.AllWhitePieces)
	attacking |= knightMoves(board.WhiteKnights, board.AllWhitePieces)
	attacking |= bishopMoves(board.WhiteBishops, board.AllPieces, board.AllWhitePieces)
	attacking |= queenMoves(board.WhiteQueens, board.AllPieces, board.AllWhitePieces)
	attacking |= kingMoves(board.WhiteKing, board.AllWhitePieces)
	return attacking
}

func (board *ChessBoard) BlackAttacking() Bitboard {
	attacking := blackPawnAttacks(board.BlackPawns)
	attacking |= rookMoves(board.BlackRooks, board.AllPieces, board.AllBlackPieces)
	attacking |= knightMoves(board.BlackKnights, board.AllBlackPieces)
	attacking |= bishopMoves(board.BlackBishops, board.AllPieces, board.AllBlackPieces)
	attacking |= queenMoves(board.BlackQueens, board.AllPieces, board.AllBlackPieces)
	attacking |= kingMoves(board.BlackKing, board.AllBlackPieces)
	return attacking
}

func (board *ChessBoard) PawnPromotions(whiteSide bool) Bitboard {
	if whiteSide {
		return board.WhitePawns & maskRank[rank8]
	}
	return board.BlackPawns & maskRank[rank1]
}

func (board *ChessBoard) PromotePawn(square Bitboard, newType PieceType) {
	if newType == WhiteQueen {
		board.WhitePawns &= ^square
		board.WhiteQueens |= square
	} else if newType == WhiteRook {
		board.WhitePawns &= ^square
		board.WhiteRooks |= square
	} else if newType == WhiteBishop {
		board.WhitePawns &= ^square
		board.WhiteBishops |= square
	} else if newType == WhiteKnight {
		board.WhitePawns &= ^square
		board.WhiteKnights |= square
	} else if newType == BlackQueen {
		board.BlackPawns &= ^square
		board.BlackQueens |= square
	} else if newType == BlackRook {
		board.BlackPawns &= ^square
		board.BlackRooks |= square
	} else if newType == BlackBishop {
		board.BlackPawns &= ^square
		board.BlackBishops |= square
	} else if newType == BlackKnight {
		board.BlackPawns &= ^square
		board.BlackKnights |= square
	}
}

func (board *ChessBoard) Init() {
	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces

	board.AllBitboards[0] = &board.WhitePawns
	board.AllBitboards[1] = &board.WhiteRooks
	board.AllBitboards[2] = &board.WhiteKnights
	board.AllBitboards[3] = &board.WhiteBishops
	board.AllBitboards[4] = &board.WhiteQueens
	board.AllBitboards[5] = &board.WhiteKing
	board.AllBitboards[6] = &board.BlackPawns
	board.AllBitboards[7] = &board.BlackRooks
	board.AllBitboards[8] = &board.BlackKnights
	board.AllBitboards[9] = &board.BlackBishops
	board.AllBitboards[10] = &board.BlackQueens
	board.AllBitboards[11] = &board.BlackKing
}

func StandardChessBoard() ChessBoard {
	var board ChessBoard

	board.WhitePawns = a2 | b2 | c2 | d2 | e2 | f2 | g2 | h2
	board.WhiteRooks = a1 | h1
	board.WhiteKnights = b1 | g1
	board.WhiteBishops = c1 | f1
	board.WhiteQueens = d1
	board.WhiteKing = e1

	board.BlackPawns = a7 | b7 | c7 | d7 | e7 | f7 | g7 | h7
	board.BlackRooks = a8 | h8
	board.BlackKnights = b8 | g8
	board.BlackBishops = c8 | f8
	board.BlackQueens = d8
	board.BlackKing = e8

	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces

	board.WhiteLongCastle = true
	board.WhiteShortCastle = true
	board.BlackLongCastle = true
	board.BlackShortCastle = true

	return board
}

func FenString(fen string) ChessBoard {
	parts := strings.Split(fen, " ")

	x := 0
	y := 0
	var board ChessBoard
	for _, char := range parts[0] {
		if char == '/' {
			x = 0
			y++
		} else if unicode.IsDigit(char) {
			x += int(char) - '0'
		} else {
			switch char {
			case 'P':
				board.WhitePawns |= 1 << (63 - (y*8 + x))
			case 'R':
				board.WhiteRooks |= 1 << (63 - (y*8 + x))
			case 'N':
				board.WhiteKnights |= 1 << (63 - (y*8 + x))
			case 'B':
				board.WhiteBishops |= 1 << (63 - (y*8 + x))
			case 'Q':
				board.WhiteQueens |= 1 << (63 - (y*8 + x))
			case 'K':
				board.WhiteKing |= 1 << (63 - (y*8 + x))
			case 'p':
				board.BlackPawns |= 1 << (63 - (y*8 + x))
			case 'r':
				board.BlackRooks |= 1 << (63 - (y*8 + x))
			case 'n':
				board.BlackKnights |= 1 << (63 - (y*8 + x))
			case 'b':
				board.BlackBishops |= 1 << (63 - (y*8 + x))
			case 'q':
				board.BlackQueens |= 1 << (63 - (y*8 + x))
			case 'k':
				board.BlackKing |= 1 << (63 - (y*8 + x))
			}
			x++
		}
	}
	if len(parts) >= 4 {
		if parts[1] == "w" {
			board.BlacksTurn = false
		} else if parts[1] == "b" {
			board.BlacksTurn = true
		}

		for _, char := range parts[2] {
			if char == 'Q' {
				board.WhiteLongCastle = true
			} else if char == 'K' {
				board.WhiteShortCastle = true
			} else if char == 'q' {
				board.BlackLongCastle = true
			} else if char == 'k' {
				board.BlackShortCastle = true
			}
		}

		/*if len(parts[3]) > 1 {
			var rank Bitboard
			if parts[3][0] == '3' {
				rank = maskRank[rank2]
				switch parts[3][1] {
				case 'a':

				}
			} else if parts[3][0] == '6' {
				rank = maskRank[rank7]
			}


		}*/
	}
	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces

	board.AllBitboards[0] = &board.WhitePawns
	board.AllBitboards[1] = &board.WhiteRooks
	board.AllBitboards[2] = &board.WhiteKnights
	board.AllBitboards[3] = &board.WhiteBishops
	board.AllBitboards[4] = &board.WhiteQueens
	board.AllBitboards[5] = &board.WhiteKing
	board.AllBitboards[6] = &board.BlackPawns
	board.AllBitboards[7] = &board.BlackRooks
	board.AllBitboards[8] = &board.BlackKnights
	board.AllBitboards[9] = &board.BlackBishops
	board.AllBitboards[10] = &board.BlackQueens
	board.AllBitboards[11] = &board.BlackKing

	return board
}

func CoordsToBitboard(x, y int) Bitboard {
	return maskFile[x] & maskRank[7-y]
}

func pawnMoves(From, To Bitboard, piece PieceType, enPassant bool) []Move {
	if piece == WhitePawn && To&maskRank[rank8] > 0 {
		return []Move{
			{From: From, To: To, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteQueen},
			{From: From, To: To, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteRook},
			{From: From, To: To, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteBishop},
			{From: From, To: To, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteKnight},
		}
	}
	if piece == BlackPawn && To&maskRank[rank1] > 0 {
		return []Move{
			{From: From, To: To, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackQueen},
			{From: From, To: To, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackRook},
			{From: From, To: To, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackBishop},
			{From: From, To: To, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackKnight},
		}
	}
	return []Move{{From: From, To: To, Piece: piece, EnPassant: enPassant}}
}
