package main

type IServices interface {
	GetUser(id int) (*User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUser(id int) (*User, error) {
	return s.repo.FindUserByID(id)
}