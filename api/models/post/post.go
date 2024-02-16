package post

type Post struct {
	active bool
	body   Body
}

func NewPost(body Body) (*Post, error) {
	return &Post{
		active: true,
		body:   body,
	}, nil
}

func (p *Post) updateActive(active bool) error {
	p.active = active
	return nil
}

func (p *Post) updateBody(newBody string) error {
	body, err := NewBody(newBody)
	if err != nil {
		return err
	}
	p.body = *body
	return nil
}

func (p *Post) GetActive() bool {
	return p.active
}

func (p *Post) GetBody() string {
	return p.body.Value()
}
