package main

import (
	"machine"
	"machine/usb/midi"
	"time"

	"tinygo.org/x/drivers/makeybutton"
)

const (
	keyOfMusic = midi.C4
)

var (
	buttonC machine.Pin = machine.D12

	key         *makeybutton.Button
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	m := midi.New()

	key = makeybutton.NewButton(buttonC)
	key.Configure()

	for {
		switch key.Get() {
		case makeybutton.Pressed:
			m.NoteOn(0, midichannel, keyOfMusic, 0x40)
		case makeybutton.Released:
			m.NoteOff(0, midichannel, keyOfMusic, 0x40)
		}
		time.Sleep(100 * time.Millisecond)
	}
}