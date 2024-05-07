package room

import (
	"context"
	"fmt"
	"time"

	"github.com/LuaSavage/double-gis-test-task/internal/services/room"
)

type Storage struct {
	lastId int64
	db     map[int64]room.RoomAvailability
}

func New() *Storage {
	return &Storage{
		lastId: 0,
		db:     make(map[int64]room.RoomAvailability),
	}
}

func (s *Storage) AddRoomAvailability(_ context.Context, req *room.RoomAvailability) (int64, error) {
	s.lastId++
	req.ID = s.lastId
	req.CreatedAt = time.Now().UTC()

	s.db[req.ID] = *req

	return req.ID, nil
}

func (s *Storage) GetRoomAvailabilities(_ context.Context) ([]room.RoomAvailability, error) {
	res := make([]room.RoomAvailability, 0, len(s.db))

	for _, v := range s.db {
		res = append(res, v)
	}

	return res, nil
}

func (s *Storage) UpdateRoomAvailability(_ context.Context, req *room.RoomAvailability) error {
	_, ok := s.db[req.ID]
	if !ok {
		return fmt.Errorf("room availability not found: id=%v", req.ID)
	}

	s.db[req.ID] = *req

	return nil
}
