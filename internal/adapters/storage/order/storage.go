package order

import (
	"context"
	"sync"
	"time"

	"github.com/LuaSavage/double-gis-test-task/internal/services/order"
)

type Storage struct {
	m sync.RWMutex

	lastId        int64
	db            map[int64]order.Order
	indexMapEmail map[string]int64 // email -> orderID
}

func New() *Storage {
	return &Storage{
		m: sync.RWMutex{},

		lastId:        0,
		db:            make(map[int64]order.Order),
		indexMapEmail: make(map[string]int64),
	}
}

func (s *Storage) AddOrder(_ context.Context, req *order.Order) (int64, error) {
	s.lastId++
	req.ID = s.lastId
	req.CreatedAt = time.Now().UTC()

	s.db[req.ID] = *req

	s.indexMapEmail[req.UserEmail] = req.ID

	return req.ID, nil
}

func (s *Storage) GetOrders(_ context.Context) ([]order.Order, error) {
	res := make([]order.Order, 0, len(s.db))

	s.m.RLock()
	defer s.m.RUnlock()

	for _, v := range s.db {
		res = append(res, v)
	}

	return res, nil
}

func (s *Storage) GetOrder(_ context.Context, params map[string]string) (*order.Order, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	email, ok := params["email"]
	if ok {
		id, found := s.indexMapEmail[email]
		if found {
			return &order.Order{
				ID:        s.db[id].ID,
				HotelID:   s.db[id].HotelID,
				RoomID:    s.db[id].RoomID,
				UserEmail: s.db[id].UserEmail,
				From:      s.db[id].From,
				To:        s.db[id].To,
				CreatedAt: s.db[id].CreatedAt,
			}, nil
		}
	}

	return &order.Order{}, nil
}
