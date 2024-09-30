package mathutil

func Add(a int16, b int16) int16 {
	return a + b
}

func Subtract(a int16, b int16) int16 {
	return a - b
}

func Multiply(a int16, b int16) int16 {
	return a * b
}

func Divide(a int16, b int16) int16 {
	return a / b
}

func Remainder(a int16, b int16) int16 {
	return a % b
}

func Min(a int16, b int16) int16 {
	if a < b {
		return a
	}
	return b
}
func Max(a int16, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
