package shapes

import "fmt"

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.Width = size
	sq.Height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.Width = width
	s.Height = width
}

func (s *Square) SetHeight(height int) {
	s.Width = height
	s.Height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}
