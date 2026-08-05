package output

import (
	"context"

	"charm.land/log/v2"
)

func (o *Output) HandleRiverOutputV1Removed(ctx context.Context) {
	log.Info("output removed", "output", o.Object)
	o.Removed = true
}

func (o *Output) HandleRiverOutputV1Position(ctx context.Context, x int32, y int32) {
	log.Debug("output position changed", "output", o.Object, "x", x, "y", y)
	o.X, o.Y = x, y
}

func (o *Output) HandleRiverOutputV1Dimensions(ctx context.Context, width int32, height int32) {
	log.Debug("output dimensions changed", "output", o.Object, "width", width, "height", height)
	o.Width, o.Height = width, height
}
