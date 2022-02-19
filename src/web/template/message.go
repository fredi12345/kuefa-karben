package template

import "encoding/gob"

func init() {
	gob.Register(&Message{})
}

type Message struct {
	Type string
	Text string
}

const (
	TypeHint  = "message-hint"
	TypeError = "message-error"
)
