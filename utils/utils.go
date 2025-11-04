package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/term"
)

func TerminalSize() (width, height int, err error) {
	return term.GetSize(int(os.Stdout.Fd()))
}

// FileStem returns path without the extension.
func FileStem(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}

func ClearScreen() {
	run := func(name string, args ...string) error {
		command := exec.Command(name, args...)
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		return command.Run()
	}

	// Change cmd based on running OS.
	switch runtime.GOOS {
	case "linux", "darwin":

		// The tput utility uses the terminfo database to make the 
		// values of terminal-dependent capabilities and information available to the shell.
		err := run("tput", "clear")
		if err != nil {
			_ = run("clear")
		}
	case "windows":
		_ = run("cls")
	}
}

func Max[T constraints.Ordered](items ...T) (max T) {
	for _, item := range items {
		if item[T] > max {
			max = item
		}
	}
	return
}

func Min[T constraints.Ordered](items ...T) (min T) {
	min = items[0]
	for _, item := range items {
		if item[T] < min {
			min = item
		}
	}
	return
}
