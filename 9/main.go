package main

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}

	var re int
	y := x
	for x != 0 {
		re = x%10 + re*10
		x = x / 10
	}
	return re == y

}
