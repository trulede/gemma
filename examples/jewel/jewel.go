// Example for NeoPixel Jewel with RGBW LEDs.
package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

// Wrap the ws2812 device and add handling for the white LED.
type ws2812W struct {
	ws2812.Device
}

func new(pin machine.Pin) ws2812W {
	return ws2812W{ws2812.Device{pin}}
}

type RGBAW struct {
	R, G, B, A, W uint8
}

func (d ws2812W) writeColorsW(buf []RGBAW) error {
	setAlpha := func(c, a uint8) uint8 {
		return uint8(int(c) * int(a) / 255)
	}
	for _, color := range buf {
		d.WriteByte(setAlpha(color.G, color.A))
		d.WriteByte(setAlpha(color.R, color.A))
		d.WriteByte(setAlpha(color.B, color.A))
		d.WriteByte(color.W)
	}
	return nil
}

// Setup the example for NewoPixel Jewel with RGBW LEDs.
var (
	jewel machine.Pin = machine.D2
	ws    ws2812W
	leds  = make([]RGBAW, 7)
	wheel = &Wheel{Brightness: 0x40}
)

func init() {
	jewel.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = new(jewel)
}

func main() {
	// Flash the board LED.
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

	// Loop through the demonstration.
	for {
		wheel.Loop(leds)
		Fade(0, 100, leds)
		Fade(100, 0, leds)
	}
}

// LED Fade example.
func Fade(start uint8, end uint8, leds []RGBAW) {
	steps, inc := func(start, end uint8) (uint8, bool) {
		if end > start {
			return end - start, true
		}
		return start - end, false
	}(start, end)
	wValue := start
	var step uint8
	for step = 0; step < steps; step++ {
		for i := range leds {
			leds[i] = RGBAW{W: wValue}
		}
		ws.writeColorsW(leds)
		time.Sleep(25 * time.Millisecond)
		if inc {
			wValue += 1
		} else {
			wValue -= 1
		}
	}
}

// LED color wheel example.
type Wheel struct {
	Brightness uint8
	pos        uint8
}

func (w *Wheel) Next() (c RGBAW) {
	pos := w.pos
	if w.pos < 85 {
		c = RGBAW{R: 0xFF - pos*3, G: pos * 3, B: 0x0, A: w.Brightness}
	} else if w.pos < 170 {
		pos -= 85
		c = RGBAW{R: 0x0, G: 0xFF - pos*3, B: pos * 3, A: w.Brightness}
	} else {
		pos -= 170
		c = RGBAW{R: pos * 3, G: 0x0, B: 0xFF - pos*3, A: w.Brightness}
	}
	w.pos++
	return
}

func (w *Wheel) Loop(leds []RGBAW) {
	w.pos = 0
	for w.pos < 255 {
		for i := range leds {
			leds[i] = wheel.Next()
		}
		ws.writeColorsW(leds)
		time.Sleep(25 * time.Millisecond)
	}
}
