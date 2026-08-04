package xkb

import (
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/window"
)

// required

type Seat interface {
	GetXkbBinding(keysym uint32, mods uint32) proto.RiverXkbBindingV1
	GetPointerBinding(button uint32, mods uint32) proto.RiverPointerBindingV1
	SetPendingAction(fn func())

	GetFocused() *window.Window
	GetHovered() *window.Window
	GetPosition() (x, y int32)
	Focus(w *window.Window)
	PointerMove(w *window.Window)
	PointerResize(w *window.Window, edges uint32)
	GetWindows() []*window.Window
	ExitSession()
}

// advertised
