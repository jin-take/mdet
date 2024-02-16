package detail

import (
	"errors"
	"strings"
	"testing"
)

func TestUserName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{"正常に登録できること", "kona", nil},
		{"バリデーションエラーになること", "ho", errors.New("ユーザー名は3文字以上である必要があります")},
		{"バリデーションエラーになること", strings.Repeat("hoge", 4), errors.New("ユーザー名は15文字以下である必要があります")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userName, err := NewUserName(tt.input)

			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("got: %v, want: %v", err, tt.want)
				}
			} else {
				if userName.Value() != tt.input {
					t.Errorf("got: %v, want: %v", userName.Value(), tt.input)
				}
			}
		})
	}
}
