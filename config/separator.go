package config

type separator struct{}

func (s *separator) Run() error {
	return nil
}

func (s *separator) Info() info {
	return info{}
}
