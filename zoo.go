package main

import "fmt"

type Animal struct {
	Name      string
	Species   string
	Happiness int
}

type Enclosure struct {
	AnimalSpecies  string
	Animals        []*Animal
	WellMaintained bool
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
	return fmt.Sprintf("AnimalSpecies: %s, Animals: %v, WellMaintained: %t", e.AnimalSpecies, e.Animals, e.WellMaintained)
}

func (e *Enclosure) AddAnimal(animal *Animal) {
	if animal.Species == e.AnimalSpecies {
		e.Animals = append(e.Animals, animal)
	}
}

func (a *Animal) Feed() {
	a.Happiness += 10
}
