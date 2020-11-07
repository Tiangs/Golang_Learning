//Longest Substring Without Repeating Character
package main

import "fmt"

func longestNonRepeatedSubstring(s string) int {
	//Initialize an map,tracking the occurence of each character
	occr := make(map[byte]int)
	maxLength := 0
	start := 0

	for i, ch := range []byte(s) {

		// if we found this current character in the map, we want the start update to last occurence 's next
		if lastI, ok := occr[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		//check all the time if maxLength is greater
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		// fill the map
		occr[ch] = i

	}
	return maxLength

}

func main() {
	fmt.Println(longestNonRepeatedSubstring("abba"))
}
