package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/apa102"
)

var (
	apa   *apa102.Device
	leds  = make([]color.RGBA, 1)
	wheel = &Wheel{Brightness: 0x10}
)

func init() {
	apa = apa102.NewSoftwareSPI(machine.SPI1_SCK_PIN, machine.SPI1_SDO_PIN, 1)
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go func() {
		for {
			led.Low()
			time.Sleep(time.Millisecond * 1500)

			led.High()
			time.Sleep(time.Millisecond * 500)

			led.Low()
			time.Sleep(time.Millisecond * 500)

			led.High()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for {
		leds[0] = wheel.Next()
		apa.WriteColors(leds)
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
		c = color.RGBA{R: 0xFF - pos*3, G: pos * 3, B: 0x0, A: w.Brightness}
	} else if w.pos < 170 {
		pos -= 85
		c = color.RGBA{R: 0x0, G: 0xFF - pos*3, B: pos * 3, A: w.Brightness}
	} else {
		pos -= 170
		c = color.RGBA{R: pos * 3, G: 0x0, B: 0xFF - pos*3, A: w.Brightness}
	}
	w.pos++
	return
}
