package main

import "fmt"

func main()  {
	array := []int{1,2,3,4,5,6}
	var array2 []int
	array2 = append(array2, array...)
	array2 = append(array2, array)
}
