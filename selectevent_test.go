package main

import (
	"testing"
)

func TestSelectEvent(t *testing.T) {
	var tests = []struct {
		description string
		state State
		eventType string
		involvedObject string
		reason string
		expected bool
	}{
		{"match_type", getState("Warning", "", ""), "Warning", "", "", true},
		{"match_involved_object", getState("", "Pod", ""), "", "Pod", "", true},
		{"match_reason", getState("", "", "Failed"), "", "", "Failed", true},
		{"mismatch_type", getState("Warning", "", ""), "Normal", "", "", false},
		{"mismatch_involved_object", getState("Warning", "ConfigMap", ""), "Warning", "", "", false},
		{"mismatch_reason", getState("Warning", "", "Completed"), "Warning", "", "Failed,Evicted,FailedMount", false},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := selectEvent(test.state, test.eventType, test.involvedObject, test.reason)
			if result != test.expected {
				t.Errorf("Expected result %t for parameters state=%v, eventType=%s, involvedObject=%s, reason=%s", result, test.state, test.eventType, test.involvedObject, test.reason)
			}
		})
	}
}

func TestSelectGeneric(t *testing.T) {
        var tests = []struct {
                description string
		items []string
		item string
                expected bool
        }{
		{"empty_list", []string{}, "Warning", true},
		{"match", []string{"Normal", "Warning"}, "Warning", true},
		{"mismatch", []string{"Failed", "Evicted"}, "FailedScheduling", false},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := selectGeneric(test.items, test.item)
			if result != test.expected {
				t.Errorf("Expected result %t for parameters items=%v, item=%s", result, test.items, test.item)
			}
		})
	}
}
