package main

import "fmt"

func main() {
	var (
		i ItalianCoffeeMachine
		c ColombiamCoffeeMachine
	)

	//Se pone el & para pasar la dirección de memoria de esa estructura
	//de lo contrario falla.
	italianCoffee := GetCoffee(&i, 10)
	italianCoffee.CoffeePrint()
	fmt.Println("**********************")
	colombiamCoffee := GetCoffee(&c, 20)
	colombiamCoffee.CoffeePrint()
}

type Coffee struct {
	Intensity int
	Region    string
}

//CoffeeMaker Interface
type CoffeeMaker interface {
	MakeCoffee(intensity int) Coffee
}

type ItalianCoffeeMachine struct{}

type ColombiamCoffeeMachine struct{}

func (c *Coffee) CoffeePrint() {
	fmt.Println(fmt.Sprintf("This coffee is from %s and instensity is %d", c.Region, c.Intensity))
}

func (i *ItalianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Italy"}
}

func (c *ColombiamCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Colombia"}
}

func GetCoffee(coffeeMaker CoffeeMaker, i int) Coffee {
	return coffeeMaker.MakeCoffee(i)
}
