package bitboard

import (
	"math/bits"
)

var (
	mgValues     [6]int32 = [6]int32{82, 477, 337, 365, 1025, 0}
	egValues     [6]int32 = [6]int32{94, 512, 281, 297, 936, 0}
	gamePhaseInc [6]int32 = [6]int32{0, 2, 1, 1, 4, 0}
)

type pieceSquareTable struct {
	table  [64]int32
	lookup [8][256]int32
}

func init() {
	for _, table := range mgTables {
		table.init()
	}
	for _, table := range egTables {
		table.init()
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 256; j++ {
			indicies := []uint8{}
			for k := 0; k < 8; k++ {
				if (1<<k)&j > 0 {
					indicies = append(indicies, uint8(k+i*8))
				}
			}
			numberToBits[i][j] = indicies
		}
	}
}

func (pst *pieceSquareTable) init() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 256; j++ {
			score := int32(0)
			for k := 0; k < 8; k++ {
				if (1<<k)&j > 0 {
					score += pst.table[(7-i)*8+(7-k)]
				}
			}
			pst.lookup[i][j] = score
		}
	}
}

func BitboardToSlice(board Bitboard) []uint8 {
	return concatCopyPreAllocate([][]uint8{
		numberToBits[0][board&maskRank[rank1]],
		numberToBits[1][(board>>8)&maskRank[rank1]],
		numberToBits[2][(board>>16)&maskRank[rank1]],
		numberToBits[3][(board>>24)&maskRank[rank1]],
		numberToBits[4][(board>>32)&maskRank[rank1]],
		numberToBits[5][(board>>40)&maskRank[rank1]],
		numberToBits[6][(board>>48)&maskRank[rank1]],
		numberToBits[7][(board>>56)&maskRank[rank1]],
	}, bits.OnesCount64(uint64(board)))
}

func concatCopyPreAllocate(slices [][]uint8, length int) []uint8 {
	tmp := make([]uint8, length)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

var mgPawnTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	0, 0, 0, 0, 0, 0, 0, 0,
	98, 134, 61, 95, 68, 126, 34, -11,
	-6, 7, 26, 31, 65, 56, 25, -20,
	-14, 13, 6, 21, 23, 12, 17, -23,
	-27, -2, -5, 12, 17, 6, 10, -25,
	-26, -4, -4, -10, 3, 3, 33, -12,
	-35, -1, -20, -23, -15, 24, 38, -22,
	0, 0, 0, 0, 0, 0, 0, 0,
}}

var egPawnTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	0, 0, 0, 0, 0, 0, 0, 0,
	178, 173, 158, 134, 147, 132, 165, 187,
	94, 100, 85, 67, 56, 53, 82, 84,
	32, 24, 13, 5, -2, 4, 17, 17,
	13, 9, -3, -7, -7, -8, 3, -1,
	4, 7, -6, 1, 0, -5, -1, -8,
	13, 8, 8, 10, 13, 0, 2, -7,
	0, 0, 0, 0, 0, 0, 0, 0,
}}

var mgKnightTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-167, -89, -34, -49, 61, -97, -15, -107,
	-73, -41, 72, 36, 23, 62, 7, -17,
	-47, 60, 37, 65, 84, 129, 73, 44,
	-9, 17, 19, 53, 37, 69, 18, 22,
	-13, 4, 16, 13, 28, 19, 21, -8,
	-23, -9, 12, 10, 19, 17, 25, -16,
	-29, -53, -12, -3, -1, 18, -14, -19,
	-105, -21, -58, -33, -17, -28, -19, -23,
}}

var egKnightTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-58, -38, -13, -28, -31, -27, -63, -99,
	-25, -8, -25, -2, -9, -25, -24, -52,
	-24, -20, 10, 9, -1, -9, -19, -41,
	-17, 3, 22, 22, 22, 11, 8, -18,
	-18, -6, 16, 25, 16, 17, 4, -18,
	-23, -3, -1, 15, 10, -3, -20, -22,
	-42, -20, -10, -5, -2, -20, -23, -44,
	-29, -51, -23, -15, -22, -18, -50, -64,
}}

var mgBishopTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-29, 4, -82, -37, -25, -42, 7, -8,
	-26, 16, -18, -13, 30, 59, 18, -47,
	-16, 37, 43, 40, 35, 50, 37, -2,
	-4, 5, 19, 50, 37, 37, 7, -2,
	-6, 13, 13, 26, 34, 12, 10, 4,
	0, 15, 15, 15, 14, 27, 18, 10,
	4, 15, 16, 0, 7, 21, 33, 1,
	-33, -3, -14, -21, -13, -12, -39, -21,
}}

var egBishopTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-14, -21, -11, -8, -7, -9, -17, -24,
	-8, -4, 7, -12, -3, -13, -4, -14,
	2, -8, 0, -1, -2, 6, 0, 4,
	-3, 9, 12, 9, 14, 10, 3, 2,
	-6, 3, 13, 19, 7, 10, -3, -9,
	-12, -3, 8, 10, 13, 3, -7, -15,
	-14, -18, -7, -1, 4, -9, -15, -27,
	-23, -9, -23, -5, -9, -16, -5, -17,
}}

var mgRookTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	32, 42, 32, 51, 63, 9, 31, 43,
	27, 32, 58, 62, 80, 67, 26, 44,
	-5, 19, 26, 36, 17, 45, 61, 16,
	-24, -11, 7, 26, 24, 35, -8, -20,
	-36, -26, -12, -1, 9, -7, 6, -23,
	-45, -25, -16, -17, 3, 0, -5, -33,
	-44, -16, -20, -9, -1, 11, -6, -71,
	-19, -13, 1, 17, 16, 7, -37, -26,
}}

var egRookTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	13, 10, 18, 15, 12, 12, 8, 5,
	11, 13, 13, 11, -3, 3, 8, 3,
	7, 7, 7, 5, 4, -3, -5, -3,
	4, 3, 13, 1, 2, 1, -1, 2,
	3, 5, 8, 4, -5, -6, -8, -11,
	-4, 0, -5, -1, -7, -12, -8, -16,
	-6, -6, 0, 2, -9, -9, -11, -3,
	-9, 2, 3, -1, -5, -13, 4, -20,
}}

var mgQueenTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-28, 0, 29, 12, 59, 44, 43, 45,
	-24, -39, -5, 1, -16, 57, 28, 54,
	-13, -17, 7, 8, 29, 56, 47, 57,
	-27, -27, -16, -16, -1, 17, -2, 1,
	-9, -26, -9, -10, -2, -4, 3, -3,
	-14, 2, -11, -2, -5, 2, 14, 5,
	-35, -8, 11, 2, 8, 15, -3, 1,
	-1, -18, -9, 10, -15, -25, -31, -50,
}}

var egQueenTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-9, 22, 22, 27, 27, 19, 10, 20,
	-17, 20, 32, 41, 58, 25, 30, 0,
	-20, 6, 9, 49, 47, 35, 19, 9,
	3, 22, 24, 45, 57, 40, 57, 36,
	-18, 28, 19, 47, 31, 34, 39, 23,
	-16, -27, 15, 6, 9, 17, 10, 5,
	-22, -23, -30, -16, -16, -23, -36, -32,
	-33, -28, -22, -43, -5, -32, -20, -41,
}}

var mgKingTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-65, 23, 16, -15, -56, -34, 2, 13,
	29, -1, -20, -7, -8, -4, -38, -29,
	-9, 24, 2, -16, -20, 6, 22, -22,
	-17, -20, -12, -27, -30, -25, -14, -36,
	-49, -1, -27, -39, -46, -44, -33, -51,
	-14, -14, -22, -46, -44, -30, -15, -27,
	1, 7, -8, -64, -43, -16, 9, 8,
	-15, 36, 12, -54, 8, -28, 24, 14,
}}

var egKingTable pieceSquareTable = pieceSquareTable{table: [64]int32{
	-74, -35, -18, -18, -11, 15, 4, -17,
	-12, 17, 14, 17, 17, 38, 23, 11,
	10, 17, 23, 15, 20, 45, 44, 13,
	-8, 22, 24, 27, 26, 33, 26, 3,
	-18, -4, 21, 24, 27, 23, 9, -11,
	-19, -3, 11, 21, 23, 16, 7, -9,
	-27, -11, 4, 13, 14, 4, -5, -17,
	-53, -34, -21, -11, -28, -14, -24, -43,
}}

var mgTables [6]*pieceSquareTable = [6]*pieceSquareTable{
	&mgPawnTable,
	&mgRookTable,
	&mgKnightTable,
	&mgBishopTable,
	&mgQueenTable,
	&mgKingTable,
}

var egTables [6]*pieceSquareTable = [6]*pieceSquareTable{
	&egPawnTable,
	&egRookTable,
	&egKnightTable,
	&egBishopTable,
	&egQueenTable,
	&egKingTable,
}

func evaluate(board *ChessBoard) int32 {
	mgScore := int32(0)
	egScore := int32(0)
	mgPhase := int32(0)

	for i := 0; i < 12; i++ {
		count := int32(bits.OnesCount64(uint64(*board.AllBitboards[i])))
		piece := i % 6
		if i < 6 {
			mgScore += count * mgValues[piece]
			egScore += count * egValues[piece]
			for j := 0; j < 8; j++ {
				index := (*board.AllBitboards[i] >> (j * 8)) & maskRank[rank1]
				mgScore += mgTables[piece].lookup[j][index]
				egScore += egTables[piece].lookup[j][index]
			}
		} else {
			mgScore -= count * mgValues[piece]
			egScore -= count * egValues[piece]
			for j := 0; j < 8; j++ {
				index := (*board.AllBitboards[i] >> ((7 - j) * 8)) & maskRank[rank1]
				mgScore -= mgTables[piece].lookup[j][index]
				egScore -= egTables[piece].lookup[j][index]
			}
		}
		mgPhase += count * gamePhaseInc[piece]
	}
	if board.BlacksTurn {
		mgScore *= -1
		egScore *= -1
	}

	if mgPhase > 24 {
		mgPhase = 24
	}
	egPhase := 24 - mgPhase

	return (mgScore*mgPhase + egScore*egPhase) / 24
}
