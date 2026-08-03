package wm

import (
	"context"
	"fmt"
	"os"

	"codeberg.org/everesh/labe/internal/proto"
	"hazelnut.eclair.cafe/wlcl"
)

type WindowManager struct {
	proto.WlRegistryStub
	proto.RiverWindowManagerV1Stub

	Registry        proto.WlRegistry
	WindowManagerV1 proto.RiverWindowManagerV1
	XkbBindingsV1   proto.RiverXkbBindingsV1

	Done bool
	Err  error
}

func New(ctx context.Context, conn *wlcl.Connection) (*WindowManager, error) {
	// as per the tinyrwm example,
	// safeguard against the wayland debug env var to prevent log pollution
	_ = os.Unsetenv("WAYLAND_DEBUG")
	display := proto.CreateDisplay(conn)

	wm := &WindowManager{}
	wm.Registry = display.GetRegistry()
	wm.Registry.SetUserData(wm)

	if err := wlcl.Roundtrip(ctx, display); err != nil {
		return nil, fmt.Errorf("wlcl roundtrip failed: %w", err)
	}

	if !wm.Registry.IsSet() ||
		!wm.WindowManagerV1.IsSet() ||
		!wm.XkbBindingsV1.IsSet() {
		return nil, fmt.Errorf("failed to register all globals")
	}

	wm.WindowManagerV1.SetUserData(wm)
	return wm, nil
}
