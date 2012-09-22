// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

import "math"

func max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

func min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

func abs(a float32) float32 {
	return float32(math.Abs(float64(a)))
}

func sqrt(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}

func sin(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func cos(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func tan(a float32) float32 {
	return float32(math.Tan(float64(a)))
}

func asin(a float32) float32 {
	return float32(math.Asin(float64(a)))
}

func acos(a float32) float32 {
	return float32(math.Acos(float64(a)))
}

func atan(a float32) float32 {
	return float32(math.Atan(float64(a)))
}
