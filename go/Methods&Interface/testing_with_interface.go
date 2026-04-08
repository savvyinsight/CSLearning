package main

import (
	"fmt"
	"testing"
)

type User struct {
	ID   string
	Name string
}

type UserReposity interface {
	GetUser(id string) (*User, error)
	SaveUser(user *User) error
}

// MockUserReposity
type MockUserRepo struct {
	users map[string]*User
}

func (m *MockUserRepo) GetUser(id string) (*User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (m *MockUserRepo) SaveUser(user *User) error {
	m.users[user.ID] = user
	return nil
}

type UserService struct {
	repo UserReposity
}

func TestUserService(t *testing.T) {
	mockRepo := &MockUserRepo{
		users: map[string]*User{
			"1": &User{"1", "Bob"},
		},
	}

	service := UserService{repo: mockRepo}
	user, err := service.repo.GetUser("1")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if user.Name != "Test User" {
		t.Errorf("Expected 'Test User', got %s", user.Name)
	}
}
