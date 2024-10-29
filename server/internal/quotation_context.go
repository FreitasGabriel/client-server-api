package internal

import (
	"context"
	"time"
)

func GetContext(cont context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(cont, timeout)
	return ctx, cancel
}
