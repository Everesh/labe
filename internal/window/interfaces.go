package window

// required

type Seat interface {
	PointerMove(w *Window)
	PointerResize(w *Window, edges uint32)
	GetX() int32
	GetY() int32
}

type WindowManager interface {
	OutputAt(x, y int32) Output
	GetDefaultOutput() Output
	GetLastActiveSeat() Seat
}

type Output interface {
	GetX() int32
	GetY() int32
	GetWidth() int32
	GetHeight() int32
}

// advertised
