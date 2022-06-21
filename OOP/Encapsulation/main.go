package main

import (
	"fmt"
	"github.com/golearning/OOP/Encapsulation/models"
)

//Test function
func main() {
	//STRUCTURE IDENTIFIER
	p := &models.Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &models.company{}
	fmt.Println(c)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//FUNCTION
	person2 := models.GetPerson()
	fmt.Println(person2)
	companyName := models.getCompanyName()
	fmt.Println(companyName)

	//VARIBLES
	fmt.Println(models.companyLocation)
	fmt.Println(models.CompanyName)
}
