package main

type bmw struct {
	car
}

func newBmw() ICar {
	return &bmw{
		car{
			name: "bmw",
			hp:   231,
		},
	}
}
