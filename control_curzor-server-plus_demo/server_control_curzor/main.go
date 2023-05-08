package main

type Comand int

const (
	UP = Comand(iota)
	DOWN
	RIGHT
	LEFT
	RIGHTCLICK
	LEFTCLICK
	PLUS
	MINUS
)

func main() {
	ch := make(chan Comand)
	go runServer(ch)
	runMove(ch)
}
