package user

type Service interface {
	GetUser(id string) (User, error)
	DeleteUser(id string) error
	AddUser(user User) (User, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (s *serviceImpl) GetUser(id string) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) AddUser(user User) (User, error) {
	//TODO implement me
	panic("implement me")
}
