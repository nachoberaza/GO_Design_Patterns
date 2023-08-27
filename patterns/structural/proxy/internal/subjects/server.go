package subjects

type Server interface {
	handleRequest(string, string) (int, string)
}
