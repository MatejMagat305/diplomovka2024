package main

import (
	"control_curzor/moveSignals"
	"fmt"
	"github.com/go-vgo/robotgo"
)

func run() {
	active := true
	step := 10
	for {
		i := moveSignals.GetKey()
		fmt.Println(i)
		if !active {
			if i.Rawcode == moveSignals.ACTIVATE() {
				active = true
			}
			continue
		}
		x, y := robotgo.GetMousePos()
		fmt.Println(x, y)
		switch i.Rawcode {
		case moveSignals.UP():
			robotgo.MoveSmooth(x, y-step)
		case moveSignals.DOWN():
			robotgo.MoveSmooth(x, y+step)
		case moveSignals.LEFT():
			robotgo.MoveSmooth(x-step, y)
		case moveSignals.RIGHT():
			robotgo.MoveSmooth(x+step, y)
		case moveSignals.LCLICK():
			robotgo.Click()
		case moveSignals.RCLICK():
			robotgo.Click("right")
		case moveSignals.DEACTIVATE():
			active = false
		case moveSignals.PLUS():
			step++
		case moveSignals.MINUS():
			step--
		default:
			fmt.Println(i)
		}
	}
}
