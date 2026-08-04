package seat

import (
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/window"
	"codeberg.org/everesh/labe/internal/xkb"
)

type Seat struct {
	proto.RiverSeatV1Stub

	Object          proto.RiverSeatV1
	WM              WindowManager
	Bindings        []*xkb.Binding
	PointerBindings []*xkb.PointerBinding

	Focused    *window.Window
	Hovered    *window.Window
	Interacted *window.Window
	X, Y       int32

	PendingAction func()

	Op         Op
	OpState    OpState
	OpReleased bool

	New     bool
	Removed bool
}

func (s *Seat) Manage() {
	if s.New {
		s.New = false
		for _, b := range s.Bindings {
			b.Object.Enable()
		}
		for _, b := range s.PointerBindings {
			b.Object.Enable()
		}
	}

	if w := s.Focused; w != nil && w.Closed {
		s.Focused = nil
	}

	s.Focus(s.Interacted)
	s.Interacted = nil

	if s.PendingAction != nil {
		s.PendingAction()
		s.PendingAction = nil
	}

	if op := s.Op; op != nil {
		switch {
		case s.OpReleased:
			op.InformEnd(s.OpState.Window)
			fallthrough
		case s.OpState.Window.Closed:
			s.Object.OpEnd()
			s.Op = nil
		default:
			op.Manage(&s.OpState)
		}
	}
	s.OpReleased = false
}

func (s *Seat) Render() {
	if op := s.Op; op != nil {
		op.Render(&s.OpState)
	}
}

func (s *Seat) StartOp(w *window.Window, op Op) {
	if s.Op != nil {
		return
	}

	s.Op = op
	s.OpState = OpState{
		Window: w,
	}

	s.Focus(w)

	s.Object.OpStartPointer()
	op.InformStart(w)
}

func (s *Seat) ExitSession() {
	s.WM.GetWindowManagerV1().ExitSession()
}

func (s *Seat) PointerMove(w *window.Window) {
	s.StartOp(w, NewOpMove(w))
}

func (s *Seat) PointerResize(w *window.Window, edges uint32) {
	s.StartOp(w, NewOpResize(w, edges))
}

func (s *Seat) Focus(w *window.Window) {
	if s.Focused == w {
		return
	}

	if w != nil {
		s.Object.FocusWindow(w.Object)
		w.Node.PlaceTop()
		s.Focused = w
	}

	if s.Focused != nil && s.Focused.Closed {
		s.Object.ClearFocus()
		s.Focused = nil
	}
}

func (s *Seat) MaybeDestroy() bool {
	if !s.Removed {
		return false
	}

	s.Object.Destroy()
	return true
}

func NewSeat(object proto.RiverSeatV1, wm WindowManager) *Seat {
	seat := &Seat{
		Object: object,
		WM:     wm,
		New:    true,
	}

	seat.Bindings, seat.PointerBindings = xkb.ConfigureBindings(seat)

	seat.Object.SetUserData(seat)
	return seat
}
