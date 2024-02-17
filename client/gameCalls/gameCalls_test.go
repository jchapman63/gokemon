package gameCalls

import "testing"

// Test that a list of pokemon names come back from the server
func TestAvailableMonsters(t *testing.T) {
	got, err := AvailableMonsters()
	if err != nil {
		t.Errorf("Did not want an error")
	}
	t.Logf("Result: %q", got)
}
