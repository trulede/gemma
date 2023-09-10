# NeoPixel Jewel (7 NeoPixel LEDs) Test Program

> Note: Example code uses pin D2 to control the Jewel LEDs. The Jewel 5 V DC pin can be connected to the Gemma M0 Vout pin (or some other 5 volt supply).


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

