package medium

import "fmt"

// a-z = 97-122
func LengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	re := 0
	start := 0
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
		_, dup := seen[s[i]]
		if !dup {
			seen[s[i]] = i
		}

		if dup {
			if len(seen) > re {
				re = len(seen)
			}
			val, _ := seen[s[i]]
			for j := start; j <= val; j++{
				delete(seen, s[j])
				start = val + 1	
			}
			seen[s[i]] = i
		}
		fmt.Printf(" %d\n", len(seen))
	}
	if len(seen) > re {
		re = len(seen)
	}
	fmt.Println(re)
	return re
}
