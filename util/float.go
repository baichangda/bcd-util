package util

import (
	"math"
)

var pows_positive [11]float64
var pows_negative [11]float64

func init() {
	for i := 0; i < len(pows_positive); i++ {
		pows_positive[i] = math.Pow(10, float64(i))
	}
	for i := 0; i < len(pows_negative); i++ {
		pows_negative[i] = math.Pow(10, float64(-i))
	}
}

func Eq(d1 float64, d2 float64, num int) bool {
	return math.Abs(d1-d2) <= pows_negative[num]
}

func Gt(d1 float64, d2 float64, num int) bool {
	return d1-d2 > pows_negative[num]
}

func Lt(d1 float64, d2 float64, num int) bool {
	return d1-d2 < pows_negative[num]
}

func Format(d float64, i int) float64 {
	if d > 0 {
		if i == 0 {
			return math.Floor(d)
		} else {
			return math.Floor(d*pows_positive[i]) / pows_positive[i]
		}

	} else if d < 0 {
		if i == 0 {
			return math.Ceil(d)
		} else {
			return math.Ceil(d*pows_positive[i]) / pows_positive[i]
		}
	} else {
		return 0
	}
}

func Round[V ~float32 | float64](v V) V {
	if v > 0 {
		if (int64(v*10) % 10) >= 5 {
			return V(int64(v) + 1)
		} else {
			return V(int64(v))
		}
	} else if v < 0 {
		if (int64(-v*10) % 10) >= 5 {
			return V(int64(v) - 1)
		} else {
			return V(int64(v))
		}

	} else {
		return 0
	}
}

func Round_slow(v float64) float64 {
	if v > 0 {
		return math.Round(v)
	} else if v < 0 {
		return -math.Round(-v)
	} else {
		return 0
	}
}
