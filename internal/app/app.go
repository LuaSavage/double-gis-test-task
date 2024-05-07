package app

import (
	"context"
	"net/http"
	"time"

	"github.com/LuaSavage/double-gis-test-task/pkg/utils"

	"github.com/go-chi/chi/v5"

	"github.com/LuaSavage/double-gis-test-task/pkg/logger"

	roomService "github.com/LuaSavage/double-gis-test-task/internal/services/room"

	orderService "github.com/LuaSavage/double-gis-test-task/internal/services/order"

	"github.com/LuaSavage/double-gis-test-task/internal/config"

	roomStore "github.com/LuaSavage/double-gis-test-task/internal/adapters/storage/room"

	orderStore "github.com/LuaSavage/double-gis-test-task/internal/adapters/storage/order"

	"github.com/LuaSavage/double-gis-test-task/internal/adapters/handler"
)

const defaultTimeout = time.Duration(30) * time.Second

type App struct {
	server *http.Server
}

func New(cfg *config.Config) (*App, error) {
	// stores
	roomStore := roomStore.New()
	orderStore := orderStore.New()

	// migrations
	migrateRoom(roomStore)

	// services
	roomService := roomService.New(roomService.OptionStorage(roomStore))

	orderService := orderService.New(
		orderService.OptionStorage(orderStore),
		orderService.OptionRoomService(roomService))

	// controller
	ctrl := handler.New(
		handler.OptionOrderService(orderService),
		handler.OptionRoomService(roomService))

	// server
	server := &http.Server{
		Addr:         ":" + cfg.PublicPort,
		ReadTimeout:  defaultTimeout,
		WriteTimeout: defaultTimeout,
		TLSNextProto: nil,
	}

	r := chi.NewRouter()
	ctrl.RegisterHandlers(r)
	server.Handler = r

	return &App{
		server: server,
	}, nil
}

func (a *App) Run() error {
	logger.Infof("server listening on localhost:8080")
	err := a.server.ListenAndServe()
	if err != nil {
		logger.Errorf("server failed: %s", err)

		return err
	}

	return nil
}

func (a *App) Stop() error {
	return a.server.Shutdown(context.Background())
}

func migrateRoom(storage *roomStore.Storage) {
	availableRooms := []roomService.RoomAvailability{
		{
			HotelID: "1",
			RoomID:  "1",
			Quota:   1,
			Date:    utils.Date(1, 1, 2024),
		},
		{
			HotelID: "1",
			RoomID:  "2",
			Quota:   1,
			Date:    utils.Date(1, 1, 2024),
		},
		{
			HotelID: "1",
			RoomID:  "2",
			Quota:   2,
			Date:    utils.Date(2, 1, 2024),
		},
		{
			HotelID: "1",
			RoomID:  "2",
			Quota:   3,
			Date:    utils.Date(3, 1, 2024),
		},
		{
			HotelID: "1",
			RoomID:  "2",
			Quota:   4,
			Date:    utils.Date(4, 1, 2024),
		},
		{
			HotelID: "1",
			RoomID:  "2",
			Quota:   0,
			Date:    utils.Date(5, 1, 2024),
		},
		{
			HotelID: "2",
			RoomID:  "1",
			Quota:   3,
			Date:    utils.Date(1, 1, 2024),
		},
	}

	for _, v := range availableRooms {
		_, _ = storage.AddRoomAvailability(context.Background(), &v)
	}
}
