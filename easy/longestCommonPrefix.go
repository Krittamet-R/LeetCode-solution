package easy

func LongestCommonPrefix(strs []string) string {
    ans := ""
	if len(strs) == 0 {
		return ""
	}
	count := 0
	for true {
		var z,y byte
		for j := 0; j<len(strs); j++ {
			x := strs[j]
			if count >= len(x) {
				return ans
			}
			y = x[count]
			if j == 0 {
				z = y
			} else if y != z && count == 0 {
				return ""
			} else if y != z{
				return ans
			}
		}
		ans = ans + string(y)
		count++
	}
	return ans
}