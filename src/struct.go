package main

import (
	"fmt"
)

/// Structures in GoLang
/// structures in GoLang are used to create user-defined data types (just like in c or c++)
/// syntax: type structName struct { body }
type User struct {
	id uint64
	firstName string
	lastName string
	age int16
}


func main() {

	/// create a variable of type User
	/// best practice is to also specify the key and value while defining the structure values
	var user User = User{id : 1, firstName: "Afroz", lastName: "shaik", age : 21}
	fmt.Println(user)

	/// 2nd way of defining a variable of type User
	/// in this syntax we must specify all the values for each key, if any of them mismatches with data type then there will be error
	user2 := User{1, "Afroz", "shaik", 21}
	fmt.Println(user2)

	/// access the data from the structure 
	fmt.Printf("%v\n", user.id)
	fmt.Printf("%v\n", user.firstName)
	fmt.Printf("%v\n", user.lastName)
	
	
}