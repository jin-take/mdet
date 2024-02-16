package post

import (
	"errors"
	"strings"
	"testing"
)

func TestBody(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{"正常に登録できること", "post body text.", nil},
		{"バリデーションエラーになること", strings.Repeat("hoge1", 1000) + "1", errors.New("本文は10000文字以下である必要があります")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := NewBody(tt.input)

			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("got: %v, want: %v", err, tt.want)
				}
			} else {
				if body.Value() != tt.input {
					t.Errorf("got: %v, want: %v", body.Value(), tt.input)
				}
			}
		})
	}
}
