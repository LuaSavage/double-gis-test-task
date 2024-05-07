package order

import (
	"context"
	"time"

	"github.com/LuaSavage/double-gis-test-task/internal/services/room"
)

type storage interface {
	AddOrder(ctx context.Context, req *Order) (id int64, err error)
	GetOrders(ctx context.Context) ([]Order, error)
	GetOrder(ctx context.Context, params map[string]string) (*Order, error)
}

type roomService interface {
	AddRoomAvailability(ctx context.Context, req *room.RoomAvailability) (int64, error)
	UpdateRoomAvailability(ctx context.Context, req *room.RoomAvailability) error
	GetAllRoomAvailabilities(ctx context.Context) ([]room.RoomAvailability, error)
	GetRoomAvailabilitiesByHotelID(ctx context.Context, roomID, hotelID string) ([]room.RoomAvailability, error)
}

type Order struct {
	ID        int64     `json:"id"`
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	CreatedAt time.Time `json:"created_at"`
}
