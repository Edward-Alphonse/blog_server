package views

type Handler interface {
	Handle()
	makeResponse(message string)
}
