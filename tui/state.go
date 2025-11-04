package tui

type state int

// The standard way of introducing enumerations in Go is 
// to declare a custom type and a const group with iota.
// Since variables have a 0 default value, you should usually start your enums on a non-zero value.

const (
	selectThingState state = iota + 1
	inputThingState
	selectJobState
	inputJobState
	errorState
	loadingState
	searchState
	confirmState
	queuedJobState
	inProgressJobState
	completedJobState
)
