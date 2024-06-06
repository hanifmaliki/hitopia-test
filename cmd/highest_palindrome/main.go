package main

import (
	"fmt"
)

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

func highestPalindrome(s string, k int) string {
	str := []byte(s)
	n := len(str)

	// Declare function first and assign the function after it, so it can call inside the function itself
	var makePalindrome func(left, right int) bool
	makePalindrome = func(left, right int) bool {
		if left >= right {
			return true
		}
		if str[left] != str[right] {
			str[left] = max(str[left], str[right])
			str[right] = str[left]
			k--
		}
		if k < 0 {
			return false
		}
		return makePalindrome(left+1, right-1)
	}

	if !makePalindrome(0, n-1) {
		return "-1"
	}

	left, right := 0, n-1
	var maximizePalindrome func(left, right int)
	maximizePalindrome = func(left, right int) {
		if left > right {
			return
		}
		if left == right {
			if k > 0 {
				str[left] = '9'
			}
			return
		}
		if str[left] < '9' {
			if k >= 2 && str[left] == s[left] && str[right] == s[right] {
				k -= 2
				str[left], str[right] = '9', '9'
			} else if k >= 1 && (str[left] != s[left] || str[right] != s[right]) {
				k--
				str[left], str[right] = '9', '9'
			}
		}
		maximizePalindrome(left+1, right-1)
	}

	maximizePalindrome(left, right)

	return string(str)
}

func main() {
	fmt.Println(highestPalindrome("3943", 1))          // Output: "3993"
	fmt.Println(highestPalindrome("932239", 2))        // Output: "992299"
	fmt.Println(highestPalindrome("12321", 1))         // Output: "12921"
	fmt.Println(highestPalindrome("12345", 1))         // Output: "-1"
	fmt.Println(highestPalindrome("5533325233355", 5)) // Output: "9933329233399"
	fmt.Println(highestPalindrome("12345", 3))         // Output: "94349"
	fmt.Println(highestPalindrome("12345", 6))         // Output: "99999" // Cukup 5 sebenernya, tapi di peraturan tidak ada keterangan jadi -1 kalo k > length
}
