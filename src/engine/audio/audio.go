package audio

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var context = audio.NewContext(44100)

type Player = audio.Player

func LoadWAV(data []byte, loop bool) (*Player, error) {
	stream, err := wav.DecodeWithSampleRate(context.SampleRate(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var player *audio.Player
	if loop {
		player, err = context.NewPlayer(audio.NewInfiniteLoop(stream, stream.Length()))
	} else {
		player, err = context.NewPlayer(stream)
	}
	if err != nil {
		return nil, fmt.Errorf("could not create audio player: %v", err)
	}

	return player, err
}

func LoadMP3(data []byte, loop bool) (*Player, error) {
	stream, err := mp3.DecodeWithSampleRate(context.SampleRate(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var player *audio.Player
	if loop {
		player, err = context.NewPlayer(audio.NewInfiniteLoop(stream, stream.Length()))
	} else {
		player, err = context.NewPlayer(stream)
	}
	if err != nil {
		return nil, fmt.Errorf("could not create audio player: %v", err)
	}

	return player, err
}

func Play(player *Player) {
	if err := player.Rewind(); err != nil {
		panic(fmt.Errorf("could not rewind audio player: %v", err))
	}
	player.Play()
}

func PausePlayer(player *Player) {
	player.Pause()
}

func ResumePlayer(player *Player) {
	player.Play()
}

func SetVolume(player *Player, volume float64) {
	player.SetVolume(volume)
}
