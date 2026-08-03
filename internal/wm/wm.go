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

	Registry proto.WlRegistry

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
	return wm, nil
}
