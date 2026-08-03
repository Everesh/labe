package wm

import (
	"context"
	"slices"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/output"
	"codeberg.org/everesh/labe/internal/seat"
	"codeberg.org/everesh/labe/internal/window"
)

func (wm *WindowManager) HandleRiverWindowManagerV1ManageStart(ctx context.Context) {
	log.Info("manage")
	wm.Windows = slices.DeleteFunc(wm.Windows, (*window.Window).MaybeDestroy)
	wm.Seats = slices.DeleteFunc(wm.Seats, (*seat.Seat).MaybeDestroy)
	wm.Outputs = slices.DeleteFunc(wm.Outputs, (*output.Output).MaybeDestroy)

	for _, w := range wm.Windows {
		w.Manage()
	}

	for _, s := range wm.Seats {
		s.Manage()
	}

	wm.WindowManagerV1.ManageFinish()
}
