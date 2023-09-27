package main

import "fmt"

// 3. Builting and custom types
// - general types
// - structs
// - maps
// - slices
// - arrays
// - Custom types

var (
	floatVar   float32 = 0.1
	floatVar64 float64 = 0.1
	myName     string  = "Foo"
	intVar32   int32   = 1
	intVar64   int64   = 11231
	intVar     int     = 32
	uintVar    uint    = 1
	uint32Var  uint32  = 32
	uint64Var  uint64  = 434
	uint8Var   uint8   = 0x1
	byteVar    byte    = 0x2
	runVar     rune    = 'a'
)

// struct
type Player struct {
	health      int
	name        string
	attackPower float64
}

func getHealth(player Player) int {

	return player.health
}

// function receiver
func (player Player) getHealth() int {
	return player.health
}

// custom types
type Weapon string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}

func datatypes() {

	player := Player{
		name:        "foo",
		health:      100,
		attackPower: 23.34,
	}
	// maps
	users := map[string]int{"tarun": 23} // empty map
	// another way for map
	user2 := make(map[string]int)
	user2["tarun"] = 12

	// accessing the user2

	age, ok := user2["tarun"] // if key doesn't exist then it will return 0 as default value of int.

	// v -> verbose v+ -> super verbose
	fmt.Printf("this is the player %+v\n", player)
	fmt.Printf("this is the player %v\n", player)
	fmt.Printf("the health of the player is %d\n", getHealth(player))
	// using function receiver
	fmt.Printf("this return health using function receiver %d\n", player.getHealth())
	fmt.Printf("result from map %+v\n", users)
	fmt.Printf("result from map %+v\n", user2)
	fmt.Printf("result from map %d\n", age)

	if !ok {
		fmt.Printf("the value doesnt. exists")
	} else {
		fmt.Printf("the value")
	}

	// delete
	delete(users, "foo")

	// iterating over the users map
	for k, v := range users {
		fmt.Printf("the key %s the value %d\n", k, v)
	}

	// slices - dynamically grow the slice
	numbers := []int{1, 2, 3}
	otherNumber := make([]int, 2)
	fmt.Printf("numbers: %d %v\n", len(numbers), numbers)
	fmt.Printf("numbers: %d %v\n", len(otherNumber), otherNumber)

	// array is static with predefined size of array
	array := [2]int{1, 2}
	fmt.Println(array)

	// custom types

}
