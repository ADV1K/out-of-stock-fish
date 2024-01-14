package main

// Players
type Player uint8

const (
	PlayerBlack Player = iota
	PlayerWhite
)

// Flags
type Bits uint32

const (
	WhiteToMove Bits = 1 << iota
	WhiteKingSideCastle
	WhiteQueenSideCastle
	BlackKingSideCastle
	BlackQueenSideCastle
	EnPassantAvailable
	Debug
)

// Pieces
type Piece uint8

const (
	EmptySquare Piece = iota

	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing

	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

var PieceToUnicodeMap = map[Piece]rune{
	EmptySquare: ' ',
	WhitePawn:   '♙',
	WhiteKnight: '♘',
	WhiteBishop: '♗',
	WhiteRook:   '♖',
	WhiteQueen:  '♕',
	WhiteKing:   '♔',
	BlackPawn:   '♟',
	BlackKnight: '♞',
	BlackBishop: '♝',
	BlackRook:   '♜',
	BlackQueen:  '♛',
	BlackKing:   '♚',
}
