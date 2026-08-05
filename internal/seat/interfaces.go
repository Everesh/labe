package seat

import (
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/window"
)

// required

type WindowManager interface {
	GetXkbBindingsV1() proto.RiverXkbBindingsV1
	GetWindowManagerV1() proto.RiverWindowManagerV1
	GetWindows() []*window.Window
	MarkSeatLastActive(s *Seat)
}

// advertised

func (s *Seat) GetXkbBinding(keysym uint32, mods uint32) proto.RiverXkbBindingV1 {
	return s.WM.GetXkbBindingsV1().GetXkbBinding(s.Object, keysym, mods)
}

func (s *Seat) GetPointerBinding(button uint32, mods uint32) proto.RiverPointerBindingV1 {
	return s.Object.GetPointerBinding(button, mods)
}

func (s *Seat) SetPendingAction(fn func()) {
	s.WM.MarkSeatLastActive(s)
	s.PendingAction = fn
}

func (s *Seat) MarkSeatLastActive() {
	s.WM.MarkSeatLastActive(s)
}

func (s *Seat) GetFocused() *window.Window {
	return s.Focused
}

func (s *Seat) GetHovered() *window.Window {
	return s.Hovered
}

func (s *Seat) GetPosition() (x, y int32) {
	return s.X, s.Y
}

func (s *Seat) GetWindows() []*window.Window {
	return s.WM.GetWindows()
}

func (s *Seat) GetX() int32 {
	return s.X
}

func (s *Seat) GetY() int32 {
	return s.Y
}
