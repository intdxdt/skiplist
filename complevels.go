package skiplist

// compute number of levels
func computeLevels(n int) int {
	var levels = 1
	for n > 0 {
		levels += 1
		n >>= 2
	}
	return levels
}
