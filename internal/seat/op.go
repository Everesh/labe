package seat

import "codeberg.org/everesh/labe/internal/window"

type Op interface {
	Manage(state *OpState)
	Render(state *OpState)
	InformStart(w *window.Window)
	InformEnd(w *window.Window)
}

type OpState struct {
	Window *window.Window
	Dx, Dy int32
}
