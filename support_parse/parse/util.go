package parse

import "math"

type RoundType interface {
	~float32 | float64
}

func Round[V RoundType](v V) V {
	if v > 0 {
		if (int64(v*10) % 10) >= 5 {
			return v + 1
		} else {
			return v
		}
	} else if v == 0 {
		return 0
	} else {
		if (int64(-v*10) % 10) >= 5 {
			return v - 1
		} else {
			return v
		}
	}
}

func Round_slow(v float64) float64 {
	if v > 0 {
		return math.Round(v)
	} else if v == 0 {
		return 0
	} else {
		return -math.Round(-v)
	}
}
