package main

import "fmt"

func intToRoman(num int) string {
	th := []string{"", "M", "MM", "MMM"}
	hu := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	te := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ge := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return th[num/1000] + hu[num%1000/100] + te[num%100/10] + ge[num%10]
}

func main() {
	num := 123
	fmt.Println(intToRoman(num))
}
