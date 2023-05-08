package main

import (
	"github.com/go-vgo/robotgo"
)

func runServer(ch chan Comand) {
	step := 5
	robotgo.SetMouseDelay(0)

	for {
		d := <-ch
		x, y := robotgo.GetMousePos()
		switch d {
		case UP:
			robotgo.Move(x, y-step)
		case DOWN:
			robotgo.Move(x, y+step)
		case LEFT:
			robotgo.Move(x-step, y)
		case RIGHT:
			robotgo.Move(x+step, y)
		case LEFTCLICK:
			robotgo.Click()
		case RIGHTCLICK:
			robotgo.Click("right")
		case PLUS:
			step++
		case MINUS:
			step--

		}
	}

}
