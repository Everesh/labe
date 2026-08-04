package wm

import (
	"context"
)

func (wm *WindowManager) HandleRiverWindowManagerV1RenderStart(ctx context.Context) {

	for _, s := range wm.Seats {
		s.Render()
	}

	wm.WindowManagerV1.RenderFinish()
}
