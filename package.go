package client

import (
	"github.com/makeitplay/arena/orders"
	"github.com/makeitplay/arena/physics"
	"github.com/makeitplay/arena/units"
)

// MakeMoveOrder creates a Move order
func MakeMoveOrder(origin, target physics.Point, speed float64) (orders.Order, error) {
	vec, err := physics.NewVector(origin, target)
	if err != nil {
		return orders.Order{}, err
	}
	vel := physics.NewZeroedVelocity(*vec.Normalize())
	vel.Speed = speed
	return orders.NewMoveOrder(vel), nil
}

// MakeMoveOrderMaxSpeed creates a MoveOrderMax order
func MakeMoveOrderMaxSpeed(origin, target physics.Point) (orders.Order, error) {
	return MakeMoveOrder(origin, target, units.PlayerMaxSpeed)
}

// MakeStopOrder creates a Stop order
func MakeStopOrder(direction physics.Vector) orders.Order {
	return orders.NewMoveOrder(physics.NewZeroedVelocity(direction))
}

// MakeKickOrder creates a Kick order
// The ball kick sums the current ball velocity with the kick velocity. This function first find the right vector
// will correctly change the ball direction, and than creates the order.
func MakeKickOrder(ball Ball, target physics.Point, speed float64) (orders.Order, error) {
	ballExpectedDirection, err := physics.NewVector(ball.Coords, target)
	if err != nil {
		return orders.Order{}, err
	}
	diffVector, err := ballExpectedDirection.Sub(ball.Velocity.Direction)
	if err != nil {
		return orders.Order{}, err
	}
	diffVector.Normalize()
	vec := physics.NewZeroedVelocity(*diffVector)
	vec.Speed = speed

	return orders.NewKickOrder(vec), nil
}

// MakeJumpOrder creates a Jump order
func MakeJumpOrder(origin, target physics.Point, speed float64) (orders.Order, error) {
	vec, err := physics.NewVector(origin, target)
	if err != nil {
		return orders.Order{}, err
	}
	vel := physics.NewZeroedVelocity(*vec.Normalize())
	vel.Speed = speed
	return orders.NewJumpOrder(vel), nil
}

// MakeCatchOrder creates a Catch order
func MakeCatchOrder() orders.Order {
	return orders.NewCatchOrder()
}
