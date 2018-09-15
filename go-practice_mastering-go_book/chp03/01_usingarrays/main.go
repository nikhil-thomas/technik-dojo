package main

import (
	"fmt"
)

func main() {
	arr1 := [2]int{1, 2}
	arr2 := [2][2]int{{1, 2}, {3, 4}}
	arr3 := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("len arr1:", len(arr1))
	fmt.Println("len arr2:", len(arr2))
	fmt.Println("len arr3:", len(arr3))

	fmt.Println("arr1:", arr1)
	fmt.Println("arr1[0]:", arr1[0])
	fmt.Println("arr2:", arr2)
	fmt.Println("arr2[0]:", arr2[0])
	fmt.Println("arr2[0][0]:", arr2[0][0])
	fmt.Println("arr3:", arr3)
	fmt.Println("arr3[0]:", arr3[0])
	fmt.Println("arr3[0][0]:", arr3[0][0])
	fmt.Println("arr3[0][0][0]:", arr3[0][0][0])

	fmt.Printf("\narr3\n")
	for _, arr2d := range arr3 {
		for _, arr1d := range arr2d {
			for _, v := range arr1d {
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}
