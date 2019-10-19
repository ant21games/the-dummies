package client

import (
	"github.com/makeitplay/arena/orders"
	"github.com/makeitplay/arena/physics"
	"github.com/makeitplay/arena/units"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeMoveOrder_CreatedAOrder(t *testing.T) {
	pointA := physics.Point{}
	pointB := physics.Point{PosX: 4, PosY: 3}
	expectedSpeed := 10.2
	order, err := MakeMoveOrder(pointA, pointB, expectedSpeed)
	assert.Nil(t, err)
	assert.Equal(t, orders.MOVE, order.Type)

	moveOrder, ok := order.Data.(orders.MoveOrderData)
	assert.True(t, ok)
	assert.Equal(t, expectedSpeed, moveOrder.Velocity.Speed)
	assert.Equal(t, 80.0, moveOrder.Velocity.Direction.GetX())
	assert.Equal(t, 60.0, moveOrder.Velocity.Direction.GetY())
}

func TestMakeMoveOrder_ReturnsErrorOnInvalidArgs(t *testing.T) {
	expectedSpeed := 10.2
	expectedError := "vector can not have zero length"

	_, err := MakeMoveOrder(physics.Point{}, physics.Point{}, expectedSpeed)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())

	_, err = MakeMoveOrder(physics.Point{PosX: 10}, physics.Point{PosX: 10}, expectedSpeed)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())

	_, err = MakeMoveOrder(physics.Point{PosY: 10}, physics.Point{PosY: 10}, expectedSpeed)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())

	_, err = MakeMoveOrder(physics.Point{PosX: -5, PosY: 10}, physics.Point{PosX: -5, PosY: 10}, expectedSpeed)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())
}

func TestMakeMoveOrder_CreatedAOrderNegativeValues(t *testing.T) {
	expectedSpeed := 10.2
	order, err := MakeMoveOrder(physics.Point{PosX: 4, PosY: 3}, physics.Point{}, expectedSpeed)
	assert.Nil(t, err)
	assert.Equal(t, orders.MOVE, order.Type)

	moveOrder, ok := order.Data.(orders.MoveOrderData)
	assert.True(t, ok)
	assert.Equal(t, expectedSpeed, moveOrder.Velocity.Speed)
	assert.Equal(t, -80.0, moveOrder.Velocity.Direction.GetX())
	assert.Equal(t, -60.0, moveOrder.Velocity.Direction.GetY())

	order, err = MakeMoveOrder(physics.Point{PosY: 3}, physics.Point{PosX: 4}, expectedSpeed)
	assert.Nil(t, err)
	assert.Equal(t, orders.MOVE, order.Type)

	moveOrder, ok = order.Data.(orders.MoveOrderData)
	assert.True(t, ok)
	assert.Equal(t, expectedSpeed, moveOrder.Velocity.Speed)
	assert.Equal(t, 80.0, moveOrder.Velocity.Direction.GetX())
	assert.Equal(t, -60.0, moveOrder.Velocity.Direction.GetY())

	order, err = MakeMoveOrder(physics.Point{PosX: 8, PosY: 3}, physics.Point{PosX: 4}, expectedSpeed)
	assert.Nil(t, err)
	assert.Equal(t, orders.MOVE, order.Type)

	moveOrder, ok = order.Data.(orders.MoveOrderData)
	assert.True(t, ok)
	assert.Equal(t, expectedSpeed, moveOrder.Velocity.Speed)
	assert.Equal(t, -80.0, moveOrder.Velocity.Direction.GetX())
	assert.Equal(t, -60.0, moveOrder.Velocity.Direction.GetY())
}

func TestMakeMoveOrderMaxSpeed_DefinesRightSpeed(t *testing.T) {
	pointA := physics.Point{}
	pointB := physics.Point{PosX: 4, PosY: 3}

	order, err := MakeMoveOrderMaxSpeed(pointA, pointB)
	assert.Nil(t, err)
	assert.Equal(t, orders.MOVE, order.Type)

	moveOrder, ok := order.Data.(orders.MoveOrderData)
	assert.True(t, ok)
	assert.Equal(t, units.PlayerMaxSpeed, moveOrder.Velocity.Speed)
	assert.Equal(t, 80.0, moveOrder.Velocity.Direction.GetX())
	assert.Equal(t, 60.0, moveOrder.Velocity.Direction.GetY())
}

func TestMakeKickOrder_SameDirection(t *testing.T) {
	ball := Ball{}
	ball.Coords = physics.Point{}
	ball.Velocity = physics.NewZeroedVelocity(physics.East)
	ball.Velocity.Speed = units.BallMaxSpeed

	target := physics.Point{PosX: 100} //Eastern from the current ball position
	order, err := MakeKickOrder(ball, target, units.BallMaxSpeed/2)
	assert.Nil(t, err)
	assert.Equal(t, orders.KICK, order.Type)

	kickOrderData, ok := order.Data.(orders.KickOrderData)
	assert.True(t, ok)
	assert.Equal(t, units.BallMaxSpeed/2, kickOrderData.Velocity.Speed)
	assert.Equal(t, 100.0, kickOrderData.Velocity.Direction.GetX())
	assert.Equal(t, 0.0, kickOrderData.Velocity.Direction.GetY())
}

func TestMakeKickOrder_Diagonal(t *testing.T) {
	ball := Ball{}
	ball.Coords = physics.Point{}
	ball.Velocity = physics.NewZeroedVelocity(physics.East)
	ball.Velocity.Speed = units.BallMaxSpeed

	target := physics.Point{PosX: 5, PosY: 3}
	order, err := MakeKickOrder(ball, target, units.BallMaxSpeed/2)
	assert.Nil(t, err)
	assert.Equal(t, orders.KICK, order.Type)

	kickOrderData, ok := order.Data.(orders.KickOrderData)
	assert.True(t, ok)
	assert.Equal(t, units.BallMaxSpeed/2, kickOrderData.Velocity.Speed)
	assert.Equal(t, 80.0, kickOrderData.Velocity.Direction.GetX())
	assert.Equal(t, 60.0, kickOrderData.Velocity.Direction.GetY())
}

func TestMakeKickOrder_DiagonalBackwards(t *testing.T) {
	ball := Ball{}
	ball.Coords = physics.Point{}
	ball.Velocity = physics.NewZeroedVelocity(physics.East)
	ball.Velocity.Speed = units.BallMaxSpeed

	target := physics.Point{PosX: -3, PosY: -3}
	order, err := MakeKickOrder(ball, target, units.BallMaxSpeed/2)
	assert.Nil(t, err)
	assert.Equal(t, orders.KICK, order.Type)

	kickOrderData, ok := order.Data.(orders.KickOrderData)
	assert.True(t, ok)
	assert.Equal(t, units.BallMaxSpeed/2, kickOrderData.Velocity.Speed)
	assert.Equal(t, -80.0, kickOrderData.Velocity.Direction.GetX())
	assert.Equal(t, -60.0, kickOrderData.Velocity.Direction.GetY())
}

func TestMakeStopOrder(t *testing.T) {
	order := MakeStopOrder(physics.North)

	assert.Equal(t, orders.MOVE, order.Type)

	moveOrder, ok := order.Data.(orders.MoveOrderData)
	assert.True(t, ok)
	assert.Equal(t, 0.0, moveOrder.Velocity.Speed)
	assert.Equal(t, 0.0, moveOrder.Velocity.Direction.GetX())
	assert.Equal(t, 1.0, moveOrder.Velocity.Direction.GetY())
}

func TestMakeJumpOrder(t *testing.T) {
	order, err := MakeJumpOrder(physics.Point{PosX: 10, PosY: 10}, physics.Point{PosX: 14, PosY: 13}, units.GoalKeeperJumpSpeed)
	assert.Nil(t, err)
	assert.Equal(t, orders.JUMP, order.Type)

	jumpOrder, ok := order.Data.(orders.JumpOrderData)
	assert.True(t, ok)

	assert.Equal(t, units.GoalKeeperJumpSpeed, jumpOrder.Velocity.Speed)
	assert.Equal(t, 80.0, jumpOrder.Velocity.Direction.GetX())
	assert.Equal(t, 60.0, jumpOrder.Velocity.Direction.GetY())
}

func TestMakeCatchOrder(t *testing.T) {
	order := MakeCatchOrder()
	assert.Equal(t, orders.CATCH, order.Type)

	catchOrder := order.Data
	assert.Nil(t, catchOrder)
}
