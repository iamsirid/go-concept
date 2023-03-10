package main

import "fmt"

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x = x * x
	sp.y = y * y
}

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {
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
