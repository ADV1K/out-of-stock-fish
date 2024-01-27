package main

// AlgebraicNotationToSquareNum converts a square in algebraic notation to a square number.
// Example: e4 -> 28 (e4 is the 28th square on the board)
func AlgebraicNotationToSquareNum(sq string) int {
	if len(sq) != 2 {
		panic("AlgebraicNotationToSquare: Invalid square")
	}
	return int((sq[1]-'1')*8 + sq[0] - 'a')
}

// generateKnightAttacks generates a lookup table for knight attacks.
// Read More: https://www.chessprogramming.org/Knight_Pattern
func generateKnightAttacks() (table [64]Bitboard) {
	for square := 0; square < 64; square++ {
		var piece Bitboard = 1 << square
		table[square] = (piece << 17) & NotAFile   // Up 2, right 1
		table[square] |= (piece << 15) & NotHFile  // Up 2, left 1
		table[square] |= (piece << 10) & NotABFile // Up 1, right 2
		table[square] |= (piece << 6) & NotGHFile  // Up 1, left 2
		table[square] |= (piece >> 17) & NotHFile  // Down 2, left 1
		table[square] |= (piece >> 15) & NotAFile  // Down 2, right 1
		table[square] |= (piece >> 10) & NotGHFile // Down 1, left 2
		table[square] |= (piece >> 6) & NotABFile  // Down 1, right 2
	}
	return
}

// generateKingAttacks generates a lookup table for king attacks.
// Read More: https://www.chessprogramming.org/King_Pattern
func generateKingAttacks() (table [64]Bitboard) {
	for square := 0; square < 64; square++ {
		var piece Bitboard = 1 << square
		var attacks Bitboard = ShiftEast(piece) | ShiftWest(piece)
		piece |= attacks
		attacks |= ShiftNorth(piece) | ShiftSouth(piece)
		table[square] = attacks
	}
	return
}

func GenerateAttackLookupTable() map[Piece][64]Bitboard {
	AttackLookupTable := make(map[Piece][64]Bitboard)
	AttackLookupTable[Knight] = generateKnightAttacks()
	AttackLookupTable[King] = generateKingAttacks()
	return AttackLookupTable
}
