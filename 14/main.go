package main

func longestCommonPrefix(strs []string) string {
	le := len(strs)

	if le == 1 {
		return strs[0]
	}

	first := strs[0]
	for i := 1; i < le; i++ {
		l1 := len(first)
		l2 := len(strs[i])
		var temp string
		for m := 0; m < l1 && m < l2; m++ {
			if first[m] == strs[i][m] {
				temp = temp + string(first[m])
			} else {
				break
			}
		}
		first = temp
	}
	return first
}
