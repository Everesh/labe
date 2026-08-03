package output

import "context"

func (o *Output) HandleRiverOutputV1Removed(ctx context.Context) {
	o.Removed = true
}
