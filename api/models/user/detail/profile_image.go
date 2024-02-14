package detail

import "github.com/jin237/mdet/models/shared"

type ProfileImage struct {
	shared.ValueObject[string]
}

func validateProfileImage(profileImage string) error {
	// 実装が必要になったときにバリデーションを行う
	return nil
}

func NewProfileImage(profileImage string) (*ProfileImage, error) {
	if err := validateProfileImage(profileImage); err != nil {
		return nil, err
	}
	return &ProfileImage{shared.NewValueObject[string](profileImage)}, nil
}
