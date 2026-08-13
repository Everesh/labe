package seat

import (
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/window"
)

var _ Op = (*OpResize)(nil)

type OpResize struct {
	StartX, StartY          int32
	StartWidth, StartHeight int32
	Edges                   uint32
}

func (s *OpResize) Manage(state *OpState) {
	w, h := s.StartWidth, s.StartHeight
	switch {
	case s.Edges&proto.RiverWindowV1EdgesLeft != 0:
		w -= state.Dx
	case s.Edges&proto.RiverWindowV1EdgesRight != 0:
		w += state.Dx
	}

	switch {
	case s.Edges&proto.RiverWindowV1EdgesTop != 0:
		h -= state.Dy
	case s.Edges&proto.RiverWindowV1EdgesBottom != 0:
		h += state.Dy
	}

	state.Window.Object.ProposeDimensions(max(w, 1), max(h, 1))
}

func (s *OpResize) Render(state *OpState) {
	x := s.StartX
	if s.Edges&proto.RiverWindowV1EdgesLeft != 0 {
		x += s.StartWidth - state.Window.Width
	}

	y := s.StartY
	if s.Edges&proto.RiverWindowV1EdgesTop != 0 {
		y += s.StartHeight - state.Window.Height
	}

	state.Window.SetPosition(x, y)
}

func (s *OpResize) InformStart(w *window.Window) {
	w.Object.InformResizeStart()
}

func (s *OpResize) InformEnd(w *window.Window) {
	w.Object.InformResizeEnd()
	w.ProposeDimensions(0, 0, false)
}

func NewOpResize(w *window.Window, edges uint32) *OpResize {
	return &OpResize{
		StartX:      w.X,
		StartY:      w.Y,
		StartWidth:  w.Width,
		StartHeight: w.Height,
		Edges:       edges,
	}
}
