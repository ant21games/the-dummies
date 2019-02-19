package client

import (
	"flag"
	"fmt"
	"github.com/makeitplay/arena/BasicTypes"
	"github.com/makeitplay/arena/units"
	"log"
	"strconv"
)

// Configuration is the set of values expected as a initial configuration of the player
type Configuration struct {
	// TeamPlace must be "home" or "away" and identifies the side of the field that the team is going to play
	TeamPlace units.TeamPlace
	// PlayerNumber must be a number between 1-11 that identifies this player in his team
	PlayerNumber BasicTypes.PlayerNumber
	// UUID is the match UUID. It will be always local for local games
	UUID string
	// WSHost is the hostname of the game server (only HTTP for now)
	WSHost string
	// WSPort is the port used by the game server
	WSPort string
	// Token is passed to the game server to bind the player to specific process, and avoid cheating
	Token string
}

// ParseFromFlags sets the flag that will allows us to change the default config
func (c *Configuration) ParseFromFlags() {
	//mandatory
	var name string
	var number int

	flag.StringVar(&name, "team", "home", "Team (home or away)")
	flag.IntVar(&number, "number", 0, "Player's number")
	flag.StringVar(&c.UUID, "uui", "local", "UUID for this player instance. (Auto-provided in production)")

	flag.StringVar(&c.WSHost, "wshost", "localhost", "Game server's websocket endpoint")
	flag.StringVar(&c.WSPort, "wsport", "8080", "Port for the websocket endpoint")

	flag.StringVar(&c.Token, "token", "", "Token used by the server to identify the right connection")

	flag.Parse()

	if name != string(units.HomeTeam) && name != string(units.AwayTeam) {
		log.Fatal("Invalid team option {" + name + "}. Must be either home or away")
	}
	if number < 1 || number > 11 {
		log.Fatal(fmt.Errorf("invalid player number {%d}. Must be 1 to 11", number))
	}
	c.PlayerNumber = BasicTypes.PlayerNumber(strconv.Itoa(number))
	c.TeamPlace = units.TeamPlace(name)
}
