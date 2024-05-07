package order

type option func(s *Service)

func OptionStorage(st storage) option {
	return func(s *Service) {
		s.storage = st
	}
}

func OptionRoomService(rs roomService) option {
	return func(s *Service) {
		s.roomService = rs
	}
}
