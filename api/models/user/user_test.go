package user

import (
	"testing"

	detail "github.com/jin237/mdet/models/user/detail"
)

func TestUser(t *testing.T) {
	email, _ := detail.NewEmail("sample@example.com")
	profileImage, _ := detail.NewProfileImage("")
	userName, _ := detail.NewUserName("kona")

	t.Run("値オブジェクトを引数に渡して新しいユーザー情報を作成できること", func(t *testing.T) {
		_, err := NewUser(*email, *profileImage, *userName)
		if err != nil {
			t.Errorf("got :%v, want: %v", err, nil)
		}
	})

	t.Run("emailを更新できること", func(t *testing.T) {
		user, _ := NewUser(*email, *profileImage, *userName)
		newEmail, _ := detail.NewEmail("new@example.com")

		user.changeEmail(*newEmail)
		if user.getEmail() != newEmail.Value() {
			t.Errorf("got :%v, want: %v", user.getEmail(), newEmail.Value())
		}
	})

	t.Run("profileImageを更新できること", func(t *testing.T) {
		user, _ := NewUser(*email, *profileImage, *userName)
		newProfileImage, _ := detail.NewProfileImage("newImageUrl")

		user.changeProfileImage(*newProfileImage)
		if user.getProfileImage() != newProfileImage.Value() {
			t.Errorf("got :%v, want: %v", user.getProfileImage(), newProfileImage.Value())
		}
	})

	t.Run("userNameを更新できること", func(t *testing.T) {
		user, _ := NewUser(*email, *profileImage, *userName)
		newUserName, _ := detail.NewUserName("changeName")

		user.changeUserName(*newUserName)
		if user.getUserName() != newUserName.Value() {
			t.Errorf("got :%v, want: %v", user.getUserName(), newUserName.Value())
		}
	})
}
