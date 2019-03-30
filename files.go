package main

import (
	"os"
)

// Returns file size in MB
func fileSize(file string) (float64, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	bytes := float64(stat.Size())
	return bytes / 1024 / 1024, nil
}

// Returns whether file exists
func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
