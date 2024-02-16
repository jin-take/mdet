package post

import (
	"errors"

	"github.com/jin237/mdet/models/shared"
)

const (
	MaxBodyLength = 10000
)

type Body struct {
	shared.ValueObject[string]
}

func validateBody(body string) error {
	if len(body) > MaxBodyLength {
		return errors.New("本文は5,000文字以下である必要があります")
	}
	return nil
}

func NewBody(body string) (*Body, error) {
	if err := validateBody(body); err != nil {
		return nil, err
	}
	return &Body{shared.NewValueObject[string](body)}, nil
}
