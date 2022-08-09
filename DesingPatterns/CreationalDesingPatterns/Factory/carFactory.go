package main

import "fmt"

func getCar(brand string) (ICar, error) {
	if brand == "mercedes" {
		return newMercedes(), nil
	}
	if brand == "bmw" {
		return newBmw(), nil
	}

	return nil, fmt.Errorf("We haven't this brand")
}
