package handler

import (
	"net/http"

	"github.com/LuaSavage/double-gis-test-task/internal/services/room"
	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
	"github.com/LuaSavage/double-gis-test-task/pkg/utils"
)

func (c *Handler) GetAllRoomAvailabilities(w http.ResponseWriter, r *http.Request) {
	var err error
	const methodName = "controller.GetAllRoomAvailabilities"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	ctx := r.Context()

	rooms, err := c.roomService.GetAllRoomAvailabilities(ctx)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, err)
		return
	}

	resp := newGetAllRoomAvailabilitiesResp(rooms)
	utils.HttpResponse(w, http.StatusOK, resp)
}

func newGetAllRoomAvailabilitiesResp(rooms []room.RoomAvailability) GetAllRoomAvailabilitiesResp {
	res := make([]RoomAvailability, 0)

	for i := range rooms {
		res = append(res, RoomAvailability{
			ID:        rooms[i].ID,
			HotelID:   rooms[i].HotelID,
			RoomID:    rooms[i].RoomID,
			Quota:     rooms[i].Quota,
			Date:      rooms[i].Date,
			CreatedAt: rooms[i].CreatedAt,
		})
	}

	return GetAllRoomAvailabilitiesResp{Availabilities: res}
}
