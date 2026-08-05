package wm

import (
	"context"
	"slices"

	"codeberg.org/everesh/labe/internal/output"
	"codeberg.org/everesh/labe/internal/seat"
	"codeberg.org/everesh/labe/internal/window"
)

func (wm *WindowManager) HandleRiverWindowManagerV1ManageStart(ctx context.Context) {
	wm.Windows = slices.DeleteFunc(wm.Windows, (*window.Window).MaybeDestroy)
	wm.Seats = slices.DeleteFunc(wm.Seats, (*seat.Seat).MaybeDestroy)
	wm.Outputs = slices.DeleteFunc(wm.Outputs, (*output.Output).MaybeDestroy)

	// if no default layer shell, pick the first in the slice
	if wm.DefaultLayerShell == nil || wm.DefaultLayerShell.Removed {
		if len(wm.Outputs) > 0 {
			wm.Outputs[0].SetDefaultLayerShell()
			wm.DefaultLayerShell = wm.Outputs[0]
		} else {
			wm.DefaultLayerShell = nil
		}
	}

	// if no default output, pick the first in the slice
	// used when there is no last active seat from which to infer active output
	if wm.DefaultOutput == nil || wm.DefaultOutput.Removed {
		if len(wm.Outputs) > 0 {
			wm.DefaultOutput = wm.Outputs[0]
		} else {
			wm.DefaultOutput = nil
		}
	}

	for _, w := range wm.Windows {
		w.Manage()
	}

	for _, s := range wm.Seats {
		s.Manage()
	}

	for _, o := range wm.Outputs {
		o.Manage()
	}

	wm.WindowManagerV1.ManageFinish()
}
