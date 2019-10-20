package broker

import (
	"errors"
	"github.com/makeitplay/arena"
)

var PlayerNotFound = errors.New("player not found")

// NewGameStateReader returns the default GameStateReader
func NewGameStateReader(gameInfo GameInfo, place arena.TeamPlace, number string) GameStateReader {
	teamState := gameInfo.HomeTeam
	if place == arena.AwayTeam {
		teamState = gameInfo.AwayTeam
	}

	var player *Player
	for _, playerInfo := range teamState.Players {
		if string(playerInfo.Number) == number {
			player = playerInfo
		}
	}
	return &reader{
		player:       player,
		GameInfo:     gameInfo,
		teamPlace:    place,
		playerNumber: number,
		teams: map[arena.TeamPlace]Team{
			arena.AwayTeam: gameInfo.AwayTeam,
			arena.HomeTeam: gameInfo.HomeTeam,
		},
	}
}

type reader struct {
	GameInfo     GameInfo
	teamPlace    arena.TeamPlace
	playerNumber string
	player       *Player
	teams        map[arena.TeamPlace]Team //optimising reading
}

func (r *reader) Ball() Ball {
	return r.GameInfo.Ball
}

func (r *reader) Turn() int {
	return r.GameInfo.Turn
}

func (r *reader) Me() *Player {
	return r.player
}

func (r *reader) GetMyTeam() Team {
	return r.teams[r.teamPlace]
}

func (r *reader) GetOpponentTeam() Team {
	otherSite := arena.AwayTeam
	if r.teamPlace == otherSite {
		otherSite = arena.HomeTeam
	}
	return r.teams[otherSite]
}

func (r *reader) ForEachPlayer(place arena.TeamPlace, callback func(index int, player *Player)) {
	for i, p := range r.teams[place].Players {
		callback(i, p)
	}
}

func (r *reader) FindPlayer(place arena.TeamPlace, playerNumber string) (*Player, error) {
	for _, playerInfo := range r.teams[place].Players {
		if string(playerInfo.Number) == playerNumber {
			return playerInfo, nil
		}
	}
	return nil, PlayerNotFound
}

func (r *reader) IHoldTheBall() bool {
	return r.GameInfo.Ball.Holder != nil && r.player != nil && r.GameInfo.Ball.Holder.Id == r.player.Id
}

func (r *reader) OpponentGoal() arena.Goal {
	if r.teamPlace == arena.HomeTeam {
		return arena.AwayTeamGoal
	}
	return arena.HomeTeamGoal
}

func (r *reader) DefenseGoal() arena.Goal {
	if r.teamPlace == arena.HomeTeam {
		return arena.HomeTeamGoal
	}
	return arena.AwayTeamGoal
}

func (r *reader) AmIGoalkeeper() bool {
	return r.playerNumber == string(arena.GoalkeeperNumber)
}
