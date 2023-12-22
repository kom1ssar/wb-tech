package event

import "context"

type OrderSubscriptions interface {
	Subscribe(_ context.Context)
}
