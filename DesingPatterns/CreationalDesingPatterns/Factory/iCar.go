package main

type ICar interface {
	setName(name string)
	setHP(power int)
	getName() string
	getHP() int
}
