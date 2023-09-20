package observers

type Observer interface {
	Notify(data interface{})
}
