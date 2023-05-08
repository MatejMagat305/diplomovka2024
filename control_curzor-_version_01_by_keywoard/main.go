package main

import (
	"control_curzor/moveSignals"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
)

func main() {
	eventsChanels := hook.Start()
	defer hook.End()
	moveSignals.Init(eventsChanels)
	robotgo.SetDelay(0)
	robotgo.SetMouseDelay(0)
	robotgo.SetKeyDelay(0)
	run()
}
