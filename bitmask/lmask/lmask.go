package lmask

import "strconv"

// LMask ...
type LMask []uint

const (
	Bits = strconv.IntSize // TODO: Remove dependency
)

var (
	U0 LMask = LMask{}
	U1 LMask = LMask{1}
)
