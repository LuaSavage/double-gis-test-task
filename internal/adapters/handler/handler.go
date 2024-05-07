package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	orderService orderService
	roomService  roomService
}

func New(opts ...option) *Handler {
	c := &Handler{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Handler) RegisterHandlers(r *chi.Mux) {
	addDefaultMiddlewares(r)

	r.Get(healthPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route(orderPath, func(r chi.Router) {
		r.Post("/", c.AddOrder)
		r.Get("/", c.GetOrders)
	})

	r.Route(roomsPath, func(r chi.Router) {
		r.Get("/", c.GetAllRoomAvailabilities)
	})
}

func addDefaultMiddlewares(r *chi.Mux) {
	r.Use(middleware.Timeout(defaultTimeout))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
}
