package window

// required

type Seat interface {
	PointerMove(w *Window)
	PointerResize(w *Window, edges uint32)
}

// advertised
