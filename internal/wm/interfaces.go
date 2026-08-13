package wm

import (
	"codeberg.org/everesh/labe/internal/layout/bsp"
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/seat"
	"codeberg.org/everesh/labe/internal/window"
)

// required

// advertised

func (wm *WindowManager) GetXkbBindingsV1() proto.RiverXkbBindingsV1 {
	return wm.XkbBindingsV1
}

func (wm *WindowManager) GetWindowManagerV1() proto.RiverWindowManagerV1 {
	return wm.WindowManagerV1
}

func (wm *WindowManager) GetWindows() []*window.Window {
	return wm.Windows
}

func (wm *WindowManager) MarkSeatLastActive(s *seat.Seat) {
	wm.LastActiveSeat = s
}

func (wm *WindowManager) GetLastActiveSeat() window.Seat {
	if wm.LastActiveSeat == nil {
		return nil
	}
	return wm.LastActiveSeat
}

func (wm *WindowManager) GetLayerShellOutput(output proto.RiverOutputV1) proto.RiverLayerShellOutputV1 {
	return wm.LayerShellV1.GetOutput(output)
}

func (wm *WindowManager) OutputAt(x, y int32) window.Output {
	for _, o := range wm.Outputs {
		if o.Contains(x, y) {
			return o
		}
	}
	return nil
}

func (wm *WindowManager) GetDefaultOutput() window.Output {
	if wm.DefaultOutput == nil {
		return nil
	}
	return wm.DefaultOutput
}

func (wm *WindowManager) AddToBSP(w *window.Window) bool {
	// TODO - only tile the default output for now
	if wm.DefaultOutput == nil || w.Output != wm.DefaultOutput {
		return false
	}

	if wm.Bsp.Find(w) != nil {
		return true
	}

	if wm.Bsp == nil {
		wm.Bsp = bsp.New(w, nil)
		wm.AlignBSP()
		return true
	}

	target := wm.Bsp
	if t := w.SplitTarget; t != nil {
		if found := wm.Bsp.Find(t); found != nil {
			target = found
		}
	}
	w.SplitTarget = nil

	bsp.New(w, target)

	wm.AlignBSP()
	return true
}

func (wm *WindowManager) RemoveFromBSP(w *window.Window) {
	if wm.Bsp == nil {
		return
	}

	if wm.Bsp.Remove(w) {
		wm.Bsp = nil
	}

	wm.AlignBSP()
}

func (wm *WindowManager) AlignBSP() {
	if wm.Bsp == nil || wm.DefaultOutput == nil {
		return
	}

	wm.Bsp.Align(wm.DefaultOutput.GetX(), wm.DefaultOutput.GetY(), wm.DefaultOutput.GetWidth(), wm.DefaultOutput.GetHeight())
}
