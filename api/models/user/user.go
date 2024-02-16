package user

import (
	detail "github.com/jin237/mdet/models/user/detail"
)

type User struct {
	UserDetail *detail.Detail
}

func NewUser(email detail.Email, profileImage detail.ProfileImage, userName detail.UserName) (*User, error) {
	userDetail, err := detail.NewDetail(email, profileImage, userName)
	if err != nil {
		return nil, err
	}
	return &User{UserDetail: userDetail}, nil
}

func (u *User) delete() error {
	// 実装が必要になったときに削除処理を行う
	return nil
}

func (u *User) changeEmail(newEmail detail.Email) error {
	return u.UserDetail.ChangeEmail(newEmail)
}

func (u *User) changeProfileImage(newProfileImage detail.ProfileImage) error {
	return u.UserDetail.ChangeProfileImage(newProfileImage)
}

func (u *User) changeUserName(newUserName detail.UserName) error {
	return u.UserDetail.ChangeUserName(newUserName)
}

func (u *User) getEmail() string {
	return u.UserDetail.GetEmail()
}

func (u *User) getProfileImage() string {
	return u.UserDetail.GetProfileImage()
}

func (u *User) getUserName() string {
	return u.UserDetail.GetUserName()
}
