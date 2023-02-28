package inheritance

import (
	"fmt"
	"testing"
)

func TestEmbbedding(t *testing.T) {
	enemy := NewEnemy()
	enemy.Move(1, 2)
	fmt.Println("enemy: ", enemy.Position)
	enemy.MoveSpecial(2, 3)
	fmt.Println("enemy: ", enemy.Position)

	player := NewPlayer()
	player.Move(1, 1)
	player.Teleport(2, 2)

	fmt.Println("player: ", player.Position)
}
