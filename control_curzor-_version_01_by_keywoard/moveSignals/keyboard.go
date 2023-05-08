package moveSignals

import (
	"encoding/json"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
)

var (
	eventsChanels chan hook.Event
)

const (
	KEYBOARD = true
	nameFile = "setting.json"
)

func Init(c chan hook.Event) {
	eventsChanels = c
	load := loadFromFileIntoGlobalState()
	if !load || notAllSet() {
		manually := robotgo.ShowAlert("setting",
			"previous setting is not found, you can choose whether you set manually or load default",
			"load default", "set manually")
		if manually {
			state = GetDefaultState()
		} else {
			setting()
		}
	}
}

func notAllSet() bool {
	all := []uint16{state.UP, state.DOWN, state.LEFT, state.RIGHT, state.LCLICK, state.RCLICK, state.DEACTIVATE, state.ACTIVATE, state.QUIT, state.PLUS, state.MINUS}
	for i := 0; i < len(all); i++ {
		if all[i] == 0 {
			return true
		}
	}
	return false
}

func GetKey() hook.Event {
	if KEYBOARD {
		return <-eventsChanels
	} else {
		return bitalino()
	}
}

func bitalino() hook.Event {
	// todo next step
	panic("unimplemented")
}

func setting() {
	ch := make(chan struct{})
	title := []string{"up", "down", "left", "right", "left click", "right click", "deactivate", "activate", "quit", "plus", "minus"}
	addr := []*uint16{&state.UP, &state.DOWN, &state.LEFT, &state.RIGHT, &state.LCLICK, &state.RCLICK, &state.DEACTIVATE, &state.ACTIVATE, &state.QUIT, &state.PLUS, &state.MINUS}
	lastPressed := uint16(0)
	for i := 0; i < len(title); i++ {
		go func() {
		f:
			for {
				if lastPressed != 0 {
					select {
					case <-ch:
						*(addr[i]) = lastPressed
						break f
					default:
					}
				}
				p := <-eventsChanels
				if p.Rawcode == 0 {
					continue
				}
				lastPressed = p.Rawcode
			}
		}()

		robotgo.ShowAlert("setting", "now press key, which will mean "+title[i]+" (it took last pressed before ok)",
			"ok", "")
		ch <- struct{}{}
	}
	save := robotgo.ShowAlert("saving", "do you want to save it",
		"yes", "no")
	if save {
		saveState()
	}
}

func saveState() {
	file, _ := json.MarshalIndent(&state, "", " ")
	_ = os.WriteFile(nameFile, file, 0644)
}

func loadFromFileIntoGlobalState() bool {
	file, err := os.ReadFile(nameFile)
	if err != nil {
		return false
	}
	err = json.Unmarshal(file, &state)
	return err == nil
}
