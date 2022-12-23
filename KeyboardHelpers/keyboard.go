package keyboard

import (
	"github.com/veandco/go-sdl2/sdl"
)

type State struct {
	s []uint8
}
type _state interface {
	KeysPressed(keycodes ...sdl.Keycode) bool
}

func NewState(input []uint8) State {
	return State{
		s: input,
	}
}

func (s *State) check(keycode sdl.Keycode) bool {
	return s.s[int(sdl.GetScancodeFromKey(keycode))] != 1
}

func (s *State) KeysPressed(keycodes ...sdl.Keycode) bool {
	for _, keycode := range keycodes {
		if s.check(keycode) {
			return false
		}
	}
	return true
}
