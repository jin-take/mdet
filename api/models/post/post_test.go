package post

import (
	"testing"
)

func TestPost(t *testing.T) {
	t.Run("値オブジェクトを引数に渡して新しい投稿を作成できること", func(t *testing.T) {
		body, _ := NewBody("post body text.")
		post, err := NewPost(*body)

		if err != nil {
			t.Errorf("got: %v, want: %v", err, nil)
		}

		if post.GetActive() != true {
			t.Errorf("got: %v, want: %v", post.GetActive(), true)
		}

		if post.GetBody() != "post body text." {
			t.Errorf("got: %v, want: %v", post.GetBody(), "post body text.")
		}
	})

	t.Run("投稿の本文を更新できること", func(t *testing.T) {
		body, _ := NewBody("new post body text.")
		post, _ := NewPost(*body)

		err := post.updateBody("new post body text.")
		if err != nil {
			t.Errorf("got: %v, want: %v", err, nil)
		}

		if post.GetBody() != "new post body text." {
			t.Errorf("got: %v, want: %v", post.GetBody(), "new post body text.")
		}
	})

	t.Run("投稿のアクティブ状態を更新できること", func(t *testing.T) {
		body, _ := NewBody("post body text.")
		post, _ := NewPost(*body)

		err := post.updateActive(false)
		if err != nil {
			t.Errorf("got: %v, want: %v", err, nil)
		}

		if post.GetActive() != false {
			t.Errorf("got: %v, want: %v", post.GetActive(), false)
		}
	})
}
