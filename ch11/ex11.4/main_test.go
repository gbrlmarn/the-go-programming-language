// Exercise 11.4: Modify randomPalindrome to exercise IsPalindrome's handling of punctuation and spaces
package main

import (
	"math/rand"
	"testing"
	"time"

	word "gopl/ch11/examples/word2"
)

// randomPalindrome generates a palindrome using only punctuation and spaces
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		// [0x0020 - 0x0030] contains spaces and punctuation
		r := rune(0x0020 + rng.Intn(0x0010))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number genrator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false\n", p)
		}
	}
}
