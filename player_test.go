package main

import (
	"fmt"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestName(t *testing.T) {
	p := Player{}
	point := sdl.Point{
		X: -10,
		Y: 10,
	}
	fmt.Print(p)
	p.Update(Move(point), Move(point), Move(point))

	fmt.Print(p)
}
