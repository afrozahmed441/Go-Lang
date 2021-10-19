package main

import (
	"fmt"
)

func main() {

	/// arrays in golang
	/// arrays are fixed size in golang
	/// creating a array of size 3
	var arr [3]int64 = [3]int64{1, 2, 3}
	/// printing the array
	fmt.Println(arr)

	/// 2nd Method or way of creating array
	var arr2 [3]string
	/// we can access the elements in the array using index values and assign or reassign the elements 
	/// index value starts from 0 (same as most programming lang's)
	arr2[0] = "Afroz"
	arr2[1] = "Josh"
	arr2[2] = "Mark"

	fmt.Println(arr2)

	/// 3rd Method or way of creating array
	/// this syntax automatically figures out the type of array (kind of like auto in c++)
	arr3 := [3]float32{1.1, 2.2, 3.3}
	fmt.Println(arr3)

	/// find the length of the array using the len() function
	fmt.Println(len(arr2))

	/// multi-diminsional arrays
	var mat [3][3]int64 = [3][3]int64{
		[3]int64{1, 0, 0},
		[3]int64{0, 1, 0},
		[3]int64{0, 0, 1}}

	/// printing the multi-d array
	fmt.Println(mat)
	fmt.Println(mat[0])
	fmt.Println(mat[1])
	fmt.Println(mat[2])

	/// In golang arrays are values just like variables
	/// arrays are primitives in golang
	/// when we assign a array to another array, copy of the array is created and assigned to the array
	a := arr
	a[0] = 2
	/// original array
	fmt.Println(arr)
	/// copy array
	fmt.Println(a)

	/// when you don't want to create a copy then you can use & to copy the address of the array
	/// It is like creating a alias name for array a, if any element is modified in the array b then it is reflected into the array a also.
	b := &a
	b[2] = 4
	/// original array
	fmt.Println(a)
	/// copy using the & syntax
	fmt.Println(b)

	///// ------------------ Slice in GOLang ----------------------------------

	/// slices are dynamic size arrays in the golang
	/// slice syntax
	var s []int64 = []int64{1, 2, 3}
	fmt.Println(s)

	/// slice has a capcity, which tells how many elements it can store before dubbling the slice
	fmt.Println(cap(s))

	/// slices are reference type, it means that when we assign one slice to another slice it does not create a copy it copies the address of the slice
	s1 := s
	/// change in the one slice also reflects in the original slice that is used to make copy 
	s1[1] = 5
	fmt.Println(s)
	fmt.Println(s1)
	
	/// using make() function to create a slice
	/// make(type, size, capacity)
	s2 := make([]int64, 3, 10)
	/// by default it contains zero values (THIS one of the feature of GOLang)
	fmt.Println(s2)

	/// we can also create a empty slice and add elements using the append() function
	/// append() function returns new slice after appending the element to slice

	var s3 []int = []int{}
	s3 = append(s3, 1)
	s3 = append(s3, 2)
	s3 = append(s3, 3)
	fmt.Println(s3)

	/// we also insert a slice into another slice using ...(spread operator) 
	/// but it must be slice literal
	s3 = append(s3, []int{4, 5, 6, 7}...)
	fmt.Println(s3)

	/// using slice operator to slice the slices :)
	/// This will print the elements from 1st index to 3rd index
	fmt.Println(s3[1:4])
	/// This will print the elemtns from 5th index till end index
	fmt.Println(s3[4:])
	/// This will print all the elements in the slice 
	fmt.Println(s3[:])

	/// basic slice operations
	/// pop operation
	s3 = s3[1:]
	/// push 
	s3 = append(s3, 8)
	/// remove element from back
	s3 = s3[:len(s3)-1]
	/// remove elements from middle of the slice
	s3 = append(s3[:2], s3[3:]...)
	fmt.Println(s3)
}