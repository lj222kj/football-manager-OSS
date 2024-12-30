package team

type Team interface {
	Get() error
	GetById() error
}
