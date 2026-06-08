package strings

func BalancedStringSplit(s string) int {
	rc := 0
	lc := 0

	bc := 0

	for _, ch := range s {
		if ch == 'L' {
			lc++
		} else {
			rc++
		}

		if (rc != 0 && lc != 0) && rc == lc {
			bc++
			rc = 0
			lc = 0
		}
	}
	return bc
}
