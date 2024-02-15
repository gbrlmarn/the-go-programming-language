// Exercise 7.10: The sort.Interface type can be adapted to other uses. Write a function IsPalindrome(s sort.Interface) bool that reports wheter the sequence s is a palindrome, in other words, reversing the sequence would not change it. Assume that the lements at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i).
package main

import "sort"

func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for j > i {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
