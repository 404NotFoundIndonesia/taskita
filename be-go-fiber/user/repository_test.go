package user_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	u "taskita/user"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Register(ctx context.Context, user u.User) (u.User, error) {
	r := m.Called(ctx, user)
	return r.Get(0).(u.User), r.Error(1)
}

func (m *mockRepository) Login(ctx context.Context, username string) (u.User, error) {
	r := m.Called(ctx, username)
	return r.Get(0).(u.User), r.Error(1)
}

func (m *mockRepository) User(ctx context.Context, id int) (u.User, error) {
	r := m.Called(ctx, id)
	return r.Get(0).(u.User), r.Error(1)
}
