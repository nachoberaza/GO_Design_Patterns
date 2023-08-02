package shapes

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}
