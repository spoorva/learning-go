package main

import "fmt"

func main() {

	x := [5]int{1: 30, 3: 40, 4: 80}
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(x)

	fmt.Println("---------------Example 1--------------------")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("---------------Example 2--------------------")
	for index, element := range a {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("---------------Example 3--------------------")
	for _, value := range a {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("---------------Example 4--------------------")
	for range a {
		fmt.Println(a[j])
		j++
	}

	array1 := [5]int{11, 12, 13, 14, 15}

	array2 := &array1 //pass by refernce
	array3 := array1 //pass by value

	fmt.Println("array2 ", *array2)
	fmt.Println("array3 ", array3)
}
