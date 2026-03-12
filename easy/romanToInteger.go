package easy

// very terrible memory usage despite good runtime
func RomanToInt(s string) int {
	t := make([]int, len(s))
	save := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			t[i] = 1
		case 'V':
			t[i] = 5
		case 'X':
			t[i] = 10
		case 'L':
			t[i] = 50
		case 'C':
			t[i] = 100
		case 'D':
			t[i] = 500
		case 'M':
			t[i] = 1000
		}
	}
	for i := 1; i < len(s); i++ {
		if t[i-1] > t[i] {
			save += t[i-1]
		} else if t[i-1] < t[i] {
			save -= t[i-1]
		} else {
			save += t[i-1]
		}
	}
	save += t[len(s)-1]

	return save
}
