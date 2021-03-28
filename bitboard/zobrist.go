package bitboard

import "math/rand"

var positionHashes [12][64]uint64
var blacksTurnHash uint64
var whiteLongCastleHash uint64
var whiteShortCastleHash uint64
var blackLongCastleHash uint64
var blackShortCastleHash uint64
var enPassantHashes [8]uint64

func init() {
	for i := 0; i < 12; i++ {
		for j := 0; j < 64; j++ {
			positionHashes[i][j] = rand.Uint64()
		}
	}
	blacksTurnHash = rand.Uint64()
	whiteLongCastleHash = rand.Uint64()
	whiteShortCastleHash = rand.Uint64()
	blackLongCastleHash = rand.Uint64()
	blackShortCastleHash = rand.Uint64()
	for i := 0; i < 8; i++ {
		enPassantHashes[i] = rand.Uint64()
	}
}

func boardToHash(board *ChessBoard) uint64 {
	hash := uint64(0)
	for i, piece := range board.AllBitboards {
		for _, pos := range BitboardToSlice(*piece) {
			hash ^= positionHashes[i][pos]
		}
	}
	if board.BlacksTurn {
		hash ^= blacksTurnHash
	}
	if board.WhiteLongCastle {
		hash ^= whiteLongCastleHash
	}
	if board.WhiteShortCastle {
		hash ^= whiteShortCastleHash
	}
	if board.BlackLongCastle {
		hash ^= blackLongCastleHash
	}
	if board.BlackShortCastle {
		hash ^= blackShortCastleHash
	}

	return hash
}
