package seat

import (
	"codeberg.org/everesh/labe/internal/window"
)

type OpMove struct {
	StartX, StartY int32
}

func (s *OpMove) Manage(state *OpState) {}

func (s *OpMove) Render(state *OpState) {
	state.Window.SetPosition(s.StartX+state.Dx, s.StartY+state.Dy)
}

func (s *OpMove) InformStart(w *window.Window) {}

func (s *OpMove) InformEnd(w *window.Window) {
	w.UpdateOutput()
	w.ProposeDimensions(false)
}

func NewOpMove(w *window.Window) *OpMove {
	return &OpMove{
		StartX: w.X,
		StartY: w.Y,
	}
}
