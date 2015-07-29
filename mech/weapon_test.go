package mech

import "testing"

func TestWeapon(t *testing.T) {
//	weapon1 := Weapon{Damage: 2, Range: 2}
	weapon1 := NewWeapon(2, 2, "test weapon1", 0.6)

	if weapon1.damage != 2 {
		t.Errorf("weapon1 damage is %d instead of 2", weapon1.damage)
	}
	if weapon1.maxRange != 2 {
		t.Errorf("weapon1 range is %d instead of 2", weapon1.maxRange)
	}
}

func TestWeaponFire(t *testing.T) {
//	weapon1 := Weapon{Damage: 2, Range: 2}
	weapon1 := NewWeapon(2, 2, "test weapon1", 0.6)

	mech1 := NewMech("testMech1", 2)
	if mech1 == nil {
		t.Errorf("mech1 was unable to be created")
	}

	weapon1.Fire(3, mech1)
	if mech1.structure != 2 {
		t.Errorf("mech destroyed at range 3 by range 2 weapon")
	}

	weapon1.Fire(2, mech1)
	if mech1.structure != 0 {
		t.Errorf("mech not destroyed at range 2 by range 2, damage 2 weapon")
	}
}
