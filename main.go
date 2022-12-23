package main

// Experiment! draw some crazy stuff!
// Gist it next week and I'll show it off on stream

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	keyboard "game/KeyboardHelpers"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}

}

//var gScreenSurface sdl.Surface

func main() {
	//var gScreenSurface sdl.Surface
	// Added after EP06 to address macosx issues
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

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(gHello.ClipRect.W), int32(gHello.ClipRect.H), sdl.WINDOW_SHOWN)
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

	err = gHello.UpperBlit(nil, windSurf, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = window.UpdateSurface(); err != nil {
		fmt.Println(err)
		return
	}
	//for {
	//
	//}
	////SDL_UpdateWindowSurface( gWindow );
	//return
	//
	//renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer renderer.Destroy()
	//
	////tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	////defer tex.Destroy()
	////te
	////pixels := make([]byte, winWidth*winHeight*4)
	////
	////for y := 0; y < winHeight; y++ {
	////	for x := 0; x < winWidth; x++ {
	////		setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
	////	}
	////}
	////unsafe.Pointer
	////tex.Update(nil, pixels, winWidth*4)
	//
	//// Changd after EP 06 to address MacOSX
	//// OSX requires that you consume events for windows to open and work properly
	////g := int32(0)
	////player := &Player{
	////	Point:  sdl.Point{},
	////	H:      0,
	////	W:      0,
	////	Color:  sdl.Color{},
	////	Health: 0,
	////}
	//ticker := time.NewTicker(time.Second)
	//board := GameBoard{nil}
	for {
		//select {
		//case t := <-ticker.C:
		//	fmt.Println("Tick at", t)
		//}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			//case *sdl.KeyboardEvent:
			//	//board.KeyEvents = append(board.KeyEvents, t)
			case *sdl.QuitEvent:
				return
			}

			state := keyboard.NewState(sdl.GetKeyboardState())
			if state.KeysPressed(sdl.K_w) && state.KeysPressed(sdl.K_a) {
				fmt.Println("press ↖")

			} else if state.KeysPressed(sdl.K_w) && state.KeysPressed(sdl.K_d) {
				fmt.Println("press ↗")
			} else if state.KeysPressed(sdl.K_s) && state.KeysPressed(sdl.K_d) {
				fmt.Println("press ↘")
			} else if state.KeysPressed(sdl.K_s) && state.KeysPressed(sdl.K_a) {
				fmt.Println("press ↙")
			} else if state.KeysPressed(sdl.K_w) {
				fmt.Println("press ↑")
			} else if state.KeysPressed(sdl.K_s) {
				fmt.Println("press ↓")
			} else if state.KeysPressed(sdl.K_d) {
				fmt.Println("press →")
			} else if state.KeysPressed(sdl.K_a) {
				fmt.Println("press ←")
			} else {
				fmt.Println("not press")
			}

		}
		//time.Tick()
		//fmt.Println(board)
		//renderer.SetDrawColor(127, 255, 212, 10)
		//renderer.Copy(tex, nil, nil)
		//rect := sdl.Rect{
		//	X: int32(winWidth / 2),
		//	Y: int32(winHeight / 2),
		//	W: 100,
		//	H: 100,
		//}
		//renderer.DrawRect(&rect)
		//
		//renderer.Present()
		//
		//renderer.Clear()
	}

}

type GameBoard struct {
	//KeyEvents []*sdl.KeyboardEvent
	Keys []uint8
}
