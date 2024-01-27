package main

import (
	"fmt"
)

type Bitboard uint64

const (
	NotAFile  Bitboard = 0xfefefefefefefefe
	NotBFile           = 0xfdfdfdfdfdfdfdfd
	NotGFile           = 0xbfbfbfbfbfbfbfbf
	NotHFile           = 0x7f7f7f7f7f7f7f7f
	NotABFile          = NotAFile & NotBFile
	NotGHFile          = NotGFile & NotHFile
)

func ShiftNorth(bb Bitboard) Bitboard {
	return bb << 8
}

func ShiftSouth(bb Bitboard) Bitboard {
	return bb >> 8
}

func ShiftEast(bb Bitboard) Bitboard {
	return (bb << 1) & NotAFile
}

func ShiftWest(bb Bitboard) Bitboard {
	return (bb >> 1) & NotHFile
}

func ShiftNorthEast(bb Bitboard) Bitboard {
	return (bb << 9) & NotAFile
}

func ShiftNorthWest(bb Bitboard) Bitboard {
	return (bb << 7) & NotHFile
}

func ShiftSouthEast(bb Bitboard) Bitboard {
	return (bb >> 7) & NotAFile
}

func ShiftSouthWest(bb Bitboard) Bitboard {
	return (bb >> 9) & NotHFile
}

// PrintBitboard prints a bitboard to the console.
func PrintBitboard(bb Bitboard) {
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)
		for file := 0; file < 8; file++ {
			if bb&(1<<uint64(rank*8+file)) != 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Print("  a b c d e f g h\n\n")
}
