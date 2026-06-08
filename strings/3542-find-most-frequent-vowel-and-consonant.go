package strings

func MaxFreqSum(s string) int {
	hashmap := make(map[rune]int)

	for _, ch := range s {
		if v, ok := hashmap[ch]; ok {
			hashmap[ch] = v + 1
		} else {
			hashmap[ch] = 1
		}
	}

	maxFreqVowel := 0
	maxFreqConst := 0

	for ch, v := range hashmap {
		switch ch {
		case 'a':
			maxFreqVowel = max(maxFreqVowel, v)

		case 'e':
			maxFreqVowel = max(maxFreqVowel, v)

		case 'i':
			maxFreqVowel = max(maxFreqVowel, v)

		case 'o':
			maxFreqVowel = max(maxFreqVowel, v)

		case 'u':
			maxFreqVowel = max(maxFreqVowel, v)

		default:
			maxFreqConst = max(maxFreqConst, v)
		}
	}

	return maxFreqConst + maxFreqVowel
}
