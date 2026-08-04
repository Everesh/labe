package wm

import (
	"context"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/output"
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/seat"
	"codeberg.org/everesh/labe/internal/window"
)

func (wm *WindowManager) HandleWlRegistryGlobal(ctx context.Context, name uint32, iface string, version uint32) {
	log.Debug("global advertised", "iface", iface, "version", version)
	switch iface {
	case proto.RiverWindowManagerV1Name:
		if version >= 4 {
			wm.WindowManagerV1 = proto.As[proto.RiverWindowManagerV1](wm.Registry.Bind(name, iface, 4))
		} else {
			log.Error("interface version too low to bind", "iface", proto.RiverWindowManagerV1Name, "want", ">=4", "got", version)
		}
	case proto.RiverXkbBindingsV1Name:
		wm.XkbBindingsV1 = proto.As[proto.RiverXkbBindingsV1](wm.Registry.Bind(name, iface, 1))
	case proto.RiverLayerShellV1Name:
		wm.LayerShellV1 = proto.As[proto.RiverLayerShellV1](wm.Registry.Bind(name, iface, 1))
	}
}

func (wm *WindowManager) HandleRiverWindowManagerV1Window(ctx context.Context, id proto.RiverWindowV1) {
	log.Info("window added", "window", id)
	wm.Windows = append(wm.Windows, window.NewWindow(id))
}

func (wm *WindowManager) HandleRiverWindowManagerV1Seat(ctx context.Context, id proto.RiverSeatV1) {
	log.Info("seat added", "seat", id)
	wm.Seats = append(wm.Seats, seat.NewSeat(id, wm))
}

func (wm *WindowManager) HandleRiverWindowManagerV1Output(ctx context.Context, id proto.RiverOutputV1) {
	log.Info("output added", "output", id)
	wm.Outputs = append(wm.Outputs, output.NewOutput(id))
}
