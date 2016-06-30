package core

func MIN(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MAX(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func NP2(n uint) uint {
	if n < 2 {
		n = 2
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++

	return n
}

func PP2(n uint) uint {
	if n < 3 {
		n = 3
	}
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16

	return n - (n >> 1)
}
