package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	keyboard "game/KeyboardHelpers"
)

const winWidth, winHeight int32 = 800, 600

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()
	rwFile := sdl.RWFromFile("/Users/yubelyusin/go/src/game/src/mario_PNG75.png", "rb")
	loadSurface, err := img.LoadPNGRW(rwFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create new surface with format loaded files
	optimizeSurface, err := loadSurface.Convert(loadSurface.Format, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer optimizeSurface.Free()

	loadSurface.Free()
	// stretchRect for change size optimize surface

	var stretchRect = &sdl.Rect{
		X: winWidth/2 - (optimizeSurface.W/10)/2,
		Y: winHeight/2 - (optimizeSurface.H/10)/2,
		W: optimizeSurface.H / 10,
		H: optimizeSurface.W / 10,
	}
	fmt.Println(stretchRect)
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
	// append Surface with stretchRect to Window
	err = optimizeSurface.BlitScaled(nil, windSurf, stretchRect)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
			state := keyboard.NewState(sdl.GetKeyboardState())
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
