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

var keys = [...]*KeyState{
	&QuitKey,
	&NextTurnKey,
}

func Process() {
	for {
		event := sdl.PollEvent()
		if event == nil {
			break
		}
		switch event := event.(type) {
		case *sdl.QuitEvent:
			QuitKey.Press()
		case *sdl.KeyDownEvent:
			switch event.Keysym.Scancode {
			case sdl.SCANCODE_SPACE:
				NextTurnKey.Press()
			}
		case *sdl.KeyUpEvent:
			switch event.Keysym.Scancode {
			case sdl.SCANCODE_SPACE:
				NextTurnKey.Release()
			}
		}
	}
}

func Update() {
	for _, key := range keys {
		switch *key {
		case KeyPressed:
			*key = KeyHeld
		case KeyReleased:
			*key = KeyEmpty
		}
	}
}

func (k *KeyState) Press() {
	if *k == KeyEmpty {
		*k = KeyPressed
	}
}

func (k *KeyState) Release() {
	if *k == KeyHeld {
		*k = KeyReleased
	}
}
