package mech

import "testing"

func TestNewMech(t *testing.T) {

	//	weapon1 := Weapon{Damage:2, Range:2}
	//	weapon2 := Weapon{Damage:1, Range:4}

	mech1 := NewMech("testMech1")
	if mech1 == nil {
		t.Errorf("mech1 was unable to be created")
	}

	mech2 := NewMech("testMech2")
	if mech2 == nil {
		t.Errorf("mech2 was unable to be created")
	}
}

func TestHit(t *testing.T) {

	//	weapon1 := Weapon{Damage:2, Range:2}

	mech1 := NewMech("testMech1")
	if mech1 == nil {
		t.Errorf("mech1 was unable to be created")
	}

	mech1.hit(0)
	if mech1.structure != 2 {
		t.Errorf("mech1 took damage when it was hit with 0")
	}

	mech1.hit(2)
	if mech1.structure != 0 {
		t.Errorf("mech1 was not destroyed by taking 2 damage")
	}
}

func TestAddWeapon(t *testing.T) {

}

func TestMechFire(t *testing.T) {

}
