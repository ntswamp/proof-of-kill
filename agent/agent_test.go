package agent

import "testing"

func TestNew(t *testing.T) {
	agent := New("PoK God", 2, 4)
	t.Log(agent)
}
