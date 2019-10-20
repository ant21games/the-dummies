package broker

import (
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/physics"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestReader_Ball(t *testing.T) {
	expectedBall := Ball{}
	expectedBall.Coords = physics.Point{PosX: 233, PosY: 300}
	expectedBall.Size = 23123

	gameState := GameInfo{
		Ball: expectedBall,
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")
	assert.Equal(t, expectedBall, reader.Ball())
}

func TestReader_Turn(t *testing.T) {
	expectedTurn := rand.Int()
	gameState := GameInfo{
		Turn: expectedTurn,
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")
	assert.Equal(t, expectedTurn, reader.Turn())
}

func TestReader_Me_WhenIAmInHomeTeam(t *testing.T) {
	expectedPlayer := Player{Id: "home-7", Number: "7"}

	myTeam := Team{
		Players: []*Player{&expectedPlayer},
	}
	gameState := GameInfo{
		HomeTeam: myTeam,
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "7")
	assert.Equal(t, &expectedPlayer, reader.Me())
}

func TestReader_Me_WhenIAmInAwayTeam(t *testing.T) {
	expectedPlayer := Player{Id: "home-7", Number: "7"}

	myTeam := Team{
		Players: []*Player{&expectedPlayer},
	}
	gameState := GameInfo{
		AwayTeam: myTeam,
	}
	reader := NewGameStateReader(gameState, arena.AwayTeam, "7")
	assert.Equal(t, &expectedPlayer, reader.Me())
}

func TestReader_GetMyTeam_WhenItIsHomeTeam(t *testing.T) {
	expectedTeam := Team{
		Name: "A cool name",
	}
	gameState := GameInfo{
		HomeTeam: expectedTeam,
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")
	assert.Equal(t, expectedTeam, reader.GetMyTeam())
}

func TestReader_GetMyTeam_WhenItIsAwayTeam(t *testing.T) {
	expectedTeam := Team{
		Name: "A cool name",
	}
	gameState := GameInfo{
		AwayTeam: expectedTeam,
	}
	reader := NewGameStateReader(gameState, arena.AwayTeam, "1")
	assert.Equal(t, expectedTeam, reader.GetMyTeam())
}

func TestReader_GetOpponentTeamWhenItIsHome(t *testing.T) {
	expectedTeam := Team{
		Name: "A cool name",
	}
	gameState := GameInfo{
		HomeTeam: expectedTeam,
	}
	reader := NewGameStateReader(gameState, arena.AwayTeam, "1")
	assert.Equal(t, expectedTeam, reader.GetOpponentTeam())
}

func TestReader_GetOpponentTeamWhenItIsAway(t *testing.T) {
	expectedTeam := Team{
		Name: "A cool name",
	}
	gameState := GameInfo{
		AwayTeam: expectedTeam,
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")
	assert.Equal(t, expectedTeam, reader.GetOpponentTeam())
}

func TestReader_ForEachPlayer_HomeTeam(t *testing.T) {
	playerExpectedA := Player{Id: "home-1"}
	playerExpectedB := Player{Id: "home-2"}
	playerExpectedC := Player{Id: "home-3"}
	gameState := GameInfo{
		HomeTeam: Team{
			Players: []*Player{&playerExpectedA, &playerExpectedB, &playerExpectedC},
		},
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")
	counter := 0
	reader.ForEachPlayer(arena.HomeTeam, func(index int, player *Player) {
		switch index {
		case 0:
			counter++
			assert.Equal(t, &playerExpectedA, player)
		case 1:
			counter++
			assert.Equal(t, &playerExpectedB, player)
		case 2:
			counter++
			assert.Equal(t, &playerExpectedC, player)
		}
	})

	reader.ForEachPlayer(arena.AwayTeam, func(index int, player *Player) {
		assert.Fail(t, "should not find player on missing teams")
	})

	assert.Equal(t, 3, counter)
}
func TestReader_ForEachPlayer_Away(t *testing.T) {
	playerExpectedA := Player{Id: "away-1"}
	playerExpectedB := Player{Id: "away-2"}
	playerExpectedC := Player{Id: "away-3"}
	gameState := GameInfo{
		AwayTeam: Team{
			Players: []*Player{&playerExpectedA, &playerExpectedB, &playerExpectedC},
		},
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")
	counter := 0
	reader.ForEachPlayer(arena.AwayTeam, func(index int, player *Player) {
		switch index {
		case 0:
			counter++
			assert.Equal(t, &playerExpectedA, player)
		case 1:
			counter++
			assert.Equal(t, &playerExpectedB, player)
		case 2:
			counter++
			assert.Equal(t, &playerExpectedC, player)
		}
	})

	reader.ForEachPlayer(arena.HomeTeam, func(index int, player *Player) {
		assert.Fail(t, "should not find player on missing teams")
	})

	assert.Equal(t, 3, counter)
}

func TestReader_FindPlayer_ThereIsThePlayer(t *testing.T) {
	playerA := Player{Number: "1"}
	playerB := Player{Number: "2"}
	playerC := Player{Number: "3"}
	playerD := Player{Number: "1"}
	playerE := Player{Number: "2"}
	playerF := Player{Number: "3"}
	gameState := GameInfo{
		HomeTeam: Team{
			Players: []*Player{&playerA, &playerB, &playerC},
		},
		AwayTeam: Team{
			Players: []*Player{&playerD, &playerE, &playerF},
		},
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")

	player, err := reader.FindPlayer(arena.HomeTeam, "3")
	assert.Nil(t, err)
	assert.Equal(t, &playerC, player)

	player, err = reader.FindPlayer(arena.AwayTeam, "2")
	assert.Nil(t, err)
	assert.Equal(t, &playerE, player)
}

func TestReader_FindPlayer_ThereIsNoSuchPlayer(t *testing.T) {
	playerA := Player{Number: "1"}
	playerB := Player{Number: "2"}
	playerC := Player{Number: "3"}
	playerD := Player{Number: "1"}
	playerE := Player{Number: "2"}
	playerF := Player{Number: "3"}
	gameState := GameInfo{
		HomeTeam: Team{
			Players: []*Player{&playerA, &playerB, &playerC},
		},
		AwayTeam: Team{
			Players: []*Player{&playerD, &playerE, &playerF},
		},
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "1")

	_, err := reader.FindPlayer(arena.HomeTeam, "4")
	assert.Equal(t, PlayerNotFound, err)

	_, err = reader.FindPlayer(arena.AwayTeam, "6")
	assert.Equal(t, PlayerNotFound, err)

	_, err = reader.FindPlayer(arena.AwayTeam, "15")
	assert.Equal(t, PlayerNotFound, err)
}

func TestReader_IHoldTheBall_IAmHoldingHomeTeam(t *testing.T) {
	myPlayer := Player{Id: "home-7", Number: "7"}

	expectedBall := Ball{}
	expectedBall.Holder = &myPlayer

	gameState := GameInfo{
		Ball: expectedBall,
		HomeTeam: Team{
			Players: []*Player{&myPlayer},
		},
	}
	reader := NewGameStateReader(gameState, arena.HomeTeam, "7")
	assert.True(t, reader.IHoldTheBall())
}

func TestReader_IHoldTheBall_IAmHoldingAwayTeam(t *testing.T) {
	myPlayer := Player{Id: "away-7", Number: "7"}

	expectedBall := Ball{}
	expectedBall.Holder = &myPlayer

	gameState := GameInfo{
		Ball: expectedBall,
		AwayTeam: Team{
			Players: []*Player{&myPlayer},
		},
	}
	reader := NewGameStateReader(gameState, arena.AwayTeam, "7")
	assert.True(t, reader.IHoldTheBall())
}

func TestReader_IHoldTheBall_IAmNotHolding(t *testing.T) {
	myPlayer := Player{Id: "away-7", Number: "7"}
	otherPlayer := Player{Id: "away-8", Number: "8"}

	expectedBall := Ball{}
	expectedBall.Holder = &otherPlayer

	gameState := GameInfo{
		Ball: expectedBall,
		AwayTeam: Team{
			Players: []*Player{&myPlayer, &otherPlayer},
		},
	}
	reader := NewGameStateReader(gameState, arena.AwayTeam, "7")
	assert.False(t, reader.IHoldTheBall())
}

func TestReader_IHoldTheBall_NobodyIsHolding(t *testing.T) {
	myPlayer := Player{Id: "away-7", Number: "7"}
	otherPlayer := Player{Id: "away-8", Number: "8"}

	expectedBall := Ball{}

	gameState := GameInfo{
		Ball: expectedBall,
		AwayTeam: Team{
			Players: []*Player{&myPlayer, &otherPlayer},
		},
	}
	reader := NewGameStateReader(gameState, arena.AwayTeam, "7")
	assert.False(t, reader.IHoldTheBall())
}

func TestReader_OpponentGoal_IAmHomeTeam(t *testing.T) {
	reader := NewGameStateReader(GameInfo{}, arena.HomeTeam, "7")
	assert.Equal(t, arena.AwayTeamGoal, reader.OpponentGoal())
}

func TestReader_OpponentGoal_IAmAwayTeam(t *testing.T) {
	reader := NewGameStateReader(GameInfo{}, arena.AwayTeam, "7")
	assert.Equal(t, arena.HomeTeamGoal, reader.OpponentGoal())
}
func TestReader_DefenseGoal_IAmHomeTeam(t *testing.T) {
	reader := NewGameStateReader(GameInfo{}, arena.HomeTeam, "7")
	assert.Equal(t, arena.HomeTeamGoal, reader.DefenseGoal())
}

func TestReader_DefenseGoal_IAmAwayTeam(t *testing.T) {
	reader := NewGameStateReader(GameInfo{}, arena.AwayTeam, "7")
	assert.Equal(t, arena.AwayTeamGoal, reader.DefenseGoal())
}

func TestReader_AmIGoalkeeper_Iam(t *testing.T) {
	reader := NewGameStateReader(GameInfo{}, arena.AwayTeam, "1")
	assert.True(t, reader.AmIGoalkeeper())
}

func TestReader_AmIGoalkeeper_IamNot(t *testing.T) {
	reader := NewGameStateReader(GameInfo{}, arena.AwayTeam, "2")
	assert.False(t, reader.AmIGoalkeeper())
}
