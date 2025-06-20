package mu

import (
	"testing"
)

func TestBoolToInt(t *testing.T) {
	trials := []struct {
		name     string
		input    bool
		expected int
	}{
		{"false", false, 0},
		{"true", true, 1},
	}

	for _, trial := range trials {
		t.Run(trial.name, func(t *testing.T) {
			got := BoolToInt(trial.input)
			if got != trial.expected {
				t.Errorf("Expected %d, but got %d", trial.expected, got)
			}
		})
	}
}

func TestIntToBool(t *testing.T) {
	trials := []struct {
		name     string
		input    int
		expected bool
	}{
		{"negative", -10, true},
		{"negative-one", -1, true},
		{"zero", 0, false},
		{"one", 1, true},
		{"postive", 10, true},
	}

	for _, trial := range trials {
		t.Run(trial.name, func(t *testing.T) {
			got := IntToBool(trial.input)
			if got != trial.expected {
				t.Errorf("Expected %t, but got %t", trial.expected, got)
			}
		})
	}
}
