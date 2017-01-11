package user

import "github.com/stretchr/testify/mock"

//ManagerMock is a mockup of the user manager service
type ManagerMock struct {
	mock.Mock
}

//Login is a mocked method
func (m *ManagerMock) Login(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

//Create is a mocked method
func (m *ManagerMock) Create(u *User) (*User, error) {
	args := m.Called(u)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}
