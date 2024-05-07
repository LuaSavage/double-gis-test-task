package order

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/LuaSavage/double-gis-test-task/internal/services/room"
	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
	"github.com/LuaSavage/double-gis-test-task/pkg/utils"
)

type Service struct {
	m sync.RWMutex

	storage     storage
	roomService roomService
}

func New(opts ...option) *Service {
	s := &Service{
		m: sync.RWMutex{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Service) AddOrder(ctx context.Context, req *Order) (id int64, err error) {
	const methodName = "service.order.AddOrder"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	s.m.Lock()
	defer s.m.Unlock()

	availabilities, err := s.roomService.GetRoomAvailabilitiesByHotelID(ctx, req.RoomID, req.HotelID)
	if err != nil {
		return id, fmt.Errorf("get room availabilities: %w", err)
	}

	validAv := findValidRoomAvailabilities(availabilities, req.From, req.To)
	if len(validAv) == 0 {
		return 0, fmt.Errorf("room is not available for selected dates")
	}

	err = s.updateRoomAvailabilities(ctx, validAv)
	if err != nil {
		return id, fmt.Errorf("update room availabilities: %w", err)
	}

	id, err = s.storage.AddOrder(ctx, req)
	if err != nil {
		return id, fmt.Errorf("add order: %w", err)
	}

	return id, nil
}

func findValidRoomAvailabilities(ra []room.RoomAvailability, from, to time.Time) []room.RoomAvailability {
	res := make([]room.RoomAvailability, 0)
	daysBetween := utils.DaysBetween(from, to)

	for _, r := range ra {
		if r.Quota > 0 && utils.IsDateBetween(r.Date, from, to, true) {
			r.Quota--

			res = append(res, r)
		}
	}

	if len(daysBetween) != len(res) {
		return nil
	}

	return res
}

func (s *Service) updateRoomAvailabilities(ctx context.Context, av []room.RoomAvailability) error {
	for _, v := range av {
		err := s.roomService.UpdateRoomAvailability(ctx, &v)
		if err != nil {
			return fmt.Errorf("update room availability: %w", err)
		}
	}

	return nil
}

func (s *Service) GetOrder(ctx context.Context, params map[string]string) (resp *Order, err error) {
	const methodName = "service.order.GetOrder"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	s.m.RLock()
	defer s.m.RUnlock()

	resp, err = s.storage.GetOrder(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}

	return resp, nil
}

func (s *Service) GetOrders(ctx context.Context) (resp []Order, err error) {
	const methodName = "service.order.AddOrder"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	s.m.RLock()
	defer s.m.RUnlock()

	resp, err = s.storage.GetOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].CreatedAt.Before(resp[j].CreatedAt)
	})

	return resp, nil
}
