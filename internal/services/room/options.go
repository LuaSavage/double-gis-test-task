package room

type option func(s *Service)

func OptionStorage(st storage) option {
	return func(s *Service) {
		s.storage = st
	}
}
