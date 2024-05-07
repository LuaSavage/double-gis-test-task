package room

import (
	"context"
	"fmt"
	"sort"

	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
)

type Service struct {
	storage storage
}

func New(opts ...option) *Service {
	s := &Service{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Service) AddRoomAvailability(ctx context.Context, req *RoomAvailability) (id int64, err error) {
	const methodName = "service.room.AddRoomAvailability"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	id, err = s.storage.AddRoomAvailability(ctx, req)
	if err != nil {
		return id, fmt.Errorf("add room availability: %w", err)
	}

	return id, nil
}

func (s *Service) UpdateRoomAvailability(ctx context.Context, req *RoomAvailability) (err error) {
	const methodName = "service.room.UpdateRoomAvailability"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	err = s.storage.UpdateRoomAvailability(ctx, req)
	if err != nil {
		return fmt.Errorf("update room availability: %w", err)
	}

	return nil
}

func (s *Service) GetRoomAvailabilitiesByHotelID(ctx context.Context, roomID, hotelID string) (resp []RoomAvailability, err error) {
	const methodName = "service.room.GetRoomAvailabilitiesByHotelID"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	rooms, err := s.storage.GetRoomAvailabilities(ctx)
	if err != nil {
		return nil, fmt.Errorf("get room availabilities by hotel: %w", err)
	}

	for i := range rooms {
		if rooms[i].HotelID == hotelID && rooms[i].RoomID == roomID {
			resp = append(resp, rooms[i])
		}
	}

	if len(resp) == 0 {
		return nil, fmt.Errorf("no available dates")
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Date.Before(resp[j].Date)
	})

	return resp, nil
}

func (s *Service) GetAllRoomAvailabilities(ctx context.Context) (resp []RoomAvailability, err error) {
	const methodName = "service.room.GetAllRoomAvailabilities"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	resp, err = s.storage.GetRoomAvailabilities(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all room availabilities: %w", err)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].CreatedAt.Before(resp[j].CreatedAt)
	})

	return resp, nil
}
