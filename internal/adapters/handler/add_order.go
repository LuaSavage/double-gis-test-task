package handler

import (
	"net/http"

	"github.com/LuaSavage/double-gis-test-task/pkg/logger"
	"github.com/LuaSavage/double-gis-test-task/pkg/utils"
)

func (c *Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var err error
	const methodName = "controller.AddOrder"
	logger.Infof("%s started", methodName)
	defer func() {
		if err != nil {
			logger.Errorf("%s failed: %v", methodName, err)
		} else {
			logger.Infof("%s success", methodName)
		}
	}()

	ctx := r.Context()
	req := AddOrderReq{}

	err = utils.ReadHttpBody(r, &req)
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, err)
		return
	}

	err = validateAddOrderReq(&req)
	if err != nil {
		utils.HttpError(w, http.StatusBadRequest, err)
		return
	}

	id, err := c.orderService.AddOrder(ctx, req.ToService())
	if err != nil {
		utils.HttpError(w, http.StatusInternalServerError, err)
		return
	}

	resp := AddOrderResp{ID: id}
	utils.HttpResponse(w, http.StatusOK, resp)
}

func validateAddOrderReq(req *AddOrderReq) error {
	if req.HotelID == "" || req.RoomID == "" || req.UserEmail == "" || req.From.IsZero() || req.To.IsZero() {
		return ErrInvalidRequest
	}

	return nil
}
