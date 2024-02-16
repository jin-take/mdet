package detail

import (
	"errors"
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{"正常に登録できること", "sample@example.com", nil},
		{"バリデーションエラーになること", "sample.com", errors.New("メールアドレスは@を含む必要があります")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.input)

			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("got: %v, want: %v", err, tt.want)
				}
			} else {
				if email.Value() != tt.input {
					t.Errorf("got: %v, want: %v", email.Value(), tt.input)
				}
			}
		})
	}
}
