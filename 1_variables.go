package main

import "fmt"

// variables
// - global variables
// - local variables
// - constants
// - variable declarations

var name = "Foo"
var firstName string = "Foo"
var lastName string = "day"

var (
	firstName2 = "Foo"
	lastName2  = "Foo"
)

const version3 = 1 // always declared on the top of package. always lowercase

func variable() {

	// local score of variable
	var version3 int
	version := 1           //infer the type
	otherVersion := "bar"  //infer string
	anotherVersion := 10.1 // infer to float32
	fmt.Println("first name: ", firstName, version, otherVersion, anotherVersion, version3)

	// constants are immutable
	fmt.Println(version3)
}
