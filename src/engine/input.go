package engine

import "github.com/hajimehoshi/ebiten/v2"

type InputKey int

const (
	InputKeyLeft InputKey = iota
	InputKeyRight
	InputKeyUp
	InputKeyDown
	InputkeyPlaceBomb
	InputKeyEscape
)

type KeyState int

const (
	KeyStateUnpressed KeyState = iota
	KeyStatePressed
	KeyStateHeld
)

type InputSate struct {
	Left      KeyState
	Right     KeyState
	Up        KeyState
	Down      KeyState
	PlaceBomb KeyState
	Escape    KeyState
}

func NewInputState() *InputSate {
	return &InputSate{}
}

func (is *InputSate) Update() {
	setKeyState(ebiten.IsKeyPressed(KeyBindings[InputKeyLeft]), &is.Left)
	setKeyState(ebiten.IsKeyPressed(KeyBindings[InputKeyRight]), &is.Right)
	setKeyState(ebiten.IsKeyPressed(KeyBindings[InputKeyUp]), &is.Up)
	setKeyState(ebiten.IsKeyPressed(KeyBindings[InputKeyDown]), &is.Down)
	setKeyState(ebiten.IsKeyPressed(KeyBindings[InputkeyPlaceBomb]), &is.PlaceBomb)
	setKeyState(ebiten.IsKeyPressed(KeyBindings[InputKeyEscape]), &is.Escape)

}

func setKeyState(isPressed bool, state *KeyState) {
	if isPressed {
		if *state > 0 {
			*state = KeyStateHeld
		} else {
			*state = KeyStatePressed
		}
	} else {
		*state = KeyStateUnpressed
	}
}
