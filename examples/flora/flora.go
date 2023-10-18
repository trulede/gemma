package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	flora machine.Pin = machine.D2
	ws    ws2812.Device
	leds  = make([]color.RGBA, 1)
	wheel = &Wheel{Brightness: 0x40}
)

func init() {
	flora.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = ws2812.New(flora)
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		for {
			led.Low()
			time.Sleep(time.Millisecond * 200)
			led.High()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for {
		for i := range leds {
			leds[i] = wheel.Next()
		}
		ws.WriteColors(leds)
		time.Sleep(25 * time.Millisecond)
	}
}

type Wheel struct {
	Brightness uint8
	pos        uint8
}

func (w *Wheel) Next() (c color.RGBA) {
	pos := w.pos
	if w.pos < 85 {
		c = color.RGBA{R: 0xFF - pos*3, G: pos * 3, B: 0x0}
	} else if w.pos < 170 {
		pos -= 85
		c = color.RGBA{R: 0x0, G: 0xFF - pos*3, B: pos * 3}
	} else {
		pos -= 170
		c = color.RGBA{R: pos * 3, G: 0x0, B: 0xFF - pos*3}
	}
	// Apply the alpha adjustment to each color (effective brightness control).
	setAlpha := func(c, a uint8) uint8 {
		return uint8(int(c) * int(a) / 255)
	}
	c.R = setAlpha(c.R, w.Brightness)
	c.G = setAlpha(c.G, w.Brightness)
	c.B = setAlpha(c.B, w.Brightness)
	// Next pos in the wheel
	w.pos++
	return
}
