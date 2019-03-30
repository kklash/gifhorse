package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	inFile    string
	outFile   string
	width     int
	height    int
	frameRate int
	offset    float64
	duration  float64
	resize    float64
	debug     bool
	force     bool
)

func main() {
	binary, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("ERROR: ffmpeg is not installed")
		fmt.Println("See https://github.com/kklash/gifhorse/blob/master/README.md#Install")
		os.Exit(2)
	}

	if !force && fileExists(outFile) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("output file '%s' exists, do you want to overwrite? [y/n]: ", outFile)
		resp, _ := reader.ReadString('\n')
		resp = strings.TrimSpace(strings.ToLower(resp))
		if resp != "y" && resp != "yes" {
			os.Exit(4)
		}
	}

	argv := []string{
		"-y", // overwrite files
		"-i",
		inFile,
		"-r",
		fmt.Sprintf("%d", frameRate),
	}

	if offset > 0 {
		argv = append(argv, "-ss", fmt.Sprintf("%f", offset))
	}

	if duration > 0 {
		argv = append(argv, "-t", fmt.Sprintf("%f", duration))
	}

	if height > 0 || width > 0 {
		argv = append(argv, "-vf", fmt.Sprintf("scale=%d:%d", width, height))
	} else if resize > 0 {
		argv = append(argv, "-vf", fmt.Sprintf("scale=iw*%f:ih*%f", resize, resize))
	}

	argv = append(argv, outFile)

	cmd := exec.Command(binary, argv...)
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error while running ffmpeg command")
		fmt.Println("use -debug option to get ffmpeg error output")
		fmt.Println(err)
		os.Exit(3)
	}

	size, err := fileSize(outFile)
	if err == nil {
		fmt.Printf("Converted to gif, size %.2f MB\n", size)
	} else {
		fmt.Println("Error getting output file size:")
		fmt.Println(err)
	}

	fmt.Printf("%s -> %s\n", inFile, outFile)
}
