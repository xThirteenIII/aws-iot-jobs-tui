package tui

import (
	"aws-iot-jobs-tui/utils"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type statefulBubble struct {
	state	state
	stateHistory utils.Stack[state]
	loading bool

	// Components
	thingList list.Model
	jobList list.Model

	input textinput.Model

	selectedThing string
	selectedJob string

	width, height int
}
