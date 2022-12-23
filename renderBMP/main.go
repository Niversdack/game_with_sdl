package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	keyboard "game/KeyboardHelpers"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	gHello, err := sdl.LoadBMP("/Users/yubelyusin/go/src/game/src/MARBLES.BMP")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gHello.Free()

	window, err := sdl.CreateWindow(
		"Testing SDL2",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		gHello.ClipRect.W,
		gHello.ClipRect.H,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()
	windSurf, err := window.GetSurface()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer windSurf.Free()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}

			state := keyboard.NewState(sdl.GetKeyboardState())
			if state.KeysPressed(sdl.K_r) {
				err = gHello.UpperBlit(nil, windSurf, nil)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if state.KeysPressed(sdl.K_ESCAPE) {
				return
			}

		}
		if err = window.UpdateSurface(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
