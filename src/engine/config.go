package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/ini.v1"
)

var (
	KeyBindings map[InputKey]ebiten.Key
	PlayerSpeed float64
)

var KeyMap = map[string]ebiten.Key{
	"A":      ebiten.KeyA,
	"D":      ebiten.KeyD,
	"W":      ebiten.KeyW,
	"S":      ebiten.KeyS,
	"Space":  ebiten.KeySpace,
	"Escape": ebiten.KeyEscape,
}

func LoadConfig(file string) error {
	cfg, err := ini.Load(file)
	if err != nil {
		return fmt.Errorf("could not load config file: %v", err)
	}

	KeyBindings = map[InputKey]ebiten.Key{
		InputKeyLeft:      KeyMap[(cfg.Section("controls").Key("left").String())],
		InputKeyRight:     KeyMap[(cfg.Section("controls").Key("right").String())],
		InputKeyUp:        KeyMap[(cfg.Section("controls").Key("up").String())],
		InputKeyDown:      KeyMap[(cfg.Section("controls").Key("down").String())],
		InputkeyPlaceBomb: KeyMap[(cfg.Section("controls").Key("place_bomb").String())],
		InputKeyEscape:    KeyMap[(cfg.Section("controls").Key("escape").String())],
	}

	PlayerSpeed, err = cfg.Section("player").Key("speed").Float64()
	if err != nil {
		return fmt.Errorf("could not load player speed: %v", err)
	}

	return nil
}
