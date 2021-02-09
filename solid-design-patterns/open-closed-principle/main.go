package main

import "fmt"

const (
	spear       = "Spear"
	sword       = "Sword"
	dragger     = "Dragger"
	bowAndArrow = "Bow and Arrow"
)

type Weapon string

type Weapons interface {
	useWeapon() Weapon
}

type King struct {
	weapon Weapon
}

func (k *King) useWeapon() Weapon {
	return k.weapon
}

type Queen struct {
	weapon Weapon
}

func (q *Queen) useWeapon() Weapon {
	return q.weapon
}

type Soldiers struct {
	weapon Weapon
}

func (s *Soldiers) useWeapon() Weapon {
	return s.weapon
}

type Fight struct{}

func (f *Fight) fight(w Weapons) {
	fmt.Println(fmt.Sprintf("%T fighting with weapon %s", w, w.useWeapon()))
}

func main() {
	f := &Fight{}
	k := &King{sword}
	q := &Queen{dragger}
	s := &Soldiers{bowAndArrow}
	f.fight(k)
	f.fight(q)
	f.fight(s)

	k = &King{spear}
	f.fight(k)

}
