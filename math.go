package flex

import "math"

// from https://github.com/rkusa/gm/blob/master/math32/bits.go

const (
	uvnan = 0x7FC00001
)

var (
	NAN = math.Float32frombits(uvnan)
)

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float32 { return math.Float32frombits(uvnan) }

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float32) (is bool) {
	return f != f
}

func fmaxf(a float32, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
