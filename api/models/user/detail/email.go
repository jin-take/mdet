package detail

import (
	"errors"
	"strings"

	shared "github.com/jin237/mdet/models/shared"
)

type Email struct {
	shared.ValueObject[string]
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("メールアドレスは@を含む必要があります")
	}
	return nil
}

func NewEmail(email string) (*Email, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	return &Email{shared.NewValueObject[string](email)}, nil
}
