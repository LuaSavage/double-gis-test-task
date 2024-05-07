package handler

import (
	"net/http"

	"github.com/LuaSavage/double-gis-test-task/internal/services/order"
	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
	"github.com/LuaSavage/double-gis-test-task/pkg/utils"
)

func (c *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var err error
	const methodName = "controller.GetOrder"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	ctx := r.Context()

	order, err := c.orderService.GetOrder(ctx, nil)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, err)
		return
	}

	resp := newGetOrderResp(order)
	utils.HttpResponse(w, http.StatusOK, resp)
}

func newGetOrderResp(order *order.Order) Order {
	res := Order{
		ID:        order.ID,
		HotelID:   order.HotelID,
		RoomID:    order.RoomID,
		UserEmail: order.UserEmail,
		From:      order.From,
		To:        order.To,
	}

	return res
}
