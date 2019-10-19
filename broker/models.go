package broker

import (
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/physics"
)

// GameMessage is the message sent from the game server to the player
type GameMessage struct {
	Type     arena.MsgType          `json:"type"`
	GameInfo GameInfo               `json:"info"`
	State    arena.GameState        `json:"state"`
	Data     map[string]interface{} `json:"data"`
	// Message is quite useless, but could help the developers to debug the game server messages
	Message string `json:"message"`
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

type Player struct {
	physics.Element
	Id        string             `json:"id"`
	Number    arena.PlayerNumber `json:"number"`
	TeamPlace arena.TeamPlace    `json:"team_place"`
}
