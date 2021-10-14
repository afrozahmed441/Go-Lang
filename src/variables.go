package main

import (
	"fmt"
)

/// this is one of the way to create variable in GO mostly at package level (global level)

var (
	fName string = "Afroz"
	lName string = "Shaik"
	age int = 20 
)

func main(){

	/**
	- variable is like a continer which can store data and this container has a label
	  with which we can access the data stored in the continer.
	**/

	/// defining a variable
	/// we start with a keyword var 
	/// then the name of the variable in our Example it is name
	/// and then the data type (This tells which type of data we are going to store in the variable) 
	var name string = "Afroz"
	/// we are printing the variable using Println in the fmt Package 
	/// Println is a function (more on this later) 
	/// This function prints the content stored in the variable and adds a new line at the end
	fmt.Println(name)

	/// other way of defining a variable(Method 2)
	/// first defining a variable 
	/// then assign a value to the variable
	var i int
	i = 2
	fmt.Println(i)

	/// other way of defining a variable. (Method 3)
	/// This way of defining a variable only works inside any function or local scope
	/// local scope means any variable that is defined inside any function (for example main function as we are doing now)
	/// This method of defining a variable does not work at Package level (or out side any function like outside main function)
	/// This way of defining a variable does not require us to specify the data type
	/// because using this way of defining variable figures out the data stored in the variable
	/// and accordingly the compiler assigns the data type to this variable
	firstName := "Afroz"
	fmt.Println(firstName)

	/**
	Naming Variables -- Rules

	- Variable name must only be one word (no spaces)
	- Variable name can't begin with a number  
	- Variable name should only made of letters, numbers and underscore(_) no other special characters
	- Variable that are already declared cannot be used to name other variable in the same package 
	
	- Variables are Case sensitive 
	- In GoLang if a variable beings with Uppercase letter then it means that the variable is exportable
	- That means the variable is also accessible outside the package
	- If a variable beings with a lowercase letter then it means that variable is only available inside the package
	
	- Scope also plays the role in naming variables
	- If your going to use the variable in your entire program then you need to give a long and meaningful name like userName for user
	- If your only using a variable in only two or three lines of code then we need to name it short

	- and acronyms should be capitalize  ex: HTTP
	
	**/

	/**
	Data Types in GOLang

	- int : Integer data type, variable of this type can store integers. Ex: var age int = 21
	- int16 : here 16 represents the size of the integer (in bits)
	- int32 
	- int64

	- uint : Unsigned Integer data type, variable of this type can only store positive integers. Ex: var num uint = 12
	- uint16
	- uint32
	- uint64

	- float : variable of this type can store Decimals. Ex: var PI float = 3.14
	- float32
	- float64

	- bool : variable of this type can store true or false.  Ex: var isStudent bool = true

	- string : variable of this type can store text or group of characters  Ex: var name string = "Afroz"
	  NOTE: string must be in double quotes "" not single quotes ''

	- rune : variable of this type can store a single character Ex: var char rune = 'a'
	  NOTE: rune must be in single quotes ''

	- complex64 : variable of this type can store complex numbers  Ex: var num complex64 = 1+2i
	- complex128 

	**/

	/// Integers
	var age int64 = 21
	var year int64 = 2000
	var num int = 23
	fmt.Println(age)
	/// NOTE: int , int32, int64 store same data but are different in-terms of size
	/// And we can't perform operations on same data of different size. (we will see later)
	fmt.Println(year)
	fmt.Println(num)

	/// Unsigned Integers
	var num1 uint = 12
	var num2 uint16 = 14
	fmt.Println(num1)
	fmt.Println(num2)

	/// float 
	var n float32 = 2.2
	var n2 float64 = 2.3
	fmt.Println(n)
	fmt.Println(n2)

	/// bool 
	var isStudent bool = true;
	var isMarried bool = false;
	fmt.Println(isStudent)
	fmt.Println(isMarried)

	/// string
	var fName string = "Afroz"
	var lName string = "Shaik"
	fmt.Println(fName)
	fmt.Println(lName)

	/// rune
	var firstLetter rune = 'a'
	var lastLetter rune = 'z'
	fmt.Println(firstLetter)
	fmt.Println(lastLetter)

	/// NOTE : variable that are deleclared but not assigned are by default zero valued.
	/// NOTE : variable that are not used in the code cause error in GOLANG.

}