package medium

import "fmt"

// a-z = 97-122
func LengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	re := 0
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
			Val, _ := seen[s[i]]
			i = Val
			seen = map[byte]int{}
		}
		fmt.Printf(" %d\n", len(seen))
	}
	if len(seen) > re {
		re = len(seen)
	}
	fmt.Println(re)
	return re
}
