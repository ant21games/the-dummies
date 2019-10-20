package broker

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
)

type turnCtxFactory struct {
	logger Logger
}

func (turnCtxFactory) CreateTurnContext(ctx context.Context, msg GameMessage) TurnContext {
	ctx, cancelFunc := context.WithCancel(ctx)
	logger := logrus.New()
	logger.SetLevel(config.LogLevel)
	log.Warnf()
	return &gameCtx{
		config:  config,
		mainCtx: ctx,
		log:     logger.WithField("player", fmt.Sprintf("%s-%s", config.TeamPlace, config.PlayerNumber)),
	}, cancelFunc
}
