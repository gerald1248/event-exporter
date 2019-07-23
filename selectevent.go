package main

func selectEvent(state State, eventType, involvedObject, reason string) bool {
	return selectGeneric(state.eventTypes, eventType) &&
		selectGeneric(state.involvedObjects, involvedObject) &&
		selectGeneric(state.reasons, reason)
}

func selectGeneric(items []string, givenItem string) bool {
	if len(items) == 0 || items == nil {
		return true
	}
	selected := false
	for _, item := range items {
		if item == givenItem {
			selected = true
			break
		}
	}
	return selected
}
