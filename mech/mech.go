// Package mech project mech.go
package mech

import (
	"fmt"
)

// Mech is a basic mech type
type mech struct {
	structure    int
	maxStructure int
	weapons      []Weapon
	name         string
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(name string, maxStructure int) *mech {
	return &mech{name: name, structure: maxStructure,
		maxStructure: maxStructure}
}

// StructureLeft Retrieves the amount of remaining structure a mech has.
func (mech mech) StructureLeft() int {
	return mech.structure
}

// internal call when a mech is hit
func (mech *mech) hit(damage int) {
	mech.structure -= damage
	fmt.Println(mech.name, "takes", damage, "damage")
	if mech.structure <= 0 {
		fmt.Println(mech.name, "destroyed")
	}
}

// AddWeapon adds a Weapon to the mech
func (mech *mech) AddWeapon(weapon Weapon) {
	mech.weapons = append(mech.weapons, weapon)
}

// Fire tells the Mech to fire at a Target
func (mech *mech) Fire(rangeToTarget int, target Target) {
	for _, weapon := range mech.weapons {
		weapon.Fire(rangeToTarget, target)
	}
}
