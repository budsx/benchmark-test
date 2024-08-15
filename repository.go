package main

type User struct {
	ID   int
	Name string
}

type Repository interface {
	FindUserByID(id int) (*User, error)
}

type DataRepo struct {
	data map[int]*User
}

func NewRepository() Repository {
	return &DataRepo{
		data: map[int]*User{
			1: {ID: 1, Name: "Budi"},
		},
	}
}

func (repo *DataRepo) FindUserByID(id int) (*User, error) {
	user, ok := repo.data[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}
