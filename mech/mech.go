// Package mech project mech.go
package mech

import (
	"fmt"
)

// Mech is a basic mech type
type Mech struct {
	structure int
	weapons   []Weapon
	name      string
}

// StructureLeft Retrieves the amount of remaining structure a mech has.
func (mech Mech) StructureLeft() int {
	return mech.structure
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(name string) *Mech {
	newMech := new(Mech)
	newMech.structure = 2
	newMech.name = name

	return newMech
}

// internal call when a mech is hit
func (mech *Mech) hit(damage int) {
	mech.structure -= damage
	fmt.Println(mech.name, "takes", damage, "damage")
	if mech.structure <= 0 {
		fmt.Println(mech.name, "destroyed")
	}
}

// AddWeapon adds a Weapon to the mech
func (mech *Mech) AddWeapon(weapon Weapon) {
	mech.weapons = append(mech.weapons, weapon)
}

// Fire tells the Mech to fire at a Target
func (mech *Mech) Fire(rangeToTarget int, target Target) {
	for _, weapon := range mech.weapons {
		weapon.Fire(rangeToTarget, target)
	}
}
