package helper

import (
)


// Bitwise oring: src and bitval results the following:
// Equals to src in case bitval (all zeors except one specific position) bit is set in src.
// > src in case bitval (all zeors except one specific position) bit is not set in src.
func IsBitSet(src int, bit int) bool {
	val := src & (1 << (bit - 1))
    return (val > 0)
}


// Sets the entire source as per bitval value.
// Sets bit value to 1 in source if it's 1 at the same bit position in bitval.
func Set(src int, bitval int) int {
	return src | bitval
}


// Sets kth bit in the source.
func SetKthBit(src int, k int) int {
	return src | (1 << (k - 1))
}


// Resets src as per bitval value.
// Flips the bit value in source if it's 1 at the same bit position in bitval.
func Reset(src int, bitval int) int {
	return src ^ bitval
}


// Resets kth bit in the source.
func ResetKthBit(src int, k int) int {
	src &= ^(1 << (k - 1))
	return src
}
