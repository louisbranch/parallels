package input

import "github.com/veandco/go-sdl2/sdl"

type KeyState int

const (
	KeyEmpty KeyState = iota //FIXME better name
	KeyPressed
	KeyHeld
	KeyReleased
)

var QuitKey KeyState
var EscKey KeyState
var NextTurnKey KeyState
var UpKey KeyState
var DownKey KeyState
var LeftKey KeyState
var RightKey KeyState
var ZoomInKey KeyState
var ZoomOutKey KeyState

var keymappings = map[sdl.Keycode]*KeyState{
	sdl.K_ESCAPE: &EscKey,
	sdl.K_SPACE:  &NextTurnKey,
	sdl.K_UP:     &UpKey,
	sdl.K_DOWN:   &DownKey,
	sdl.K_LEFT:   &LeftKey,
	sdl.K_RIGHT:  &RightKey,
	sdl.K_EQUALS: &ZoomInKey,
	sdl.K_MINUS:  &ZoomOutKey,
}

func Process() {
	for {
		event := sdl.PollEvent()
		if event == nil {
			break
		}
		switch event := event.(type) {
		case *sdl.QuitEvent:
			QuitKey = KeyPressed
		case *sdl.KeyDownEvent:
			key, ok := keymappings[event.Keysym.Sym]
			if ok && *key == KeyEmpty {
				*key = KeyPressed
			}
		case *sdl.KeyUpEvent:
			key, ok := keymappings[event.Keysym.Sym]
			if ok && *key == KeyHeld {
				*key = KeyReleased
			}
		}
	}
}

func Update() {
	for _, key := range keymappings {
		switch *key {
		case KeyPressed:
			*key = KeyHeld
		case KeyReleased:
			*key = KeyEmpty
		}
	}
}
