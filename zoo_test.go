package main

import (
	"testing"
)

func TestNewAnimal(t *testing.T) {
	animal := NewAnimal("Guela", "Lion", 50)
	if animal == nil {
		t.Error("Expected a non-nil animal, got nil")
	}
	if animal.Name != "Guela" {
		t.Errorf("Expected animal name to be 'Guela', got '%s'", animal.Name)
	}
	if animal.Species != "Lion" {
		t.Errorf("Expected animal species to be 'Lion', got '%s'", animal.Species)
	}
	if animal.Happiness != 50 {
		t.Errorf("Expected animal happiness to be 50, got %d", animal.Happiness)
	}
}

func TestNewEnclosure(t *testing.T) {
	enclosure := NewEnclosure("Lion", true)
	if enclosure == nil {
		t.Error("Expected a non-nil enclosure, got nil")
	}
	if enclosure.AnimalSpecies != "Lion" {
		t.Errorf("Expected enclosure animal species to be 'Lion', got '%s'", enclosure.AnimalSpecies)
	}
	if !enclosure.WellMaintained {
		t.Error("Expected enclosure to be well maintained, got not well maintained")
	}
}

func TestAddAnimal(t *testing.T) {
	enclosure := NewEnclosure("Lion", true)
	animal := NewAnimal("Guela", "Lion", 50)
	enclosure.AddAnimal(animal)

	if len(enclosure.Animals) != 1 {
		t.Errorf("Expected enclosure to have 1 animal, got %d", len(enclosure.Animals))
	}
	if enclosure.Animals[0] != animal {
		t.Error("Expected enclosure's first animal to be the added animal")
	}
}

func TestFeed(t *testing.T) {
	animal := NewAnimal("Guela", "Lion", 50)
	animal.Feed()

	if animal.Happiness != 60 {
		t.Errorf("Expected animal's happiness to be 60 after feeding, got %d", animal.Happiness)
	}
}
