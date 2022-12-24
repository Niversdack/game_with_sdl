package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
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
	rwFile := sdl.RWFromFile("/Users/yubelyusin/go/src/game/src/texture.png", "rb")
	pngLoaded, err := img.LoadPNGRW(rwFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	rect := pngLoaded.ClipRect
	texture, err := renderer.CreateTextureFromSurface(pngLoaded)
	if err != nil {
		fmt.Println(err)
		return
	}
	pngLoaded.Free()
	defer texture.Destroy()
	RenderBackGround(renderer, rect, texture, winWidth, winHeight)

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
				w, h := window.GetSize()
				RenderBackGround(renderer, rect, texture, w, h)
			}
			if state.KeysPressed(sdl.K_ESCAPE) {
				return
			}
		}
		//if err = window.UpdateSurface(); err != nil {
		//	fmt.Println(err)
		//	return
		//}
	}
}

// Fill textures all back ground with W,H
func RenderBackGround(renderer *sdl.Renderer, rect sdl.Rect, texture *sdl.Texture, w, h int32) {
	renderer.Clear()
	src := &sdl.Rect{
		X: rect.X,
		Y: rect.Y,
		W: rect.W / 2,
		H: rect.H / 2,
	}
	for {
		if src.Y > h {
			break
		}
		for {
			if src.X > w {
				src.X = 0
				break
			}
			renderer.Copy(texture, nil, src)
			src.X = src.X + src.W
		}
		src.Y = src.Y + src.H
	}
	renderer.Present()
}
