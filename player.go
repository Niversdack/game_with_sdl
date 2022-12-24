package main

import "C"
import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	sync.Mutex
	sdl.Point
	H, W      int32
	moveSpeed int32
	sdl.Color
	Health int32
}
type optFn func(player *Player)

func Move(point sdl.Point) func(player *Player) {
	return func(player *Player) {
		player.X += point.X * player.moveSpeed
		player.Y += point.Y * player.moveSpeed
	}
}

func (p *Player) Update(fns ...optFn) {
	if len(fns) == 0 {
		return
	}
	for _, fn := range fns {
		fn(p)
	}
}
func (p *Player) GetRect() *sdl.Rect {
	rect := &sdl.Rect{
		X: p.X,
		Y: p.Y,
		W: p.W,
		H: p.H,
	}
	return rect
}
