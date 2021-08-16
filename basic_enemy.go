package main
import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)
const basicEnemySize = 105
func newBasicEnemy(renderer *sdl.Renderer, x, y float64) *element {
	basicEnemy := &element{}
	basicEnemy.position = vector{x: x, y: y}
	basicEnemy.rotation = 180
	idleSequence, err := newSequence("img/basic_enemy/idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence("img/basic_enemy/destroy", 60, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}
	sequences := map[string]*sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}
	animator := newAnimator(basicEnemy, sequences, "idle")
	basicEnemy.addComponent(animator)
	vtb := newVulnerableToBullets(basicEnemy)
	basicEnemy.addComponent(vtb)
	col := circle{ center: basicEnemy.position, radius: 38 }
	basicEnemy.collisions = append(basicEnemy.collisions, col)
	basicEnemy.active = true
	return basicEnemy
}