package main

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是
*/

// 9-回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	ans := 0
	for x > ans {
		ans = ans*10 + x%10
		x /= 10
	}
	return x == ans || x == ans/10
}
