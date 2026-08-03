package output

import (
	"context"

	"charm.land/log/v2"
)

func (o *Output) HandleRiverOutputV1Removed(ctx context.Context) {
	log.Info("output removed", "id", o.Object)
	o.Removed = true
}
