package detail

import "testing"

func TestDetail(t *testing.T) {
	email, _ := NewEmail("sample@example.com")
	profileImage, _ := NewProfileImage("")
	userName, _ := NewUserName("kona")

	t.Run("値オブジェクトを引数に渡して新しいユーザー情報を作成できること", func(t *testing.T) {
		_, err := NewDetail(*email, *profileImage, *userName)
		if err != nil {
			t.Errorf("got :%v, want: %v", err, nil)
		}
	})

	t.Run("emailを更新できること", func(t *testing.T) {
		detail, _ := NewDetail(*email, *profileImage, *userName)
		newEmail, _ := NewEmail("change@example.com")

		detail.ChangeEmail(*newEmail)
		if detail.email.Value() != newEmail.Value() {
			t.Errorf("got :%v, want: %v", detail.email.Value(), newEmail.Value())
		}
	})

	t.Run("profileImageを更新できること", func(t *testing.T) {
		detail, _ := NewDetail(*email, *profileImage, *userName)
		newProfileImage, _ := NewProfileImage("newImageUrl")

		detail.ChangeProfileImage(*newProfileImage)
		if detail.profileImage.Value() != newProfileImage.Value() {
			t.Errorf("got :%v, want: %v", detail.profileImage.Value(), newProfileImage.Value())
		}
	})

	t.Run("userNameを更新できること", func(t *testing.T) {
		detail, _ := NewDetail(*email, *profileImage, *userName)
		newUserName, _ := NewUserName("changeName")

		detail.ChangeUserName(*newUserName)
		if detail.userName.Value() != newUserName.Value() {
			t.Errorf("got :%v, want: %v", detail.userName.Value(), newUserName.Value())
		}
	})
}
