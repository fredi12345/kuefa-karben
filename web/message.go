package web

import "encoding/gob"

func init() {
	gob.Register(&message{})
}

type message struct {
	Type string
	Text string
}

const (
	TypeHint  = "message-hint"
	TypeError = "message-error"
)
