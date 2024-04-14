// ncurses_example.go
package goncurses_test

import (
	"testing"

	"github.com/mintalk/goncurses"
)

func TestInit(t *testing.T) {
	_, err := goncurses.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer goncurses.End()
}
