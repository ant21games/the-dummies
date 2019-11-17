package client

import (
	"github.com/lugobots/arena"
	"github.com/lugobots/arena/orders"
	"github.com/lugobots/arena/physics"
)

// Ball is the ball :-)
type Ball struct {
	physics.Element
	// Holder identifies the player who is holding the ball
	Holder *Player
}

// Team groups the player team info based on the status sent by the game server
type Team struct {
	Name    arena.TeamPlace `json:"name"`
	Score   int             `json:"score"`
	Players []*Player       `json:"players"`
}

// PlayerMessage is the message sent from a player to the game server
type PlayerMessage struct {
	Type   arena.MsgType  `json:"type"`
	Orders []orders.Order `json:"orders"`
	// Debug is a message the will be only visible in the game server log (used for debugging purposes)
	Debug string `json:"message"`
}

// GameMessage is the message sent from the game server to the player
type GameMessage struct {
	Type     arena.MsgType          `json:"type"`
	GameInfo GameInfo               `json:"info"`
	State    arena.GameState        `json:"state"`
	Data     map[string]interface{} `json:"data"`
	// Message is quite useless, but could help the developers to debug the game server messages
	Message string `json:"message"`
}

func (g *GameMessage) Ball() Ball {
	return g.GameInfo.Ball
}

func (g *GameMessage) Turn() int {
	return g.GameInfo.Turn
}

func (g *GameMessage) GetTeam(place arena.TeamPlace) Team {
	team := g.GameInfo.HomeTeam
	if place == arena.AwayTeam {
		team = g.GameInfo.AwayTeam
	}
	return team
}

func (g *GameMessage) ForEachPlayByTeam(place arena.TeamPlace, callback func(index int, player *Player)) {
	for i, p := range g.GetTeam(place).Players {
		callback(i, p)
	}
}

// GameInfo is the set of values that defines the current game state
type GameInfo struct {
	State arena.GameState `json:"state"`
	// Turn is the sequential number of turns. Read the game documentation to understand what a turn is
	Turn     int  `json:"turn"`
	Ball     Ball `json:"ball"`
	HomeTeam Team `json:"home"`
	AwayTeam Team `json:"away"`
}

type Player struct {
	physics.Element
	Id        string             `json:"id"`
	Number    arena.PlayerNumber `json:"number"`
	TeamPlace arena.TeamPlace    `json:"team_place"`
}
