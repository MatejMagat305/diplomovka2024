package moveSignals

type State struct {
	UP, DOWN, LEFT, RIGHT, LCLICK, RCLICK, DEACTIVATE, ACTIVATE, QUIT, PLUS, MINUS uint16
}

var (
	state State
)

func GetDefaultState() State {
	return State{
		UP:     65362,
		DOWN:   65364,
		LEFT:   65361,
		RIGHT:  65363,
		LCLICK: 65293,
		RCLICK: 65535,
	}
}

func UP() uint16 {
	return state.UP
}

func DOWN() uint16 {
	return state.DOWN
}

func LEFT() uint16 {
	return state.LEFT
}

func RIGHT() uint16 {
	return state.RIGHT
}

func LCLICK() uint16 {
	return state.LCLICK
}

func RCLICK() uint16 {
	return state.RCLICK
}

func DEACTIVATE() uint16 {
	return state.DEACTIVATE
}

func ACTIVATE() uint16 {
	return state.ACTIVATE
}

func QUIT() uint16 {
	return state.QUIT
}

func PLUS() uint16 {
	return state.PLUS
}

func MINUS() uint16 {
	return state.MINUS
}
