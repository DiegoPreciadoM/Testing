package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: false, speed: 100}
	prey := Prey{name: "pez payaso", speed: 10}

	// Act
	h := shark.Hunt(&prey)

	// Assert
	assert.Equal(t, false, shark.hungry, "deben ser iguales")
	assert.Equal(t, true, shark.tired, "Deben ser iguales")
	assert.Equal(t, nil, h, "deben ser iguales")
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: true, speed: 100}
	prey := Prey{name: "pez payaso", speed: 10}

	// Act
	h := shark.Hunt(&prey)

	// Assert
	assert.EqualError(t, h, "cannot hunt, i am really tired")

}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// Arrange
	shark := Shark{hungry: false, tired: false, speed: 100}
	prey := Prey{name: "pez payaso", speed: 10}

	// Act
	h := shark.Hunt(&prey)

	// Assert
	assert.EqualError(t, h, "cannot hunt, i am not hungry")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: false, speed: 50}
	prey := Prey{name: "pez payaso", speed: 100}

	// Act
	h := shark.Hunt(&prey)

	// Assert
	assert.EqualError(t, h, "could not catch it")
}

func TestSharkHuntNilPrey(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: false, speed: 50}

	// Act
	h := shark.Hunt(nil)

	// Assert
	assert.EqualError(t, h, "cannot hunt a nil prey")
}
