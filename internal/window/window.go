package window

import (
	"time"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
)

type Window struct {
	proto.RiverWindowV1Stub

	Object proto.RiverWindowV1
	Node   proto.RiverNodeV1
	WM     WindowManager
	Output Output

	SplitTarget *Window

	X, Y                int32
	Width, Height       int32
	MaxWidth, MaxHeight int32
	MinWidth, MinHeight int32

	PointerMoveRequested        Seat
	PointerResizeRequested      Seat
	PointerResizeRequestedEdges uint32

	DecorationHint uint32
	TiledHint      uint32

	New              bool
	PendingPosition  bool
	Closed           bool
	DecorationHinted bool
	TiledHinted      bool

	debounceTimer *time.Timer
}

func (w *Window) SetPosition(x, y int32) {
	w.Node.SetPosition(x, y)
	w.X, w.Y = x, y
}

func (w *Window) HintTiled(sides uint32) {
	w.TiledHint = sides
	w.TiledHinted = true
}

func (w *Window) Manage() {
	if w.New {
		log.Debug("managing new window", "window", w.Object)
		w.New = false
		w.UpdateOutput()

		if !w.WM.AddToLayout(w) {
			w.ProposeDimensions(0, 0, true)
		}
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

func (w *Window) FlushHints() {
	if w.DecorationHinted {
		log.Debug("adjusting decorations", "window", w.Object)
		w.DecorationHinted = false
		switch w.DecorationHint {
		case 0b00: // only csd supported
			fallthrough
		case 0b01: // prefered csd
			w.Object.UseCsd()
		case 0b10: // prefered ssd
			fallthrough
		case 0b11: // only ssd supported
			w.Object.UseSsd()
		default:
			log.Error("invalid decoration hint recieved", "want", "0-3", "got", w.DecorationHint)
		}
	}

	if w.TiledHinted {
		w.TiledHinted = false
		w.Object.SetTiled(w.TiledHint)
	}
}

func (w *Window) Render() {
	if w.PendingPosition {
		log.Debug("adjusting postition", "window", w.Object)
		w.PendingPosition = false

		if w.Output != nil {
			w.X = w.Output.GetX() + ((w.Output.GetWidth() - w.Width) / 2)
			w.Y = w.Output.GetY() + ((w.Output.GetHeight() - w.Height) / 2)
		} else {
			log.Error("window render loop could not find any outputs ?!??", "window", w.Object)
		}

		w.SetPosition(w.X, w.Y)
	}
}

func (w *Window) UpdateOutput() {
	var o Output
	if s := w.WM.GetLastActiveSeat(); s != nil {
		o = w.WM.OutputAt(s.GetX(), s.GetY())
	} else {
		o = w.WM.GetDefaultOutput()
	}

	if o != nil && o != w.Output {
		defOut := w.WM.GetDefaultOutput()
		old := w.Output
		w.Output = o

		if old != nil {
			if old == defOut && o != defOut {
				w.WM.RemoveFromLayout(w)
			}

			if old != defOut && o == defOut {
				w.WM.AddToLayout(w)
			}
		}
	} else {
		log.Warn("failed to update window output, retaining previous one", "window", w.Object)
		if w.Output == nil {
			log.Error("THERE IS NO PREVIOUS OUTPUT FOR THE WINDOW, THE FUCK IS HAPPENING", "window", w.Object)
		}
	}
}

func (w *Window) ProposeDimensions(width, height int32, fill bool) {
	log.Debug("new dimensions proposed", "window", w.Object, "width", width, "height", height, "fill", fill)

	if width != 0 {
		w.Width = width
	}

	if height != 0 {
		w.Height = height
	}

	oWidth := w.Output.GetWidth()
	oHeight := w.Output.GetHeight()
	oX := w.Output.GetX()
	oY := w.Output.GetY()

	if w.Output != nil {
		if oWidth < w.Width || fill {
			w.Width = oWidth
		}
		if oHeight < w.Height || fill {
			w.Height = oHeight
		}
	} else {
		log.Error("undefined output for window, could not propose dimensions, handing decision off to the client", "window", w.Object)
	}

	if w.X+w.Width > oX+oWidth ||
		w.Y+w.Height > oY+oHeight ||
		w.Y < oY || w.X < oX {
		w.PendingPosition = true
	}

	// 0,0 is a special case that lets the window decide on its own
	w.Object.ProposeDimensions(w.Width, w.Height)
}

func (w *Window) MaybeDestroy() bool {
	if !w.Closed {
		return false
	}

	w.WM.RemoveFromLayout(w)
	w.Node.Destroy()
	w.Object.Destroy()
	return true
}

func NewWindow(object proto.RiverWindowV1, wm WindowManager) *Window {
	window := &Window{
		Object:          object,
		Node:            object.GetNode(),
		WM:              wm,
		New:             true,
		PendingPosition: true,
	}

	object.SetUserData(window)
	return window
}
