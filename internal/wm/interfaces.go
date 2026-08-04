package wm

import (
	"codeberg.org/everesh/labe/internal/proto"
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
