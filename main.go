package main

// Experiment! draw some crazy stuff!
// Gist it next week and I'll show it off on stream

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	keyboard "game/KeyboardHelpers"
)

const winWidth, winHeight int32 = 800, 600

type color struct {
	r, g, b byte
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

	player := &Player{
		moveSpeed: 5,
		Point: sdl.Point{
			X: int32(winWidth / 2),
			Y: int32(winHeight / 2),
		},
		H:      50,
		W:      50,
		Color:  sdl.Color{},
		Health: 0,
	}
	ticker := time.NewTicker(time.Millisecond * 16)
	board := &GameBoard{}

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
	for {

		select {
		case t := <-ticker.C:
			err = renderer.SetDrawColor(0, 0, 0, 0)
			renderer.Clear()
			renderer.SetDrawColor(255, 0, 0, 0)
			rect := player.GetRect()
			err = renderer.FillRect(rect)
			if err != nil {
				fmt.Println(err)
			}
			renderer.Present()
			fmt.Println(player.Point)

			tmp := [maxEventLen]optFn{}
			copy(tmp[:], board.events)
			board.events = nil
			//!!!
			go func(events [maxEventLen]optFn) {
				move := make([]optFn, 0, 5)
				for i := 0; i < len(events); i++ {
					if events[i] == nil {
						continue
					}
					move = append(move, events[i])
				}
				player.Update(move...)
			}(tmp)
			_ = t
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {

			case *sdl.QuitEvent:
				return
			}

			state := keyboard.NewState(sdl.GetKeyboardState())
			if state.KeysPressed(sdl.K_w) && state.KeysPressed(sdl.K_a) {
				board.events = append(board.events, Move(sdl.Point{
					X: -1,
					Y: -1,
				}))
				//fmt.Println("press ↖")
			} else if state.KeysPressed(sdl.K_w) && state.KeysPressed(sdl.K_d) {
				board.events = append(board.events, Move(sdl.Point{
					X: 1,
					Y: -1,
				}))
				//fmt.Println("press ↗")
			} else if state.KeysPressed(sdl.K_s) && state.KeysPressed(sdl.K_d) {
				board.events = append(board.events, Move(sdl.Point{
					X: 1,
					Y: 1,
				}))
				//fmt.Println("press ↘")
			} else if state.KeysPressed(sdl.K_s) && state.KeysPressed(sdl.K_a) {
				board.events = append(board.events, Move(sdl.Point{
					X: -1,
					Y: 1,
				}))
				//fmt.Println("press ↙")
			} else if state.KeysPressed(sdl.K_w) {
				board.events = append(board.events, Move(sdl.Point{
					Y: -1,
				}))
				//fmt.Println("press ↑")
			} else if state.KeysPressed(sdl.K_s) {
				board.events = append(board.events, Move(sdl.Point{
					Y: 1,
				}))
				//fmt.Println("press ↓")
			} else if state.KeysPressed(sdl.K_d) {
				board.events = append(board.events, Move(sdl.Point{
					X: 1,
				}))
				//fmt.Println("press →")
			} else if state.KeysPressed(sdl.K_a) {
				board.events = append(board.events, Move(sdl.Point{
					X: -1,
				}))
				//fmt.Println("press ←")
			} else if state.KeysPressed(sdl.K_ESCAPE) {
				return
				//fmt.Println("not press")
			}

		}
	}

}

const maxEventLen = 5

type GameBoard struct {
	events []optFn
	Keys   []uint8
}
