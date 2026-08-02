package main

import (
	"context"
	"os"

	"codeberg.org/everesh/labe/internal/proto"
	"hazelnut.eclair.cafe/wlcl"
)

func main() {
	ctx := context.Background()
	conn, err := wlcl.Connect(ctx, "")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	_ = os.Unsetenv("WAYLAND_DEBUG")

	display := proto.CreateDisplay(conn)
	if err := wlcl.Roundtrip(ctx, display); err != nil {
		panic(err)
	}

	// TODO
}
