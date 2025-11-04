package tui

import "github.com/charmbracelet/bubbles/key"

type statefulKeymap struct {
	state state

	quit, forceQuit,
	selectOne,
	confirm,
	openFolder,
	back,
	filter,
	up, down, left, right,
	top, bottom,
	showHelp key.Binding
}

func (k *statefulKeymap) setState(newState state) {
	k.state = newState
}

func newStatefulKeymap() *statefulKeymap {
	k := key.NewBinding
	keys := key.WithKeys
	help := key.WithHelp

	return &statefulKeymap {
		quit: k(
			keys("q"),
			help("q", "quit"),
		),
		forceQuit: k(
			keys("ctrl+c", "ctrl+d"),
			help("ctrl+c", "quit"),
		),
		selectOne: k(
			keys(" "),
			help("space", "select item"),
		),
		confirm: k(
			keys("enter"),
			help("enter", "confirm"),
		),
		openFolder: k(
			keys("o"),
			help("o", "open folder"),
		),
		back: k(
			keys("b"),
			help("b", "back"),
		),
		filter: k(
			keys("/"),
			help("/", "filter"),
		),
		up: k(
			keys("up", "k"),
			help("↑", "up"),
		),
		down: k(
			keys("down", "j"),
			help("↑", "down"),
		),
		left: k(
			keys("left", "h"),
			help("←", "left"),
		),
		right: k(
			keys("right", "l"),
			help("→", "right"),
		),
		top: k(
			keys("g"),
			help("g", "top"),
		),
		bottom: k(
			keys("G"),
			help("G", "bottom"),
		),
		showHelp: k(
			keys("?", "h"),
			help("?", "help"),
		),
	}
}
