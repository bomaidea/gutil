package strutil

import (
	"github.com/bomaidea/gutil/charutil"
)

// Reverse reverts a string ASD -> DSA, abc -> cba
func Reverse(s string) string {
	var out string

	for _, v := range s {
		out = string(v) + out
	}

	return out
}

// ContainsOnlyLetter returns true if the string
// given as parameter contains only strings
func ContainsOnlyLetter(s string) bool {
	for _, v := range s {
		if !charutil.IsLetter(v) {
			return false
		}
	}
	return true
}
