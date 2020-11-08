package main

import "fmt"

func main() {
	s := "This is a string"
	sChinese := "先天下之忧而忧"
	fmt.Println(s)
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X \n", b)
	}

	for i, ch := range []rune(sChinese) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

}
