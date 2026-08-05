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

	if o.NonExclusiveX == 0 || o.NonExclusiveY == 0 {
		o.NonExclusiveX, o.NonExclusiveY = x, y
	}
}

func (o *Output) HandleRiverOutputV1Dimensions(ctx context.Context, width int32, height int32) {
	log.Debug("output dimensions changed", "output", o.Object, "width", width, "height", height)
	o.Width, o.Height = width, height

	if o.NonExclusiveWidth == 0 || o.NonExclusiveHeight == 0 {
		o.NonExclusiveWidth, o.NonExclusiveHeight = width, height
	}
}

func (o *Output) HandleRiverLayerShellOutputV1NonExclusiveArea(ctx context.Context, x int32, y int32, width int32, height int32) {
	log.Debug("output non exclusive area changed", "output", o.Object, "width", width, "height", height)
	o.NonExclusiveX, o.NonExclusiveY = x, y
	o.NonExclusiveWidth, o.NonExclusiveHeight = width, height
}
