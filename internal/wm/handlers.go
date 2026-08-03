package wm

import (
	"context"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/seat"
	"codeberg.org/everesh/labe/internal/window"
)

func (wm *WindowManager) HandleWlRegistryGlobal(ctx context.Context, name uint32, iface string, version uint32) {
	switch iface {
	case proto.RiverWindowManagerV1Name:
		if version >= 4 {
			wm.WindowManagerV1 = proto.As[proto.RiverWindowManagerV1](wm.Registry.Bind(name, iface, 4))
		} else {
			log.Error("failed to bind "+proto.RiverWindowManagerV1Name+", version too low",
				"want", ">=4", "got", version)
		}
	case proto.RiverXkbBindingsV1Name:
		wm.XkbBindingsV1 = proto.As[proto.RiverXkbBindingsV1](wm.Registry.Bind(name, iface, 1))
	}
}

func (wm *WindowManager) HandleRiverWindowManagerV1Window(ctx context.Context, id proto.RiverWindowV1) {
	wm.Windows = append(wm.Windows, window.NewWindow(id))
}

func (wm *WindowManager) HandleRiverWindowManagerV1Seat(ctx context.Context, id proto.RiverSeatV1) {
	wm.Seats = append(wm.Seats, seat.NewSeat(id))
}
