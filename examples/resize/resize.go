// goncurses - ncurses library for Go.

/* This example demonstrates the ability to resize. Only one of detecting SIGWINCH or KEY_RESIZE
 * is strictly needed, but depending on the options ncurses was built with, one or the other may
 * work better. */
package main

import (
	"os"
	"os/signal"
	"syscall"

	gc "github.com/mintalk/goncurses"
)

var stdscr *gc.Window
var sigWinChCount, keyResizeCount int

func main() {
	sigWinChCount = 0
	keyResizeCount = 0
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGWINCH)

	// Errors should not be ignored in production code
	stdscr, _ = gc.Init()
	defer gc.End()
	stdscr.Timeout(0)
	redrawDisplay()

	for {
		select {
		case <-sigs:
			sigWinChCount++
			resize()
		default:
			c := stdscr.GetChar()
			switch c {
			case gc.KEY_RESIZE:
				keyResizeCount++
				//resize()
			case 'q':
				return
			}
		}
	}
}

func redrawDisplay() {
	// Errors should not be ignored in production code
	stdscr.Clear()
	row, col := stdscr.MaxYX()
	tRow, tCol, _ := osTermSize()
	stdscr.MovePrintf(1, 1, "     MaxYX shows %d rows and %d columns", row, col)
	stdscr.MovePrintf(2, 1, "osTermSize shows %d rows and %d columns", tRow, tCol)
	stdscr.MovePrintf(3, 1, "  SIGWINCH has been sent %d times", sigWinChCount)
	stdscr.MovePrintf(4, 1, "KEY_RESIZE has been sent %d times", keyResizeCount)
	stdscr.MovePrint(6, 1, "Press 'q' to quit")
	stdscr.Box(0, 0)
	stdscr.Refresh()
}

func resize() {
	// Errors should not be ignored in production code
	tRow, tCol, _ := osTermSize()
	gc.ResizeTerm(tRow, tCol)

	redrawDisplay()
}
