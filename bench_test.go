package main

import "testing"

func BenchmarkNewRepository(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewRepository()
	}
}

func BenchmarkFindUserByID(b *testing.B) { 
	repo := NewRepository()
	for i := 0; i < b.N; i++ {
		repo.FindUserByID(1)
	}
}

func BenchmarkNewService(b *testing.B) {
	repo := NewRepository()
	for i := 0; i < b.N; i++ {
		_ = NewService(repo)
	}
}

func BenchmarkService_GetUser(b *testing.B) {
	repo := NewRepository()
	service := NewService(repo)

	for i := 0; i < b.N; i++ {
		service.GetUser(1)
	}
}
