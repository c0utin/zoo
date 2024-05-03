package main

import (
	"fmt"
)

type Animal struct {
	Name      string
	Species   string
	Happiness int
}

type Enclosure struct {
	AnimalSpecies  string
	Animals        []*Animal
	WellMaintained bool
	Visitors       int
}

func NewAnimal(name, species string, happiness int) *Animal {
	return &Animal{
		Name:      name,
		Species:   species,
		Happiness: happiness,
	}
}

func NewEnclosure(animalSpecies string, wellMaintained bool) *Enclosure {
	return &Enclosure{
		AnimalSpecies:  animalSpecies,
		WellMaintained: wellMaintained,
	}
}

func (a *Animal) String() string {
	return fmt.Sprintf("Name: %s, Species: %s, Happiness: %d", a.Name, a.Species, a.Happiness)
}

func (e *Enclosure) String() string {
	return fmt.Sprintf("AnimalSpecies: %s, Animals: %v, WellMaintained: %t, Visitors: %d", e.AnimalSpecies, e.Animals, e.WellMaintained, e.Visitors)
}

func (e *Enclosure) AddAnimal(animal *Animal) {
	if animal.Species == e.AnimalSpecies {
		e.Animals = append(e.Animals, animal)
	}
}

func (a *Animal) Feed() {
	a.Happiness += 10
}

func (e *Enclosure) AttractVisitors() {
	if e.WellMaintained && e.calculateAverageHappiness() >= 50 {
		e.Visitors = 10
	} else {
		e.Visitors = 0
	}
}

func (e *Enclosure) calculateAverageHappiness() int {
	if len(e.Animals) == 0 {
		return 0
	}
	totalHappiness := 0
	for _, animal := range e.Animals {
		totalHappiness += animal.Happiness
	}
	return totalHappiness / len(e.Animals)
}

func (e *Enclosure) CalculateRevenue() int {
	return e.Visitors * 5
}
