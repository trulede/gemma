# Board LED Test Program

## Build

```bash
# Setup GOROOT if necessary.
tinygo info gemma-m0
export GOROOT=...

# Build the UF2 image file.
tinygo build -target gemma-m0 -o boardleds.uf2 boardleds.go
```

## Deploy

1. Connect Gemma M0 device to USB and power on.

2. Double press the reset button, red LED comes on, DotStar LED is green (which indicates USB connection is OK).  

   GEMMABOOT drive should be mounted.

3. Drag the UF2 file directly to the mounted drive, boot loader will then restart and run the executable.  

   Red LED should flash, and DotStar LED should cycle through its full range of colors.

