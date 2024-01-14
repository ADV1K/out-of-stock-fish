package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Pawns           uint64
	Knights         uint64
	Bishops         uint64
	Rooks           uint64
	Queens          uint64
	Kings           uint64
	WhitePieces     uint64
	BlackPieces     uint64
	EnPassantSquare uint8
	HalfMoveClock   uint16
	FullMoveCount   uint16
	Flags           Flag
}

func (b *Board) SetFlag(flag Flag) {
	b.Flags |= flag
}

func (b *Board) ClearFlag(flag Flag) {
	b.Flags &= ^flag
}

func (b *Board) ToggleFlag(flag Flag) {
	b.Flags ^= flag
}

func (b *Board) HasFlag(flag Flag) bool {
	return b.Flags&flag != 0
}

func (b *Board) PutPieceAt(piece rune, rank, file int) {
	var mask uint64 = 1 << ((rank-1)*8 + file - 1)

	// Set piece on the bitboard and 8x8 board
	switch piece {
	case 'p', 'P':
		b.Pawns |= mask
	case 'n', 'N':
		b.Knights |= mask
	case 'b', 'B':
		b.Bishops |= mask
	case 'r', 'R':
		b.Rooks |= mask
	case 'q', 'Q':
		b.Queens |= mask
	case 'k', 'K':
		b.Kings |= mask
	}

	if piece >= 'a' {
		b.BlackPieces |= mask
	} else {
		b.WhitePieces |= mask
	}
}

func (b *Board) Print() {
	pieces := b.Pawns | b.Knights | b.Bishops | b.Rooks | b.Queens | b.Kings
	for i := 63; i >= 0; i-- {
		var c rune
		piece := pieces & (1 << i)

		switch {
		case piece&b.Pawns&b.WhitePieces != 0:
			c = PieceToUnicodeMap[WhitePawn]
		case piece&b.Knights&b.WhitePieces != 0:
			c = PieceToUnicodeMap[WhiteKnight]
		case piece&b.Bishops&b.WhitePieces != 0:
			c = PieceToUnicodeMap[WhiteBishop]
		case piece&b.Rooks&b.WhitePieces != 0:
			c = PieceToUnicodeMap[WhiteRook]
		case piece&b.Queens&b.WhitePieces != 0:
			c = PieceToUnicodeMap[WhiteQueen]
		case piece&b.Kings&b.WhitePieces != 0:
			c = PieceToUnicodeMap[WhiteKing]

		case piece&b.Pawns&b.BlackPieces != 0:
			c = PieceToUnicodeMap[BlackPawn]
		case piece&b.Knights&b.BlackPieces != 0:
			c = PieceToUnicodeMap[BlackKnight]
		case piece&b.Bishops&b.BlackPieces != 0:
			c = PieceToUnicodeMap[BlackBishop]
		case piece&b.Rooks&b.BlackPieces != 0:
			c = PieceToUnicodeMap[BlackRook]
		case piece&b.Queens&b.BlackPieces != 0:
			c = PieceToUnicodeMap[BlackQueen]
		case piece&b.Kings&b.BlackPieces != 0:
			c = PieceToUnicodeMap[BlackKing]

		default:
			c = PieceToUnicodeMap[EmptySquare]
		}

		if (i+1)%8 == 0 {
			if i != 63 {
				fmt.Printf("\n")
			}
			fmt.Printf("%d ", (i+1)/8)
		}
		fmt.Printf("%c ", c)
	}
	fmt.Println("\n  a b c d e f g h")
}

func NewBoard(fen string) (board Board) {
	parts := strings.Split(fen, " ")

	// Piece Placement
	rank := 8
	file := 1
	for _, c := range parts[0] {
		switch c {
		case '/':
			rank--
			file = 1
		case '1', '2', '3', '4', '5', '6', '7', '8':
			file += int(c) - '0'
		case 'p', 'n', 'b', 'r', 'q', 'k', 'P', 'N', 'B', 'R', 'Q', 'K':
			board.PutPieceAt(c, rank, file)
			file++
		default:
			panic(fmt.Sprintf("Invalid FEN: Invalid character in piece placement: '%c'", c))
		}
	}

	// Side to move: w or b
	if parts[1] == "w" {
		board.SetFlag(WhiteToMove)
	} else if parts[1] != "b" {
		panic(fmt.Sprintf("Invalid FEN: Invalid side to move: '%s'", parts[1]))
	}

	// Castling Availability: KQkq or -
	for _, c := range parts[2] {
		switch c {
		case 'K':
			board.SetFlag(WhiteKingSideCastle)
		case 'Q':
			board.SetFlag(WhiteQueenSideCastle)
		case 'k':
			board.SetFlag(BlackKingSideCastle)
		case 'q':
			board.SetFlag(BlackQueenSideCastle)
		case '-':
			continue
		default:
			panic(fmt.Sprintf("Invalid FEN: Invalid castling availability: '%c'", c))
		}
	}

	// En Passant Availability
	if parts[3] != "-" {
		board.EnPassantSquare = (parts[3][0] - 'a') + (parts[3][1]-'1')*8
	}

	// HalfMoveClock
	i, err := strconv.Atoi(parts[4])
	if err != nil {
		panic(fmt.Sprintf("Invalid FEN: Invalid value for HalfMoveClock: '%s'", parts[4]))
	}
	board.HalfMoveClock = uint16(i)

	// FullMoveCount
	i, err = strconv.Atoi(parts[5])
	if err != nil {
		panic(fmt.Sprintf("Invalid FEN: Invalid value for FullMoveCount: '%s'", parts[5]))
	}
	board.FullMoveCount = uint16(i)

	return board
}

func main() {
	board := NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	// board := NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	// board := NewBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	// board := NewBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	board.Print()
	// fmt.Printf("%032b\n", board.Flags)
	// fmt.Printf("%064b pawn\n", board.Pawns)
	// fmt.Printf("%064b knight\n", board.Knights)
	// fmt.Printf("%064b bishop\n", board.Bishops)
	// fmt.Printf("%064b rook\n", board.Rooks)
	// fmt.Printf("%064b king\n", board.Kings)
	// fmt.Printf("%064b queen\n", board.Queens)
}
