package wrappers

import "fmt"

type TransparentShape struct {
	ColorShape   ColoredShape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.ColorShape.Render(), t.Transparency*100.0)
}
