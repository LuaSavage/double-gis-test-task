package handler

import (
	"net/http"

	"github.com/LuaSavage/double-gis-test-task/internal/services/order"
	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
	"github.com/LuaSavage/double-gis-test-task/pkg/utils"
)

func (c *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var err error
	const methodName = "controller.GetOrders"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	ctx := r.Context()

	orders, err := c.orderService.GetOrders(ctx)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, err)
		return
	}

	resp := newGetOrdersResp(orders)
	utils.HttpResponse(w, http.StatusOK, resp)
}

func newGetOrdersResp(orders []order.Order) GetOrdersResp {
	res := make([]Order, 0, len(orders))

	for i := range orders {
		res = append(res, Order{
			ID:        orders[i].ID,
			HotelID:   orders[i].HotelID,
			RoomID:    orders[i].RoomID,
			UserEmail: orders[i].UserEmail,
			From:      orders[i].From,
			To:        orders[i].To,
			CreatedAt: orders[i].CreatedAt,
		})
	}

	return GetOrdersResp{Orders: res}
}
