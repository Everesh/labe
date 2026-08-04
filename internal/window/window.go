package window

import (
	"codeberg.org/everesh/labe/internal/proto"
)

type Window struct {
	proto.RiverWindowV1Stub

	Object proto.RiverWindowV1
	Node   proto.RiverNodeV1

	X, Y          int32
	Width, Height int32

	PointerMoveRequested        Seat
	PointerResizeRequested      Seat
	PointerResizeRequestedEdges uint32

	New    bool
	Closed bool
}

func (w *Window) SetPosition(x, y int32) {
	w.Node.SetPosition(x, y)
	w.X, w.Y = x, y
}

func (w *Window) Manage() {
	if w.New {
		w.New = false
		w.Node.SetPosition(0, 0)
		w.Object.ProposeDimensions(0, 0)
	}

	if w.PointerMoveRequested != nil {
		w.PointerMoveRequested.PointerMove(w)
		w.PointerMoveRequested = nil
	}

	if w.PointerResizeRequested != nil {
		w.PointerResizeRequested.PointerResize(w, w.PointerResizeRequestedEdges)
		w.PointerResizeRequested = nil
	}
}

func (w *Window) MaybeDestroy() bool {
	if !w.Closed {
		return false
	}

	w.Node.Destroy()
	w.Object.Destroy()
	return true
}

func NewWindow(object proto.RiverWindowV1) *Window {
	window := &Window{
		Object: object,
		Node:   object.GetNode(),
		New:    true,
	}

	object.SetUserData(window)
	return window
}
