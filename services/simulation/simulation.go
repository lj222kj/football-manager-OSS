package simulation

type service struct{}

func New() Simulation {
	return &service{}
}
