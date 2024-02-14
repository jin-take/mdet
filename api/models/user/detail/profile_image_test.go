package detail

import "testing"

func TestProfileImage(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"正常に登録できること", "sample.jpg", "sample.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profileImage, _ := NewProfileImage(tt.input)

			if profileImage.Value() != tt.want {
				t.Errorf("got: %v, want: %v", profileImage.Value(), tt.want)
			}
		})
	}
}
