package wm

import (
	"charm.land/log/v2"
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

func (wm *WindowManager) AddToLayout(w *window.Window) bool {
	// TODO - only tile the default output for now
	if wm.DefaultOutput == nil || w.Output != wm.DefaultOutput {
		return false
	}

	if err := wm.Layout.Add(w); err != nil {
		log.Error("failed to add window to layout", "window", w.Object, "error", err)
		return false
	}

	wm.AlignLayout()
	return true
}

func (wm *WindowManager) RemoveFromLayout(w *window.Window) {
	if err := wm.Layout.Remove(w); err != nil {
		log.Error("failed to remove window from layout", "window", w.Object, "error", err)
	}

	wm.AlignLayout()
}

func (wm *WindowManager) AlignLayout() {
	if wm.DefaultOutput == nil {
		return
	}

	err := wm.Layout.Align(wm.DefaultOutput.GetX(), wm.DefaultOutput.GetY(), wm.DefaultOutput.GetWidth(), wm.DefaultOutput.GetHeight())
	if err != nil {
		log.Error("failed to align layout", "error", err)
	}
}
