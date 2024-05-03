package main

import "fmt"

func main() {
	guela := NewAnimal("Guela", "Lion", 50)
	fmt.Println("New animal created:", guela)

	enclosure := NewEnclosure("Lion", true)
	fmt.Println("New enclosure created:", enclosure)

	enclosure.AddAnimal(guela)
	fmt.Println("Animal added to the enclosure:", enclosure)

	guela.Feed()
	fmt.Println("Animal fed, current happiness:", guela.Happiness)

	enclosure.AttractVisitors()
	fmt.Printf("Enclosure attracted %d visitors\n", enclosure.Visitors)

	revenue := enclosure.CalculateRevenue()
	fmt.Printf("Revenue generated from visitors: $%d\n", revenue)
}
