package main

// Players
type Player uint8

const (
	PlayerBlack Player = iota
	PlayerWhite
)

// Flags
type Flag uint16

const (
	WhiteToMove Flag = 1 << iota
	WhiteKingSideCastle
	WhiteQueenSideCastle
	BlackKingSideCastle
	BlackQueenSideCastle
	Debug
)

// Pieces
type Piece uint8

const (
	EmptySquare Piece = 0

	Pawn   = 1
	Knight = 2
	Bishop = 3
	Rook   = 4
	Queen  = 5
	King   = 6

	WhitePiece = 8
	BlackPiece = 16
)

const (
	WhitePawn   = WhitePiece | Pawn
	WhiteKnight = WhitePiece | Knight
	WhiteBishop = WhitePiece | Bishop
	WhiteRook   = WhitePiece | Rook
	WhiteQueen  = WhitePiece | Queen
	WhiteKing   = WhitePiece | King

	BlackPawn   = BlackPiece | Pawn
	BlackKnight = BlackPiece | Knight
	BlackBishop = BlackPiece | Bishop
	BlackRook   = BlackPiece | Rook
	BlackQueen  = BlackPiece | Queen
	BlackKing   = BlackPiece | King
)

var PieceToUnicodeMap = map[Piece]rune{
	EmptySquare: '.',
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
