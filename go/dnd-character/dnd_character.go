package dndcharacter

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	dieBox := make([]int, 4)
	for i := 0; i < 4; i++ {
		dieBox[i] = 1 + rand.IntN(6)
	}
	slices.Sort(dieBox)
	abilityScore := 0
	for i := 1; i < 4; i++ {
		abilityScore += dieBox[i]
	}
	return abilityScore
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Ability(),
		Ability(),
		constitution,
		Ability(),
		Ability(),
		Ability(),
		10 + Modifier(constitution),
	}
}
