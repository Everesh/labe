package wm

import (
	"context"

	"charm.land/log/v2"
)

func (wm *WindowManager) HandlerRiverWindowManagerV1RenderStart(ctx context.Context) {

	// TODO
	log.Info("render")

	wm.WindowManagerV1.RenderFinish()
}
