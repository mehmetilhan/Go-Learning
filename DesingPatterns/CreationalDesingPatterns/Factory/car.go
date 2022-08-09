package main

type car struct {
	name string
	hp   int
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) setHP(power int) {
	c.hp = power
}

func (c *car) getHP() int {
	return c.hp
}
