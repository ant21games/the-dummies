package broker

import (
	"context"
	"fmt"
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/orders"
	"github.com/makeitplay/arena/physics"
	"github.com/makeitplay/arena/talk"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/url"
	"time"
)

func NewBrocker(logger Logger) Broker {
	return &gameBroker{
		logger: logger,
		turnCtxFactory: func(msg GameMessage, place arena.TeamPlace, number string) (turnContext TurnContext, cancelFunc context.CancelFunc) {
			//note that the turn ctx is not a sub ctx of the game context. We must watch if both are not done before trying send orders
			ctx, turnBreaker := context.WithCancel(context.Background())
			return &turnCtx{
				ctx:     ctx,
				logger:  logger,
				reader:  NewGameStateReader(msg.GameInfo, place, number),
				gameMsg: &msg,
			}, turnBreaker
		},
	}
}

type gameBroker struct {
	mainCtx        context.Context
	gameCtxCancel  context.CancelFunc
	logger         Logger
	config         *Configuration
	turnCtxFactory TurnContextFactory
	Talker         talk.Talker
}

func (b *gameBroker) Dial(configuration *Configuration) (context.Context, error) {
	b.config = configuration
	b.mainCtx, b.gameCtxCancel = context.WithCancel(context.Background())

	talkerCtx, talker, err := TalkerSetup(b.mainCtx, configuration, physics.Point{})
	if err != nil {
		return nil, err
	}

	b.Talker = talker
	go listenServerMessages(p)

	go func() {
		select {
		case <-talkerCtx.Done(): //if the talker stops, we must stop the broker
			b.logger.Warnf("was connection lost: %s", talkerCtx.Err())
			b.Stop()
		case <-b.mainCtx.Done(): //if the broker stops, we must stop the talker
			talker.Close()
			b.logger.Warnf("player stopped: %s", b.mainCtx.Err())

		}
	}()
	return b.mainCtx, nil
}

func (b *gameBroker) Stop() error {
	b.gameCtxCancel()
	return nil
}

func (b *gameBroker) SendOrders(message string, ordersList ...orders.Order) error {
	panic("implement me")
}

func (b *gameBroker) SetTurnContextFactory(turnContextFactory TurnContextFactory) {
	b.turnCtxFactory = turnContextFactory
}

func (b *gameBroker) OnMessage(func(msg GameMessage)) {
	panic("implement me")
}

func (b *gameBroker) OnListeningState(func(turnTx TurnContext, breaker context.CancelFunc)) {
	panic("implement me")
}

func TalkerSetup(mainCtx context.Context, config *Configuration, initialPos physics.Point) (context.Context, talk.Talker, error) {
	rand.Seed(time.Now().UnixNano())
	// First we have to get the command line arguments to identify this bot in its game
	uri := new(url.URL)
	uri.Scheme = "ws"
	uri.Host = fmt.Sprintf("%s:%s", config.WSHost, config.WSPort)
	uri.Path = fmt.Sprintf("/announcements/%s/%s", config.UUID, config.TeamPlace)

	playerSpec := arena.PlayerSpecifications{
		Number:          config.PlayerNumber,
		InitialCoords:   initialPos,
		Token:           config.Token,
		ProtocolVersion: "1.0",
	}

	ignoreLogs := logrus.New()
	ignoreLogs.Level = logrus.PanicLevel
	talker := talk.NewTalker(ignoreLogs.WithField("internal", "websocket"))
	talkerCtx, err := talker.Connect(mainCtx, *uri, playerSpec)
	if err != nil {
		return nil, nil, fmt.Errorf("fail on opening the websocket connection: %s", err)
	}
	return talkerCtx, talker, nil
}
