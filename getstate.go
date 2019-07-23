package main

import (
	"strings"
)

func getState(eventTypes, involvedObjects, reasons string) State {
	var state State
	state.eventTypes = splitParameter(eventTypes)
	state.involvedObjects = splitParameter(involvedObjects)
	state.reasons = splitParameter(reasons)

	return state
}

func splitParameter(s string) []string {
	var items []string
	if len(s) > 0 {
		items = strings.Split(s, ",")
	}
	return items
}
