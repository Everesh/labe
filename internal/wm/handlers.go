package wm

import (
	"context"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
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

func (wm *WindowManager) HandlerRiverWindowManagerV1RenderStart(ctx context.Context) {

	// TODO
	log.Info("render")

	wm.WindowManagerV1.RenderFinish()
}

func (wm *WindowManager) HandleRiverWindowManagerV1ManageStart(ctx context.Context) {

	// TODO
	log.Info("manage")

	wm.WindowManagerV1.ManageFinish()
}
