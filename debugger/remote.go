package debugger

import (
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/orders"
	"github.com/makeitplay/arena/physics"
	"github.com/makeitplay/client-player-go/broker"
)

type DebugCmd struct {
	Cmd  string
	Args map[string]interface{}
}

type TrainingQuestion struct {
	Question     string   `json:"question"`
	PlayerId     string   `json:"player_id"`
	QuestionId   string   `json:"question_id"`
	Alternatives []string `json:"alternatives"`
}

type DebugSender interface {
	SendCommand(msg DebugCmd) (newState broker.GameMessage, err error)
}

type Controller interface {
	SetDebugSender(DebugSender)
	NextTurn() (newState broker.GameMessage, err error)
	LoadArrangement(name string) (newState broker.GameMessage, err error)
	SetBallProperties(v physics.Velocity, position physics.Point) (newState broker.GameMessage, err error)
	SetPlayerPosition(place arena.TeamPlace, number arena.PlayerNumber, position physics.Point) (newState broker.GameMessage, err error)
	SetGameTurn(turn int) (newState broker.GameMessage, err error)
	ResetScore() (newState broker.GameMessage, err error)
	AskQuestion(question TrainingQuestion, confg broker.Configuration) (err error)
}

type Mimic interface {
	CreatePlayers() error
	GetLastTurnContext(place arena.TeamPlace, number arena.PlayerNumber) (ctx broker.TurnContext, err error)
	GetPlayerBroker(place arena.TeamPlace, number arena.PlayerNumber) (ctx broker.Broker, err error)
	SendOrders(place arena.TeamPlace, number arena.PlayerNumber, orderList []orders.Order)
}
