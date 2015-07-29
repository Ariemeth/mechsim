// Project to simulate mechs to learn go.
package main

import (
	"fmt"

	"github.com/Ariemeth/mechsim/mech"
	"github.com/Ariemeth/mechsim/mech/weapon"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	mech1 := mech.NewMech("Mech1", 2)
	mech2 := mech.NewMech("Mech2", 2)

	weapon1 := weapon.Create(5, 1, "rifle", .75)
	weapon2 := weapon.Create(3, 3, "shotgun", .4)

	mech1.AddWeapon(weapon1)
	mech2.AddWeapon(weapon2)

	mech1.Fire(4, mech2)

	fmt.Println("Mech1 has ", mech1.StructureLeft(), " structure left.")
	fmt.Println("Mech2 has ", mech2.StructureLeft(), " structure left.")
}
