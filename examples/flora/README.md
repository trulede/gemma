# NeoPixel Flora (1 NeoPixel RGB LED) Test Program

Flora LEDs can be daisy chained by connecting the "Out" pad of one Flora to
the "In" pad of the next Flora. Adjust the code accordingly.

> Note: Example code uses pin D2 to control the Flora LED(s). The Jewel 5 V DC pin can be connected to the Gemma M0 Vout pin (or some other 5 volt supply).


## Build

```bash
# Setup GOROOT if necessary.
tinygo info gemma-m0
export GOROOT=...

# Build the UF2 image file.
tinygo build -target gemma-m0 -o flora.uf2 flora.go
```


## Deploy

1. Connect Gemma M0 device to USB and power on.

2. Double press the reset button, red LED comes on, DotStar LED is green (which indicates USB connection is OK).  

   GEMMABOOT drive should be mounted.

3. Drag the UF2 file directly to the mounted drive, boot loader will then restart and run the executable.  

   Red LED should flash, and NeoPixel Flora LEDs should cycle through their full range of colors.
