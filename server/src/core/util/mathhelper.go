package util

func ClampInt32(min, max, value int32) int32 {
	if value < min {
		value = min
	}

	if value > max {
		value = max
	}

	return value
}

func ClampInt64(min, max, value int64) int64 {
	if value < min {
		value = min
	}

	if value > max {
		value = max
	}

	return value
}

func ClampFloat32(min, max, value float32) float32 {
	if value < min {
		value = min
	}

	if value > max {
		value = max
	}

	return value
}

func ClampFloat64(min, max, value float64) float64 {
	if value < min {
		value = min
	}

	if value > max {
		value = max
	}

	return value
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func MaxUInt32(a, b uint32) uint32 {
	if a > b {
		return a
	}

	return b
}

func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}

	return b
}

func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}

	return b
}

func MinUInt32(a, b uint32) uint32 {
	if a < b {
		return a
	}

	return b
}

func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}

	return b
}

func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
