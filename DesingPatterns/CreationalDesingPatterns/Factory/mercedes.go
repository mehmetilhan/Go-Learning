package main

type mercedes struct {
	car
}

func newMercedes() ICar {
	return &mercedes{
		car{
			name: "mercedes",
			hp: 190,
		},
	}
}


