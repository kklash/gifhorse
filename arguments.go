package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage(reason string) {
	fmt.Printf("ERROR: %s\n", reason)
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	flag.StringVar(&inFile, "in", "", "input video file name (REQUIRED)")
	flag.StringVar(&outFile, "out", "out.gif", "output .gif file name (must end with .gif)")
	flag.IntVar(&width, "width", 0, "output image frame width (in pixels). incompatible with -resize")
	flag.IntVar(&height, "height", 0, "output image frame height (in pixels). incompatible with -resize")
	flag.Float64Var(&resize, "resize", 0, "image resizing multiplier. for example, '-resize 0.5' means the output frame size is half of the input size.")
	flag.IntVar(&frameRate, "framerate", 15, "output gif framerate")
	flag.Float64Var(&offset, "offset", 0, "offset time (in seconds) from start of video to convert from")
	flag.Float64Var(&duration, "duration", 0, "duration of the video (in seconds) from -offset to convert from")
	flag.BoolVar(&force, "force", false, "automatically overwrite output files")
	flag.BoolVar(&debug, "debug", false, "debug mode, print ffmpeg output")

	flag.Parse()

	if len(inFile) == 0 {
		usage("must provide -in file")
	} else if !strings.HasSuffix(outFile, ".gif") {
		usage("-out file must have .gif extension")
	} else if resize > 0 && (width > 0 || height > 0) {
		usage("cannot mix -resize with -height and -width")
	} else if resize < 0 {
		usage("-resize cannot be negative")
	} else if width < 0 {
		usage("-width cannot be negative")
	} else if height < 0 {
		usage("-height cannot be negative")
	} else if frameRate <= 0 {
		usage("-framerate must be > 0")
	}

	if width == 0 {
		width = -1 // Cue to ffmpeg to scale with height if height is provided
	}
	if height == 0 {
		height = -1 // Cue to ffmpeg to scale with width if width is provided
	}
}
