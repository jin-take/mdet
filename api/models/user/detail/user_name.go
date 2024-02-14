package detail

import (
	"errors"

	"github.com/jin237/mdet/models/shared"
)

const (
	MIN_USER_NAME_LENGTH = 3
	MAX_USER_NAME_LENGTH = 15
)

type UserName struct {
	shared.ValueObject[string]
}

func validateUserName(userName string) error {
	if len(userName) < MIN_USER_NAME_LENGTH {
		return errors.New("ユーザー名は3文字以上である必要があります")
	}
	if len(userName) > MAX_USER_NAME_LENGTH {
		return errors.New("ユーザー名は15文字以下である必要があります")
	}
	return nil
}

func NewUserName(userName string) (*UserName, error) {
	if err := validateUserName(userName); err != nil {
		return nil, err
	}
	return &UserName{shared.NewValueObject[string](userName)}, nil
}
