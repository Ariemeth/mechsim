package mech

// Weapon is weapon with specific characteristics
type weapon struct {
	maxRange, damage int
	name          string
	hitRate       float64
}

// Target is an interface used by objects that can be hit and take damage
type Target interface {
	hit(int)
}

// NewWeapon creates a new Weapon.
func NewWeapon(maxRange int, damage int, name string,
	hitRate float64) weapon {

	return weapon{maxRange: maxRange, damage: damage, name: name,
		hitRate: hitRate}
}

// Fire is used by an object to fire at a Target.
// Requires the range to the Target and the Target.
func (weapon weapon) Fire(rangeToTarget int, target Target) {
	if rangeToTarget <= weapon.maxRange {
		target.hit(weapon.damage)
	}
}
