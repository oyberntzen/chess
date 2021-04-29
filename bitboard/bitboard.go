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

	WhiteEnPassant uint8
	BlackEnPassant uint8

	BlacksTurn bool

	Zobrist uint64

	LastHashes []uint64
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

	FromIndex uint8
	ToIndex   uint8

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
)

var (
	a1, a2, a3, a4, a5, a6, a7, a8 Bitboard = maskRank[0] & maskFile[0], maskRank[1] & maskFile[0], maskRank[2] & maskFile[0], maskRank[3] & maskFile[0], maskRank[4] & maskFile[0], maskRank[5] & maskFile[0], maskRank[6] & maskFile[0], maskRank[7] & maskFile[0]
	b1, b2, b3, b4, b5, b6, b7, b8 Bitboard = maskRank[0] & maskFile[1], maskRank[1] & maskFile[1], maskRank[2] & maskFile[1], maskRank[3] & maskFile[1], maskRank[4] & maskFile[1], maskRank[5] & maskFile[1], maskRank[6] & maskFile[1], maskRank[7] & maskFile[1]
	c1, c2, c3, c4, c5, c6, c7, c8 Bitboard = maskRank[0] & maskFile[2], maskRank[1] & maskFile[2], maskRank[2] & maskFile[2], maskRank[3] & maskFile[2], maskRank[4] & maskFile[2], maskRank[5] & maskFile[2], maskRank[6] & maskFile[2], maskRank[7] & maskFile[2]
	d1, d2, d3, d4, d5, d6, d7, d8 Bitboard = maskRank[0] & maskFile[3], maskRank[1] & maskFile[3], maskRank[2] & maskFile[3], maskRank[3] & maskFile[3], maskRank[4] & maskFile[3], maskRank[5] & maskFile[3], maskRank[6] & maskFile[3], maskRank[7] & maskFile[3]
	e1, e2, e3, e4, e5, e6, e7, e8 Bitboard = maskRank[0] & maskFile[4], maskRank[1] & maskFile[4], maskRank[2] & maskFile[4], maskRank[3] & maskFile[4], maskRank[4] & maskFile[4], maskRank[5] & maskFile[4], maskRank[6] & maskFile[4], maskRank[7] & maskFile[4]
	f1, f2, f3, f4, f5, f6, f7, f8 Bitboard = maskRank[0] & maskFile[5], maskRank[1] & maskFile[5], maskRank[2] & maskFile[5], maskRank[3] & maskFile[5], maskRank[4] & maskFile[5], maskRank[5] & maskFile[5], maskRank[6] & maskFile[5], maskRank[7] & maskFile[5]
	g1, g2, g3, g4, g5, g6, g7, g8 Bitboard = maskRank[0] & maskFile[6], maskRank[1] & maskFile[6], maskRank[2] & maskFile[6], maskRank[3] & maskFile[6], maskRank[4] & maskFile[6], maskRank[5] & maskFile[6], maskRank[6] & maskFile[6], maskRank[7] & maskFile[6]
	h1, h2, h3, h4, h5, h6, h7, h8 Bitboard = maskRank[0] & maskFile[7], maskRank[1] & maskFile[7], maskRank[2] & maskFile[7], maskRank[3] & maskFile[7], maskRank[4] & maskFile[7], maskRank[5] & maskFile[7], maskRank[6] & maskFile[7], maskRank[7] & maskFile[7]
)

var posToString map[Bitboard]string = map[Bitboard]string{
	a1: "a1", a2: "a2", a3: "a3", a4: "a4", a5: "a5", a6: "a6", a7: "a7", a8: "a8",
	b1: "b1", b2: "b2", b3: "b3", b4: "b4", b5: "b5", b6: "b6", b7: "b7", b8: "b8",
	c1: "c1", c2: "c2", c3: "c3", c4: "c4", c5: "c5", c6: "c6", c7: "c7", c8: "c8",
	d1: "d1", d2: "d2", d3: "d3", d4: "d4", d5: "d5", d6: "d6", d7: "d7", d8: "d8",
	e1: "e1", e2: "e2", e3: "e3", e4: "e4", e5: "e5", e6: "e6", e7: "e7", e8: "e8",
	f1: "f1", f2: "f2", f3: "f3", f4: "f4", f5: "f5", f6: "f6", f7: "f7", f8: "f8",
	g1: "g1", g2: "g2", g3: "g3", g4: "g4", g5: "g5", g6: "g6", g7: "g7", g8: "g8",
	h1: "h1", h2: "h2", h3: "h3", h4: "h4", h5: "h5", h6: "h6", h7: "h7", h8: "h8",
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
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<To, From, To, WhitePawn, false)...)
			}
			left, right := whiteEnPassant(1<<From, board.BlackEnPassant)
			if left > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From+9), From, From+9, WhitePawn, true)...)
			}
			if right > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From+7), From, From+7, WhitePawn, true)...)
			}
		}

		for _, From := range BitboardToSlice(board.WhiteRooks) {
			moves := rookMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: WhiteRook})
			}
		}

		for _, From := range BitboardToSlice(board.WhiteKnights) {
			moves := knightMoves(1<<From, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: WhiteKnight})
			}
		}

		for _, From := range BitboardToSlice(board.WhiteBishops) {
			moves := bishopMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: WhiteBishop})
			}
		}

		for _, From := range BitboardToSlice(board.WhiteQueens) {
			moves := queenMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: WhiteQueen})
			}
		}

		moves := kingMoves(board.WhiteKing, board.AllWhitePieces) & andWith
		fromIndex := BitboardToSlice(board.WhiteKing)[0]
		for _, To := range BitboardToSlice(moves) {
			psudomoves = append(psudomoves, Move{From: board.WhiteKing, To: 1 << To, FromIndex: fromIndex, ToIndex: To, Piece: WhiteKing})
		}
		attacking := board.BlackAttacking()
		if whiteLongCastle(board.WhiteLongCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.WhiteKing, To: c1, FromIndex: fromIndex, ToIndex: 5, Piece: WhiteKing, LongCastle: true})
		}
		if whiteShortCastle(board.WhiteShortCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.WhiteKing, To: g1, FromIndex: fromIndex, ToIndex: 1, Piece: WhiteKing, ShortCastle: true})
		}
	} else {
		for _, From := range BitboardToSlice(board.BlackPawns) {
			moves := blackPawnMoves(1<<From, board.AllPieces, board.AllWhitePieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<To, From, To, BlackPawn, false)...)
			}
			left, right := blackEnPassant(1<<From, board.WhiteEnPassant)
			if left > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From-7), From, From-7, BlackPawn, true)...)
			}
			if right > 0 {
				psudomoves = append(psudomoves, pawnMoves(1<<From, 1<<(From-9), From, From-9, BlackPawn, true)...)
			}
		}

		for _, From := range BitboardToSlice(board.BlackRooks) {
			moves := rookMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: BlackRook})
			}
		}

		for _, From := range BitboardToSlice(board.BlackKnights) {
			moves := knightMoves(1<<From, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: BlackKnight})
			}
		}

		for _, From := range BitboardToSlice(board.BlackBishops) {
			moves := bishopMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: BlackBishop})
			}
		}

		for _, From := range BitboardToSlice(board.BlackQueens) {
			moves := queenMoves(1<<From, board.AllPieces, board.AllBlackPieces) & andWith
			for _, To := range BitboardToSlice(moves) {
				psudomoves = append(psudomoves, Move{From: 1 << From, To: 1 << To, FromIndex: From, ToIndex: To, Piece: BlackQueen})
			}
		}

		moves := kingMoves(board.BlackKing, board.AllBlackPieces) & andWith
		fromIndex := BitboardToSlice(board.BlackKing)[0]
		for _, To := range BitboardToSlice(moves) {
			psudomoves = append(psudomoves, Move{From: board.BlackKing, To: 1 << To, FromIndex: fromIndex, ToIndex: To, Piece: BlackKing})
		}
		attacking := board.WhiteAttacking()
		if blackLongCastle(board.BlackLongCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.BlackKing, To: c8, FromIndex: fromIndex, ToIndex: 61, Piece: BlackKing, LongCastle: true})
		}
		if blackShortCastle(board.BlackShortCastle, board.AllPieces, attacking) && andWith == ^Bitboard(0) {
			psudomoves = append(psudomoves, Move{From: board.BlackKing, To: g8, FromIndex: fromIndex, ToIndex: 57, Piece: BlackKing, ShortCastle: true})
		}
	}
	return psudomoves
}

func (board *ChessBoard) DoMove(m Move) {
	if board.WhiteEnPassant != 8 {
		board.Zobrist ^= enPassantHashes[board.WhiteEnPassant]
		board.WhiteEnPassant = 8
	}
	if board.BlackEnPassant != 8 {
		board.Zobrist ^= enPassantHashes[board.BlackEnPassant]
		board.BlackEnPassant = 8
	}

	if m.Piece == WhitePawn {
		if m.EnPassant {
			board.DeleteOnSquare(m.To>>8, m.ToIndex-8)
		} else {
			board.DeleteOnSquare(m.To, m.ToIndex)
		}
		board.WhitePawns = (board.WhitePawns & ^m.From) | m.To
		board.Zobrist ^= positionHashes[WhitePawn-1][m.FromIndex] ^ positionHashes[WhitePawn-1][m.ToIndex]
		if m.PawnPromotionPiece != 0 {
			board.PromotePawn(m.To, m.PawnPromotionPiece)
		}
		if m.ToIndex-m.FromIndex == 16 {
			board.WhiteEnPassant = m.FromIndex % 8
			board.Zobrist ^= enPassantHashes[m.FromIndex%8]
		}
		board.LastHashes = []uint64{}
	} else if m.Piece == WhiteRook {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.WhiteRooks = (board.WhiteRooks & ^m.From) | m.To
		board.Zobrist ^= positionHashes[WhiteRook-1][m.FromIndex] ^ positionHashes[WhiteRook-1][m.ToIndex]
	} else if m.Piece == WhiteKnight {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.Zobrist ^= positionHashes[WhiteKnight-1][m.FromIndex] ^ positionHashes[WhiteKnight-1][m.ToIndex]
		board.WhiteKnights = (board.WhiteKnights & ^m.From) | m.To
	} else if m.Piece == WhiteBishop {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.Zobrist ^= positionHashes[WhiteBishop-1][m.FromIndex] ^ positionHashes[WhiteBishop-1][m.ToIndex]
		board.WhiteBishops = (board.WhiteBishops & ^m.From) | m.To
	} else if m.Piece == WhiteQueen {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.Zobrist ^= positionHashes[WhiteQueen-1][m.FromIndex] ^ positionHashes[WhiteQueen-1][m.ToIndex]
		board.WhiteQueens = (board.WhiteQueens & ^m.From) | m.To
	} else if m.Piece == WhiteKing {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.WhiteKing = (board.WhiteKing & ^m.From) | m.To
		board.Zobrist ^= positionHashes[WhiteKing-1][m.FromIndex] ^ positionHashes[WhiteKing-1][m.ToIndex]
		if m.LongCastle {
			board.WhiteRooks = (board.WhiteRooks & ^a1) | d1
			board.Zobrist ^= positionHashes[WhiteRook-1][7] ^ positionHashes[WhiteRook-1][4]
		} else if m.ShortCastle {
			board.WhiteRooks = (board.WhiteRooks & ^h1) | f1
			board.Zobrist ^= positionHashes[WhiteRook-1][0] ^ positionHashes[WhiteRook-1][2]
		}
		if board.WhiteLongCastle {
			board.WhiteLongCastle = false
			board.Zobrist ^= whiteLongCastleHash
		}
		if board.WhiteShortCastle {
			board.WhiteShortCastle = false
			board.Zobrist ^= whiteShortCastleHash
		}
	}

	if m.Piece == BlackPawn {
		if m.EnPassant {
			board.DeleteOnSquare(m.To<<8, m.ToIndex+8)
		} else {
			board.DeleteOnSquare(m.To, m.ToIndex)
		}
		board.BlackPawns = (board.BlackPawns & ^m.From) | m.To
		board.Zobrist ^= positionHashes[BlackPawn-1][m.FromIndex] ^ positionHashes[BlackPawn-1][m.ToIndex]
		if m.PawnPromotionPiece != 0 {
			board.PromotePawn(m.To, m.PawnPromotionPiece)
		}
		if m.FromIndex-m.ToIndex == 16 {
			board.BlackEnPassant = m.FromIndex % 8
			board.Zobrist ^= enPassantHashes[m.FromIndex%8]
		}
		board.LastHashes = []uint64{}
	} else if m.Piece == BlackRook {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.BlackRooks = (board.BlackRooks & ^m.From) | m.To
		board.Zobrist ^= positionHashes[BlackRook-1][m.FromIndex] ^ positionHashes[BlackRook-1][m.ToIndex]
	} else if m.Piece == BlackKnight {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.BlackKnights = (board.BlackKnights & ^m.From) | m.To
		board.Zobrist ^= positionHashes[BlackKnight-1][m.FromIndex] ^ positionHashes[BlackKnight-1][m.ToIndex]
	} else if m.Piece == BlackBishop {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.BlackBishops = (board.BlackBishops & ^m.From) | m.To
		board.Zobrist ^= positionHashes[BlackBishop-1][m.FromIndex] ^ positionHashes[BlackBishop-1][m.ToIndex]
	} else if m.Piece == BlackQueen {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.BlackQueens = (board.BlackQueens & ^m.From) | m.To
		board.Zobrist ^= positionHashes[BlackQueen-1][m.FromIndex] ^ positionHashes[BlackQueen-1][m.ToIndex]
	} else if m.Piece == BlackKing {
		board.DeleteOnSquare(m.To, m.ToIndex)
		board.BlackKing = (board.BlackKing & ^m.From) | m.To
		board.Zobrist ^= positionHashes[BlackKing-1][m.FromIndex] ^ positionHashes[BlackKing-1][m.ToIndex]
		if m.LongCastle {
			board.BlackRooks = (board.BlackRooks & ^a8) | d8
			board.Zobrist ^= positionHashes[BlackRook-1][63] ^ positionHashes[BlackRook-1][60]
		} else if m.ShortCastle {
			board.BlackRooks = (board.BlackRooks & ^h8) | f8
			board.Zobrist ^= positionHashes[BlackRook-1][58] ^ positionHashes[BlackRook-1][56]
		}
		if board.BlackLongCastle {
			board.BlackLongCastle = false
			board.Zobrist ^= blackLongCastleHash
		}
		if board.BlackShortCastle {
			board.BlackShortCastle = false
			board.Zobrist ^= blackShortCastleHash
		}
	}

	if board.WhiteRooks&a1 == 0 && board.WhiteLongCastle {
		board.WhiteLongCastle = false
		board.Zobrist ^= whiteLongCastleHash
	}
	if board.WhiteRooks&h1 == 0 && board.WhiteShortCastle {
		board.WhiteShortCastle = false
		board.Zobrist ^= whiteShortCastleHash
	}
	if board.BlackRooks&a8 == 0 && board.BlackLongCastle {
		board.BlackLongCastle = false
		board.Zobrist ^= blackLongCastleHash
	}
	if board.BlackRooks&h8 == 0 && board.BlackShortCastle {
		board.BlackShortCastle = false
		board.Zobrist ^= blackShortCastleHash
	}

	board.AllWhitePieces = board.WhitePawns | board.WhiteRooks | board.WhiteKnights | board.WhiteBishops | board.WhiteQueens | board.WhiteKing
	board.AllBlackPieces = board.BlackPawns | board.BlackRooks | board.BlackKnights | board.BlackBishops | board.BlackQueens | board.BlackKing
	board.AllPieces = board.AllWhitePieces | board.AllBlackPieces

	board.BlacksTurn = !board.BlacksTurn
	board.Zobrist ^= blacksTurnHash

	board.LastHashes = append(board.LastHashes, board.Zobrist)
}

func (board *ChessBoard) DeleteOnSquare(square Bitboard, index uint8) PieceType {
	deleted := true
	if board.WhitePawns&square > 0 {
		board.WhitePawns &= ^square
		board.Zobrist ^= positionHashes[WhitePawn-1][index]
		return WhitePawn
	} else if board.WhiteRooks&square > 0 {
		board.WhiteRooks &= ^square
		board.Zobrist ^= positionHashes[WhiteRook-1][index]
		return WhiteRook
	} else if board.WhiteKnights&square > 0 {
		board.WhiteKnights &= ^square
		board.Zobrist ^= positionHashes[WhiteKnight-1][index]
		return WhiteKnight
	} else if board.WhiteBishops&square > 0 {
		board.WhiteBishops &= ^square
		board.Zobrist ^= positionHashes[WhiteBishop-1][index]
		return WhiteBishop
	} else if board.WhiteQueens&square > 0 {
		board.WhiteQueens &= ^square
		board.Zobrist ^= positionHashes[WhiteQueen-1][index]
		return WhiteQueen
	} else if board.WhiteKing&square > 0 {
		board.WhiteKing &= ^square
		board.Zobrist ^= positionHashes[WhiteKing-1][index]
		return WhiteKing
	} else if board.BlackPawns&square > 0 {
		board.BlackPawns &= ^square
		board.Zobrist ^= positionHashes[BlackPawn-1][index]
		return BlackPawn
	} else if board.BlackRooks&square > 0 {
		board.BlackRooks &= ^square
		board.Zobrist ^= positionHashes[BlackRook-1][index]
		return BlackRook
	} else if board.BlackKnights&square > 0 {
		board.BlackKnights &= ^square
		board.Zobrist ^= positionHashes[BlackKnight-1][index]
		return BlackKnight
	} else if board.BlackBishops&square > 0 {
		board.BlackBishops &= ^square
		board.Zobrist ^= positionHashes[BlackBishop-1][index]
		return BlackBishop
	} else if board.BlackQueens&square > 0 {
		board.BlackQueens &= ^square
		board.Zobrist ^= positionHashes[BlackQueen-1][index]
		return BlackQueen
	} else if board.BlackKing&square > 0 {
		board.BlackKing &= ^square
		board.Zobrist ^= positionHashes[BlackKing-1][index]
		return BlackKing
	} else {
		deleted = false
	}
	if deleted {
		board.LastHashes = []uint64{}
	}
	return 0
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

	board.WhiteEnPassant = 8
	board.BlackEnPassant = 8

	board.Zobrist = boardToHash(board)
	board.LastHashes = []uint64{board.Zobrist}
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

	return board
}

func CoordsToBitboard(x, y int) Bitboard {
	return maskFile[x] & maskRank[7-y]
}

func pawnMoves(From, To Bitboard, fromIndex, toIndex uint8, piece PieceType, enPassant bool) []Move {
	if piece == WhitePawn && To&maskRank[rank8] > 0 {
		return []Move{
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteQueen},
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteRook},
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteBishop},
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: WhitePawn, EnPassant: enPassant, PawnPromotionPiece: WhiteKnight},
		}
	}
	if piece == BlackPawn && To&maskRank[rank1] > 0 {
		return []Move{
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackQueen},
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackRook},
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackBishop},
			{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: BlackPawn, EnPassant: enPassant, PawnPromotionPiece: BlackKnight},
		}
	}
	return []Move{{From: From, To: To, FromIndex: fromIndex, ToIndex: toIndex, Piece: piece, EnPassant: enPassant}}
}
