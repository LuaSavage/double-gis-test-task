package handler

type option func(c *Handler)

func OptionRoomService(rs roomService) option {
	return func(c *Handler) {
		c.roomService = rs
	}
}

func OptionOrderService(os orderService) option {
	return func(c *Handler) {
		c.orderService = os
	}
}
