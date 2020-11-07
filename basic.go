package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)

}

func swap(a, b *int) {

	// fmt.Println(*a)
	// fmt.Println(*b)
	*a, *b = *b, *a

	// fmt.Println(*a)
	// fmt.Println(*b)
	// *b = *a

	// fmt.Println(*a)
	// fmt.Println(*b)
}

func printArray(arr *[5]int) {
	arr[0] = 20000
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {
	fmt.Println("hello world")
	fmt.Println("==== Pointer=====")
	variableZeroValue()
	a, b := 3, 4
	fmt.Println(&a)
	fmt.Println(&b)
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("==== Array =====")
	arr1 := [5]int{1, 2, 3, 6, 7}
	printArray(&arr1)
	fmt.Println(arr1)

	fmt.Println("==== Slice =====")
	s1 := arr1[1:3]
	s2 := s1[1:3]
	s1[0] = 133333
	fmt.Println("arr1 is")
	fmt.Println(arr1)
	fmt.Printf("s2 = %v, length s2 = %d , capacity s2 = %d\n", s2, len(s2), cap(s2))

}
