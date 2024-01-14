package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	WhitePawn   uint64
	WhiteKnight uint64
	WhiteBishop uint64
	WhiteRook   uint64
	WhiteQueen  uint64
	WhiteKing   uint64

	BlackPawn   uint64
	BlackKnight uint64
	BlackBishop uint64
	BlackRook   uint64
	BlackQueen  uint64
	BlackKing   uint64

	Squares       [8][8]Piece
	Flags         Bits
	HalfMoveClock uint16
	FullMoveCount uint16
}

func (b *Board) SetFlag(flag Bits) {
	b.Flags |= flag
}

func (b *Board) ClearFlag(flag Bits) {
	b.Flags &= ^flag
}

func (b *Board) ToggleFlag(flag Bits) {
	b.Flags ^= flag
}

func (b *Board) HasFlag(flag Bits) bool {
	return b.Flags&flag != 0
}

func (b *Board) PutPieceAt(piece rune, rank, file int) {
	var mask uint64 = 1 << ((rank-1)*8 + file - 1)

	// Set piece on the bitboard and 8x8 board
	switch piece {
	case 'p':
		b.BlackPawn |= mask
		b.Squares[rank-1][file-1] = BlackPawn
	case 'n':
		b.BlackKnight |= mask
		b.Squares[rank-1][file-1] = BlackKnight
	case 'b':
		b.BlackBishop |= mask
		b.Squares[rank-1][file-1] = BlackBishop
	case 'r':
		b.BlackRook |= mask
		b.Squares[rank-1][file-1] = BlackRook
	case 'q':
		b.BlackQueen |= mask
		b.Squares[rank-1][file-1] = BlackQueen
	case 'k':
		b.BlackKing |= mask
		b.Squares[rank-1][file-1] = BlackKing
	case 'P':
		b.WhitePawn |= mask
		b.Squares[rank-1][file-1] = WhitePawn
	case 'N':
		b.WhiteKnight |= mask
		b.Squares[rank-1][file-1] = WhiteKnight
	case 'B':
		b.WhiteBishop |= mask
		b.Squares[rank-1][file-1] = WhiteBishop
	case 'R':
		b.WhiteRook |= mask
		b.Squares[rank-1][file-1] = WhiteRook
	case 'Q':
		b.WhiteQueen |= mask
		b.Squares[rank-1][file-1] = WhiteQueen
	case 'K':
		b.WhiteKing |= mask
		b.Squares[rank-1][file-1] = WhiteKing
	}
}

func (b *Board) Print() {
	for rank := 8; rank > 0; rank-- {
		fmt.Printf("%d ", rank)
		for _, piece := range b.Squares[rank-1] {
			char := PieceToUnicodeMap[piece]
			// highlight empty black squares
			// if char == ' ' && (rank+file+1)%2 == 0 {
			// 	char = '.'
			// }
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
	fmt.Println("  a b c d e f g h")
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
		board.SetFlag(EnPassantAvailable)
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
	// board := NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	board := NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	// board := NewBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	// board := NewBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	board.Print()
	// fmt.Printf("%032b\n", board.Flags)
	// fmt.Printf("%064b pawn\n", board.BlackPawn)
	// fmt.Printf("%064b knight\n", board.BlackKnight)
	// fmt.Printf("%064b bishop\n", board.BlackBishop)
	// fmt.Printf("%064b rook\n", board.BlackRook)
	// fmt.Printf("%064b king\n", board.BlackKing)
	// fmt.Printf("%064b queen\n", board.BlackQueen)
	// fmt.Println()
	// fmt.Printf("%064b Pawn\n", board.WhitePawn)
	// fmt.Printf("%064b Knight\n", board.WhiteKnight)
	// fmt.Printf("%064b Bishop\n", board.WhiteBishop)
	// fmt.Printf("%064b Rook\n", board.WhiteRook)
	// fmt.Printf("%064b King\n", board.WhiteKing)
	// fmt.Printf("%064b Queen\n", board.WhiteQueen)
}
