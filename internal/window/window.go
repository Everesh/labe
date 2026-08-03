package window

import "codeberg.org/everesh/labe/internal/proto"

type Window struct {
	proto.RiverWindowV1Stub

	Object proto.RiverWindowV1
	Node   proto.RiverNodeV1

	New    bool
	Closed bool
}

func (w *Window) Manage() {
	if w.New {
		w.New = false
		w.Node.SetPosition(0, 0)
		w.Object.ProposeDimensions(0, 0)
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
