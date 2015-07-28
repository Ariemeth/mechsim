package mech

// Weapon is weapon with specific characteristics
type Weapon struct {
	Range, Damage int
	Name          string
	HitRate       float64
}

// Target is an interface used by objects that can be hit and take damage
type Target interface {
	hit(int)
}

// Fire is used by an object to fire at a Target.
// Requires the range to the Target and the Target.
func (weapon Weapon) Fire(rangeToTarget int, target Target) {
	if rangeToTarget <= weapon.Range {
		target.hit(weapon.Damage)
	}
}
