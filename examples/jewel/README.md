# NeoPixel Jewel (7 NeoPixel RGBW LEDs) Test Program

> Note: Example code uses pin D2 to control the Jewel LEDs. The Jewel 5 V DC pin can be connected to the Gemma M0 Vout pin (or some other 5 volt supply).

## Data Sheets

 * [WS2812](https://cdn-shop.adafruit.com/datasheets/WS2812.pdf) [^1]

[^1]: The RGBW LEDs require an additional byte of data (for the white LED) to be transmitted. The ws2812 driver only supports 3 LEDS, however this example adds support for the 4th white LED.

## Build

```bash
# Setup GOROOT if necessary.
tinygo info gemma-m0
export GOROOT=...

# Build the UF2 image file.
tinygo build -target gemma-m0 -o jewel.uf2 jewel.go
```


## Deploy

1. Connect Gemma M0 device to USB and power on.

2. Double press the reset button, red LED comes on, DotStar LED is green (which indicates USB connection is OK).  

   GEMMABOOT drive should be mounted.

3. Drag the UF2 file directly to the mounted drive, boot loader will then restart and run the executable.  

   Red LED should flash, and NeoPixel Jewel 7 LEDs should cycle through their full range of colors.

