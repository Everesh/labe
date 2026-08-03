package wm

import (
	"context"
	"slices"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/window"
)

func (wm *WindowManager) HandleRiverWindowManagerV1ManageStart(ctx context.Context) {
	log.Info("manage")
	wm.Windows = slices.DeleteFunc(wm.Windows, (*window.Window).MaybeDestroy)

	for _, w := range wm.Windows {
		w.Manage()
	}

	wm.WindowManagerV1.ManageFinish()
}
