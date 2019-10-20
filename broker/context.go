package broker

import (
	"context"
	"time"
)

//type turnCtxFactory struct {
//
//}
//
//func (t turnCtxFactory) CreateTurnContext(gameMsg GameMessage, place arena.TeamPlace, number string) (TurnContext, context.CancelFunc) {
//	ctx, breaker := context.WithCancel(context.Background())
//
//	reader := NewGameStateReader(gameMsg.GameInfo, place, number)
//
//	return &turnCtx{
//		ctx: ctx,
//		logger: logger,
//		reader: reader,
//		gameMsg: &gameMsg,
//	}, breaker
//}

type turnCtx struct {
	ctx     context.Context
	logger  Logger
	reader  GameStateReader
	gameMsg *GameMessage
}

func (t *turnCtx) Deadline() (deadline time.Time, ok bool) {
	return t.ctx.Deadline()
}

func (t *turnCtx) Done() <-chan struct{} {
	return t.ctx.Done()
}

func (t *turnCtx) Err() error {
	return t.ctx.Err()
}

func (t *turnCtx) Value(key interface{}) interface{} {
	return t.ctx.Value(key)
}

func (t *turnCtx) Logger() Logger {
	return t.logger
}

func (t *turnCtx) Reader() GameStateReader {
	return t.reader
}

func (t *turnCtx) GameMsg() *GameMessage {
	return t.gameMsg
}
