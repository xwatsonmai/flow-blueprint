package runTime

import (
	"context"
	"flow-blueprint/blogger"
)

type RunTime interface {
	context.Context
	blogger.ILogger
	GetTradeId() string
	IsDebugger() bool
	SetTradeId(tradeId string)
	GetGo() error
	ExecuteGo(any, ...any)
}
