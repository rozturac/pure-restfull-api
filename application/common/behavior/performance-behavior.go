package behavior

import (
	"context"
	"fmt"
	"github.com/rozturac/go-mediator"
	"log"
	"pure-restfull-api/application/common"
	"time"
)

type PerformanceBehavior struct {
	logger *log.Logger
}

func NewPerformanceBehavior(logger *log.Logger) *PerformanceBehavior {
	return &PerformanceBehavior{
		logger: logger,
	}
}

func (p *PerformanceBehavior) Execute(ctx context.Context, command mediator.Command, next mediator.CommandHandler) (interface{}, error) {
	defer func(start time.Time) {
		if time.Since(start).Milliseconds() > 500 {
			p.logger.Println(fmt.Sprintf("Performance Warning Log Command:%v, CorrelationId: %v, Duration(ms): %v", command, ctx.Value(common.CorrelationId), time.Since(start).Milliseconds()))
		}
	}(time.Now())
	return next(ctx, command)
}
