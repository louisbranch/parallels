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
var NextTurnKey KeyState

var keymappings = map[sdl.Scancode]*KeyState{
	sdl.SCANCODE_SPACE: &NextTurnKey,
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
			key, ok := keymappings[event.Keysym.Scancode]
			if ok && *key == KeyEmpty {
				*key = KeyPressed
			}
		case *sdl.KeyUpEvent:
			key, ok := keymappings[event.Keysym.Scancode]
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
