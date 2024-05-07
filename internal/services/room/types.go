package room

import (
	"context"
	"time"
)

type storage interface {
	AddRoomAvailability(ctx context.Context, req *RoomAvailability) (id int64, err error)
	UpdateRoomAvailability(ctx context.Context, req *RoomAvailability) error
	GetRoomAvailabilities(ctx context.Context) ([]RoomAvailability, error)
}

type RoomAvailability struct {
	ID        int64     `json:"id"`
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	Quota     int       `json:"quota"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}
