package wm

import (
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

func (wm *WindowManager) GetDefaultOutput() window.Output {
	if wm.DefaultOutput == nil {
		return nil
	}
	return wm.DefaultOutput
}
