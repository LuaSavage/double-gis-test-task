package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/LuaSavage/double-gis-test-task/internal/services/order"
	"github.com/LuaSavage/double-gis-test-task/internal/services/room"
)

const (
	defaultTimeout = time.Duration(30) * time.Second

	healthPath = "/health"
	orderPath  = "/orders"
	roomsPath  = "/room-availabilities"
)

var ErrInvalidRequest = fmt.Errorf("request is invalid")

type orderService interface {
	AddOrder(ctx context.Context, req *order.Order) (id int64, err error)
	GetOrders(ctx context.Context) ([]order.Order, error)
	GetOrder(ctx context.Context, params map[string]string) (resp *order.Order, err error)
}

type roomService interface {
	GetAllRoomAvailabilities(ctx context.Context) (resp []room.RoomAvailability, err error)
}

type AddOrderReq struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (r *AddOrderReq) ToService() *order.Order {
	if r == nil {
		return nil
	}

	return &order.Order{
		HotelID:   r.HotelID,
		RoomID:    r.RoomID,
		UserEmail: r.UserEmail,
		From:      r.From,
		To:        r.To,
	}
}

type AddOrderResp struct {
	ID int64 `json:"id"`
}

type GetOrdersResp struct {
	Orders []Order `json:"orders"`
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

type GetAllRoomAvailabilitiesResp struct {
	Availabilities []RoomAvailability `json:"availabilities"`
}

type RoomAvailability struct {
	ID        int64     `json:"id"`
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	Quota     int       `json:"quota"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}
