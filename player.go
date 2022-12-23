package main

import "C"
import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	sync.Mutex
	sdl.Point
	H, W int32
	sdl.Color
	Health int32
}
type optFn func(player *Player)

func Move(point sdl.Point) func(player *Player) {
	return func(player *Player) {
		player.X += point.X
		player.Y += point.Y
	}
}

func (p *Player) Update(fns ...optFn) {
	for _, fn := range fns {
		fn(p)
	}
}
