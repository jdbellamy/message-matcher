package matcher

//go:generate mockery -name=Message

import (
	"errors"

	"github.com/tidwall/gjson"
)

type Message interface {
	Value() string
}

type message struct {
	raw string
}

func NewMessage(json string) (Message, error) {
	if !gjson.Valid(json) {
		return nil, errors.New("invalid json")
	}
	return &message{json}, nil
}

// TODO - return a more useful type here... validation has already been performed
func (msg *message) Value() string {
	return msg.raw
}
