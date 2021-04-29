package bitboard

import (
	"unsafe"
)

type Entry struct {
	Zobrist uint64
	Data    uint64
}

type NodeType uint8

type Position struct {
	Zobrist  uint64
	BestMove Move
	Depth    uint8
	Score    int32
	Node     NodeType
	Age      uint8
}

const (
	UpperBoundNode NodeType = 1
	LowerBoundNode NodeType = 2
	ExactNode      NodeType = 3

	Mask8bit  uint64 = 0x00000000000000ff
	Mask16bit uint64 = 0x000000000000ffff
	Mask32bit uint64 = 0x00000000ffffffff
)

const tableSize uint64 = 100_000

var transpositionTable [tableSize]Entry

func StoreEntry(zobrist uint64, bestMoveIndex uint8, depth uint8, score int32, node NodeType, age uint8) {
	bestMoveIndexData := uint64(bestMoveIndex)                  //8-bit
	depthData := uint64(depth)                                  //8-bit
	scoreData := *(*uint64)(unsafe.Pointer(&score)) & Mask32bit //32-bit
	nodeData := uint64(node)                                    //8-bit
	ageData := uint64(age)                                      //8-bit

	data := (bestMoveIndexData) | (depthData << 8) | (scoreData << 16) | (nodeData << 48) | (ageData << 56)

	index := zobrist % tableSize
	transpositionTable[index].Zobrist = zobrist ^ data
	transpositionTable[index].Data = data
}

func GetEntry(zobrist uint64) (uint8, uint8, int32, NodeType, uint8, bool) {
	index := zobrist % tableSize
	matching := false
	if transpositionTable[index].Zobrist^transpositionTable[index].Data == zobrist {
		matching = true
	}
	data := transpositionTable[index].Data
	bestMoveIndexData := data & Mask8bit
	depthData := (data >> 8) & Mask8bit
	scoreData := (data >> 16) & Mask32bit
	nodeData := (data >> 48) & Mask8bit
	ageData := (data >> 56) & Mask8bit
	return uint8(bestMoveIndexData), uint8(depthData), *(*int32)(unsafe.Pointer(&scoreData)), NodeType(nodeData), uint8(ageData), matching
}
