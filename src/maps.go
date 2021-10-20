package main

import (
	"fmt"
)

func main() {

	 /// Maps in Golang
     /// Maps are key value pairs 
     /// Syntax for creating a map in golang
	 /// map keyword followed by square brackets in sqaure brackets specify the key data type and outside the sqaure brackets specify the value data type
	 /// This is a map literal example here 
	 var m map[int]string = map[int]string{
		1 : "Hello",
		2 : "Hey",
		3 : "Hi",
	}
	/// each key value pair in the map is seperated by comma

	/// access the map values using the key
	fmt.Printf("%v\n", m[1])
	fmt.Printf("%v\n", m[2])
	fmt.Printf("%v\n", m[3])

	/// inserting into map
	m[4] = "yo"

	/// deleteing a key value pair from the map using delete() function
	/// delete function takes name of the map and the key as arguments
	delete(m, 4)

	/// check if the key is present in the map or not using ok syntax
	/// we are not using the value so we declare the variable name as _, so we don't get any variable not used error in go (check the note in the variables.go)
	/// if the key is present in the map then k is true, else it is false
	_, k := m[5]
	fmt.Printf("%v\n", k)

	/// using len() function to get length(number of key value pairs) of map
	fmt.Println(len(m))

	/// using auto syntax to create map
	newMap := map[string]string {
		"Name" : "Afroz Ahmed",
		"Hobbies" : "Reading Tech Blogs",
	}

	fmt.Println(newMap)

	/// using make function to create map
	nMap := make(map[int]string)
	fmt.Printf("%v\n", nMap)

	/// NOTE : map is also a reference type like slice 
	/// reference type means that when you copy a variable of refernce type to another variable of same type then the value is not copied, but it points to the original data container
	/// Example : copying one map into another map 
	/// making another map m2 and storing m into m3
	var m2 map[int]string = m
	/// interseting a new value into the map m2
	m2[5] = "ola"
	/// printing two maps 
	fmt.Println(m)
	fmt.Println(m2)
	/// both the maps contains the new key value pair, because they are reference type


}