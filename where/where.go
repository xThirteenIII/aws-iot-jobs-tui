package where

import (
	"aws-iot-jobs-tui/helper"
	"log"
	"os"
	"path/filepath"
)

// mkdir creates a directory named path, along with
// necessary parents, and returns path.
// If path is already a directory, mkdir does nothing and returns path.
func mkdir(path string) string {

	// MkdirAll creates a directory named path,
	// along with any necessary parents, and returns nil,
	// or else returns an error.
	// The permission bits perm (before umask) are used for all
	// directories that MkdirAll creates.
	// If path is already a directory, MkdirAll does nothing
	// and returns nil.
	helper.Must(os.MkdirAll(path, os.ModePerm))
	return path
}

// Config path is just the working directory.
// ./.env
func Config() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldn't get working directory")
	}
	return path
}

func Logs() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldn't get working directory")
	}

	path := filepath.Join(workingDir, "logs")
	return mkdir(path)
}
