package debugger

import (
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/orders"
	"github.com/makeitplay/arena/physics"
)

type DebugCmd struct {
	Cmd  string
	Args map[string]interface{}
}

type DebugSender interface {
	SendCommand(msg DebugCmd) (newState GameMessage, err error)
}

type Controller interface {
	SetDebugSender(DebugSender)
	NextTurn() (newState GameMessage, err error)
	LoadArrangement(name string) (newState GameMessage, err error)
	SetBallProperties(v physics.Velocity, position physics.Point) (newState GameMessage, err error)
	SetPlayerPosition(place arena.TeamPlace, number arena.PlayerNumber, position physics.Point) (newState GameMessage, err error)
	SetGameTurn(turn int) (newState GameMessage, err error)
	ResetScore() (newState GameMessage, err error)
	AskQuestion(question TrainingQuestion, confg Configuration) (err error)
}

type Mimic interface {
	CreatePlayers() error
	GetTurnContext(place arena.TeamPlace, number arena.PlayerNumber) (ctx TurnContext, err error)
	GetPlayerBroker(place arena.TeamPlace, number arena.PlayerNumber) (ctx Broker, err error)
	SendOrders(place arena.TeamPlace, number arena.PlayerNumber, orderList []orders.Order)
}
