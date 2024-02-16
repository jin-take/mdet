package detail

type Detail struct {
	email        Email
	profileImage ProfileImage
	userName     UserName
}

func NewDetail(email Email, profileImage ProfileImage, userName UserName) (*Detail, error) {
	return &Detail{
		email:        email,
		profileImage: profileImage,
		userName:     userName,
	}, nil
}

func (d *Detail) ChangeEmail(newEmail Email) error {
	d.email = newEmail
	return nil
}

func (d *Detail) ChangeProfileImage(newProfileImage ProfileImage) error {
	d.profileImage = newProfileImage
	return nil
}

func (d *Detail) ChangeUserName(newUserName UserName) error {
	d.userName = newUserName
	return nil
}

func (d *Detail) GetEmail() string {
	return d.email.Value()
}

func (d *Detail) GetProfileImage() string {
	return d.profileImage.Value()
}

func (d *Detail) GetUserName() string {
	return d.userName.Value()
}
