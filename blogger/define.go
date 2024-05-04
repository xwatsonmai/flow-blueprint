package blogger

import "context"

type ILogger interface {
	Info(ctx context.Context, msg ...interface{})
	Error(ctx context.Context, msg ...interface{})
}
