package broker

import (
	"errors"
	"github.com/makeitplay/arena"
)

func NewGameReader(gameInfo GameInfo, config Configuration) GameReader {

	teamState := gameInfo.HomeTeam
	if config.TeamPlace == arena.AwayTeam {
		teamState = gameInfo.AwayTeam
	}

	var player *Player
	for _, playerInfo := range teamState.Players {
		if playerInfo.Number == config.PlayerNumber {
			player = playerInfo
		}
	}
	return &reader{
		player:   player,
		GameInfo: gameInfo,
		mySide:   config.TeamPlace,
		teams: map[arena.TeamPlace]Team{
			arena.AwayTeam: gameInfo.AwayTeam,
			arena.HomeTeam: gameInfo.HomeTeam,
		},
	}
}

type reader struct {
	GameInfo GameInfo
	mySide   arena.TeamPlace
	player   *Player
	teams    map[arena.TeamPlace]Team //optimising reading
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
	return r.teams[r.mySide]
}

func (r *reader) GetOpponentTeam() Team {
	otherSite := arena.AwayTeam
	if r.mySide == otherSite {
		otherSite = arena.HomeTeam
	}
	return r.teams[otherSite]
}

func (r *reader) ForEachPlayByTeam(place arena.TeamPlace, callback func(index int, player *Player)) {
	for i, p := range r.teams[place].Players {
		callback(i, p)
	}
}

func (r *reader) FindPlayer(place arena.TeamPlace, playerNumber arena.PlayerNumber) (*Player, error) {
	for _, playerInfo := range r.teams[place].Players {
		if playerInfo.Number == playerNumber {
			return playerInfo, nil
		}
	}
	return nil, PlayerNotFound
}

func (r *reader) IHoldTheBall() bool {
	return r.GameInfo.Ball.Holder != nil && r.GameInfo.Ball.Holder.Id == r.player.Id
}
func (r *reader) OpponentGoal() arena.Goal {

}
func (r *reader) DefenseGoal() arena.Goal {

}
func (r *reader) IsGoalkeeper() bool {

}

var PlayerNotFound = errors.New("player not found")
