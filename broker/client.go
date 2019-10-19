package broker

import (
	"context"
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/orders"
	"github.com/sirupsen/logrus"
)

type TurnContextFactory interface {
	CreateTurnContext(msg GameMessage) TurnContext
}

type Broker interface {
	Dial(configuration *Configuration) (context.Context, error)
	Stop() error
	SendOrders(message string, ordersList ...orders.Order) error
	OnMessage(func(msg GameMessage))
	OnAnnouncement(func(turnTx TurnContext))
}

type TurnContext interface {
	context.Context
	Logger() *logrus.Entry
	Reader() GameReader
	GameMsg() *GameMessage
}

type GameReader interface {
	Ball() Ball
	Turn() int
	Me() *Player
	GetMyTeam() Team
	GetOpponentTeam() Team
	ForEachPlayByTeam(place arena.TeamPlace, callback func(index int, player *Player))
	FindPlayer(place arena.TeamPlace, playerNumber arena.PlayerNumber) (*Player, error)
}
