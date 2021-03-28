package bitboard

func kingMoves(kingLoc Bitboard, ownSide Bitboard) Bitboard {
	kingClearFileA := kingLoc & clearFile[fileA]
	kingClearFileH := kingLoc & clearFile[fileH]

	spot1 := kingClearFileA << 9
	spot2 := kingLoc << 8
	spot3 := kingClearFileH << 7
	spot4 := kingClearFileA << 1
	spot5 := kingClearFileH >> 1
	spot6 := kingClearFileA >> 7
	spot7 := kingLoc >> 8
	spot8 := kingClearFileH >> 9

	kingValid := spot1 | spot2 | spot3 | spot4 | spot5 | spot6 | spot7 | spot8
	kingValid = kingValid & ^ownSide
	return kingValid
}

func knightMoves(knightLoc Bitboard, ownSide Bitboard) Bitboard {
	knightClearFileA := knightLoc & clearFile[fileA]
	knightClearFileAB := knightLoc & clearFile[fileA] & clearFile[fileB]
	knightClearFileH := knightLoc & clearFile[fileH]
	knightClearFileGH := knightLoc & clearFile[fileH] & clearFile[fileG]

	spot1 := knightClearFileA << 17
	spot2 := knightClearFileH << 15
	spot3 := knightClearFileAB << 10
	spot4 := knightClearFileGH << 6
	spot5 := knightClearFileAB >> 6
	spot6 := knightClearFileGH >> 10
	spot7 := knightClearFileA >> 15
	spot8 := knightClearFileH >> 17

	knightValid := spot1 | spot2 | spot3 | spot4 | spot5 | spot6 | spot7 | spot8
	knightValid = knightValid & ^ownSide
	return knightValid
}

func whitePawnMoves(pawnLoc Bitboard, allPieces Bitboard, allBlackPieces Bitboard) Bitboard {
	oneStep := (pawnLoc << 8) & ^allPieces
	twoSteps := ((oneStep & maskRank[rank3]) << 8) & ^allPieces
	leftAttack := (pawnLoc & clearFile[fileA]) << 9
	rightAttack := (pawnLoc & clearFile[fileH]) << 7
	pawnValid := oneStep | twoSteps | ((leftAttack | rightAttack) & allBlackPieces)
	return pawnValid
}

func blackPawnMoves(pawnLoc Bitboard, allPieces Bitboard, allWhitePieces Bitboard) Bitboard {
	oneStep := (pawnLoc >> 8) & ^allPieces
	twoSteps := ((oneStep & maskRank[rank6]) >> 8) & ^allPieces
	leftAttack := (pawnLoc & clearFile[fileH]) >> 9
	rightAttack := (pawnLoc & clearFile[fileA]) >> 7
	pawnValid := oneStep | twoSteps | ((leftAttack | rightAttack) & allWhitePieces)
	return pawnValid
}

func whiteEnPassant(pawnLoc Bitboard, allPieces, blackPawns, lastBlackPawns Bitboard) (Bitboard, Bitboard, Bitboard, Bitboard) {
	leftSide := (pawnLoc & clearFile[fileA]) << 9
	leftSide &= ^(allPieces | lastBlackPawns)
	leftSide &= ((lastBlackPawns & ^allPieces) >> 8) & ((blackPawns & ^lastBlackPawns) << 8)

	rightSide := (pawnLoc & clearFile[fileH]) << 7
	rightSide &= ^(allPieces | lastBlackPawns)
	rightSide &= ((lastBlackPawns & ^allPieces) >> 8) & ((blackPawns & ^lastBlackPawns) << 8)

	return leftSide, leftSide >> 8, rightSide, rightSide >> 8
}

func blackEnPassant(pawnLoc Bitboard, allPieces, whitePawns, lastWhitePawns Bitboard) (Bitboard, Bitboard, Bitboard, Bitboard) {
	leftSide := (pawnLoc & clearFile[fileA]) >> 7
	leftSide &= ^(allPieces | lastWhitePawns)
	leftSide &= ((lastWhitePawns & ^allPieces) << 8) & ((whitePawns & ^lastWhitePawns) >> 8)

	rightSide := (pawnLoc & clearFile[fileH]) >> 9
	rightSide &= ^(allPieces | lastWhitePawns)
	rightSide &= ((lastWhitePawns & ^allPieces) << 8) & ((whitePawns & ^lastWhitePawns) >> 8)

	return leftSide, leftSide << 8, rightSide, rightSide << 8
}

func whiteEnPassant2(pawnLoc Bitboard, enPassant uint8) (Bitboard, Bitboard) {
	if enPassant != 8 {
		leftSide := (pawnLoc & clearFile[fileA]) & (1 << ((enPassant - 1) + 32))
		rightSide := (pawnLoc & clearFile[fileH]) & (1 << ((enPassant + 1) + 32))
		return leftSide << 9, rightSide << 7
	}
	return 0, 0
}

func blackEnPassant2(pawnLoc Bitboard, enPassant uint8) (Bitboard, Bitboard) {
	if enPassant != 8 {
		leftSide := (pawnLoc & clearFile[fileA]) & (1 << ((enPassant - 1) + 24))
		rightSide := (pawnLoc & clearFile[fileH]) & (1 << ((enPassant + 1) + 24))
		return leftSide >> 7, rightSide >> 9
	}
	return 0, 0
}

func whitePawnAttacks(pawnLoc Bitboard) Bitboard {
	leftAttack := (pawnLoc & clearFile[fileA]) << 9
	rightAttack := (pawnLoc & clearFile[fileH]) << 7
	return leftAttack | rightAttack
}
func blackPawnAttacks(pawnLoc Bitboard) Bitboard {
	leftAttack := (pawnLoc & clearFile[fileH]) >> 9
	rightAttack := (pawnLoc & clearFile[fileA]) >> 7
	return leftAttack | rightAttack
}

func upAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	pieceLoc = (pieceLoc << 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 8) & emptySpots
	flood |= pieceLoc
	return flood << 8
}

func rightAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	emptySpots = emptySpots & clearFile[fileA]
	pieceLoc = (pieceLoc >> 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 1) & emptySpots
	flood |= pieceLoc
	return (flood >> 1) & clearFile[fileA]
}

func downAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	pieceLoc = (pieceLoc >> 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 8) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 8) & emptySpots
	flood |= pieceLoc
	return flood >> 8
}

func leftAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	emptySpots = emptySpots & clearFile[fileH]
	pieceLoc = (pieceLoc << 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 1) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 1) & emptySpots
	flood |= pieceLoc
	return (flood << 1) & clearFile[fileH]
}

func rightUpAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	emptySpots = emptySpots & clearFile[fileA]
	pieceLoc = (pieceLoc << 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 7) & emptySpots
	flood |= pieceLoc
	return (flood << 7) & clearFile[fileA]
}

func rightDownAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	emptySpots = emptySpots & clearFile[fileA]
	pieceLoc = (pieceLoc >> 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 9) & emptySpots
	flood |= pieceLoc
	return (flood >> 9) & clearFile[fileA]
}

func leftUpAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	emptySpots = emptySpots & clearFile[fileH]
	pieceLoc = (pieceLoc << 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 9) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc << 9) & emptySpots
	flood |= pieceLoc
	return (flood << 9) & clearFile[fileH]
}

func leftDownAttacks(pieceLoc Bitboard, emptySpots Bitboard) Bitboard {
	flood := pieceLoc
	emptySpots = emptySpots & clearFile[fileH]
	pieceLoc = (pieceLoc >> 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 7) & emptySpots
	flood |= pieceLoc
	pieceLoc = (pieceLoc >> 7) & emptySpots
	flood |= pieceLoc
	return (flood >> 7) & clearFile[fileH]
}

func bishopMoves(bishopLoc Bitboard, allPieces Bitboard, ownSide Bitboard) Bitboard {
	rightUp := rightUpAttacks(bishopLoc, ^allPieces)
	rightDown := rightDownAttacks(bishopLoc, ^allPieces)
	leftUp := leftUpAttacks(bishopLoc, ^allPieces)
	leftDown := leftDownAttacks(bishopLoc, ^allPieces)
	return (rightUp | rightDown | leftUp | leftDown) & ^ownSide
}

func rookMoves(rookLoc Bitboard, allPieces Bitboard, ownSide Bitboard) Bitboard {
	up := upAttacks(rookLoc, ^allPieces)
	right := rightAttacks(rookLoc, ^allPieces)
	down := downAttacks(rookLoc, ^allPieces)
	left := leftAttacks(rookLoc, ^allPieces)
	return (up | right | down | left) & ^ownSide
}

func queenMoves(queenLoc Bitboard, allPieces Bitboard, ownSide Bitboard) Bitboard {
	up := upAttacks(queenLoc, ^allPieces)
	right := rightAttacks(queenLoc, ^allPieces)
	down := downAttacks(queenLoc, ^allPieces)
	left := leftAttacks(queenLoc, ^allPieces)
	rightUp := rightUpAttacks(queenLoc, ^allPieces)
	rightDown := rightDownAttacks(queenLoc, ^allPieces)
	leftUp := leftUpAttacks(queenLoc, ^allPieces)
	leftDown := leftDownAttacks(queenLoc, ^allPieces)
	return (up | right | down | left | rightUp | rightDown | leftUp | leftDown) & ^ownSide
}

func whiteLongCastle(castleValid bool, allPieces Bitboard, blackAttacking Bitboard) bool {
	if !castleValid {
		return false
	}
	if blackAttacking&(c1|d1|e1) == 0 && allPieces&(b1|c1|d1) == 0 {
		return true
	}
	return false
}

func whiteShortCastle(castleValid bool, allPieces Bitboard, blackAttacking Bitboard) bool {
	if !castleValid {
		return false
	}
	if blackAttacking&(e1|f1|g1) == 0 && allPieces&(f1|g1) == 0 {
		return true
	}
	return false
}

func blackLongCastle(castleValid bool, allPieces Bitboard, whiteAttacking Bitboard) bool {
	if !castleValid {
		return false
	}
	if whiteAttacking&(c8|d8|e8) == 0 && allPieces&(b8|c8|d8) == 0 {
		return true
	}
	return false
}

func blackShortCastle(castleValid bool, allPieces Bitboard, whiteAttacking Bitboard) bool {
	if !castleValid {
		return false
	}
	if whiteAttacking&(e8|f8|g8) == 0 && allPieces&(f8|g8) == 0 {
		return true
	}
	return false
}
