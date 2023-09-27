package main

import "fmt"

// Enums - by default it not available in go with tweaks we can achieve that.

// weapon type  -> axe, sword, wooden stick, knife

// without enum
func getDamage(weaponType string) int {
	switch weaponType {
	case "axe":
		return 100
	case "sword":
		return 89
	case "wooden stick":
		return 40
	case "knife":
		return 99
	default:
		panic("weapon does not exists") // after encounter it stops the execution of function with a stack trace
	}
}

// with enums

type WeaponType int

// const (
//
//	Axe         WeaponType = 1
//	Sword       WeaponType = 2
//	WoodenStick WeaponType = 3
//	Knife       WeaponType = 4
//
// )

// iota helps to increase the value one after another in the type
const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func getDamage2(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 89
	case WoodenStick:
		return 40
	case Knife:
		return 99
	default:
		panic("weapon does not exists") // after encounter it stops the execution of function with a stack trace
	}
}

func enums() {

	fmt.Println("damage of the given weapon: ", getDamage("knife"))
	fmt.Println("damage of the given weapon: ", getDamage2(Knife))
	fmt.Printf("damage of the given weapon: %d %d\n", Knife, getDamage2(Knife))
}
