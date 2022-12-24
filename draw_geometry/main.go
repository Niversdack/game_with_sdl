package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	keyboard "game/KeyboardHelpers"
)

const winWidth, winHeight int32 = 1920, 1080

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()
	window, err := sdl.CreateWindow(
		"Mario Mushroom",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth,
		winHeight,
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

	renderer, err := window.GetRenderer()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = renderer.SetDrawColor(0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	renderer.Clear()
	fillRect := &sdl.Rect{
		X: winWidth / 4,
		Y: winHeight / 4,
		W: winWidth / 2,
		H: winHeight / 2,
	}
	renderer.SetDrawColor(214, 59, 139, 0)
	renderer.FillRect(fillRect)

	outlineRect := &sdl.Rect{
		X: winWidth / 6,
		Y: winHeight / 6,
		W: winWidth * 2 / 3,
		H: winHeight * 2 / 3,
	}
	renderer.SetDrawColor(211, 231, 20, 0)
	renderer.DrawRect(outlineRect)
	renderer.SetDrawColor(255, 130, 0, 0)
	renderer.DrawLine(0, winHeight/2, winWidth, winHeight/2)
	//SDL_SetRenderDrawColor( gRenderer, 0xFF, 0xFF, 0x00, 0xFF );
	renderer.SetDrawColor(255, 130, 0, 0)

	//for( int i = 0; i < SCREEN_HEIGHT; i += 4 )
	//{
	for i := int32(0); i < winHeight; i += 4 {
		renderer.DrawPoint(winWidth/2, i)
	}

	renderer.Present()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
			state := keyboard.NewState(sdl.GetKeyboardState())
			if state.KeysPressed(sdl.K_SPACE) {
				err = window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if state.KeysPressed(sdl.K_ESCAPE) {
				return
			}
		}

	}
}
