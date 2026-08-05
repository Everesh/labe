package wm

import (
	"context"
)

func (wm *WindowManager) HandleRiverWindowManagerV1RenderStart(ctx context.Context) {

	for _, w := range wm.Windows {
		w.Render()
	}

	for _, s := range wm.Seats {
		s.Render()
	}

	wm.WindowManagerV1.RenderFinish()
}
