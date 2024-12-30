package team

type service struct{}

func New() Team {
	return &service{}
}

func (s service) Get() error {
	return nil
}

func (s service) GetById() error {
	return nil
}
