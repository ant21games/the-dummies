package broker

import (
	"context"
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/orders"
)

type TurnContextFactory interface {
	CreateTurnContext(ctx context.Context, msg GameMessage) TurnContext
}

type Broker interface {
	Dial(configuration *Configuration) (context.Context, error)
	Stop() error
	SendOrders(message string, ordersList ...orders.Order) error
	SetTurnContextFactory(TurnContextFactory)
	OnMessage(func(msg GameMessage))
	OnListeningState(func(turnTx TurnContext))
}

type TurnContext interface {
	context.Context
	Logger() Logger
	Reader() GameStateReader
	GameMsg() *GameMessage
}

// A GameStateReader is a helper to optimise methods that are heavily used by bots.
// The GameStateReader usage is optional. All data provided by this interface may also be directly retrieved
// from the GameMessage
type GameStateReader interface {
	Ball() Ball
	Turn() int
	Me() *Player
	GetMyTeam() Team
	GetOpponentTeam() Team
	ForEachPlayer(place arena.TeamPlace, callback func(index int, player *Player))
	FindPlayer(place arena.TeamPlace, playerNumber string) (*Player, error)
	IHoldTheBall() bool
	OpponentGoal() arena.Goal
	DefenseGoal() arena.Goal
	AmIGoalkeeper() bool
}

type Logger interface {
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}
