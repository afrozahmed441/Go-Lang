package main

import (
	"fmt"
)

/// one way of defining constants at package level (global level)

const (
	PI = 3.14
	YEAR = 365
)

/// enum constants
/// ioat assigns value 0 for first time (starts from 0)
const (
	a = iota
	b
	c
	d
)

func main() {

	/// Constants in GO
	/// There are Two Types of Constant in GO
	/// Untyped Constants and Typed Constants

	/// UnTyped Constants
	const DOB = 2000
	const PI = 3.14
	const NAME = "afroz"

	fmt.Println(DOB);
	fmt.Println(PI);
	fmt.Println(NAME);

	/// Typed Constants
	const YEAR int32 = 2021
	const SUM float32 = 7.14
	const SNAME string = "shaik";

	fmt.Println(YEAR);
	fmt.Println(SUM);
	fmt.Println(SNAME);

	/// enum constants
	fmt.Printf("%v\n", a);
	fmt.Printf("%v\n", b);
	fmt.Printf("%v\n", c);
	fmt.Printf("%v\n", d);

/// NOTE : Arrays and slices can't be assigned as const (we will see in array and slice section)
}