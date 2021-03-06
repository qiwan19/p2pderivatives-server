package mock_userservice

import (
	"context"
	"errors"
	"p2pderivatives-server/internal/user/usercommon"
	"p2pderivatives-server/internal/user/userservice"
	"p2pderivatives-server/test/mocks/mock_userrepository"
)

var tokens = []string{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDYzNDU4MDAsImp0aSI6IjEifQ.QfDVtZbYnXaaQ3_vBJow-s9KT5OKAIT7O3dc9hR_yoc",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDYzNDU4MDAsImp0aSI6InRlc3QxIn0._HUTrOKAtYzLLUrMzhpA7TOrkl1NEp_M5YoRDDZsDmg",
}

var tokenIndex = 0

// ServiceMock is a mock for the usercommon.ServiceIf interface
type ServiceMock struct {
	repo *mock_userrepository.RepositoryMock
}

// NewServiceMock creates a new ServiceMock instance
func NewServiceMock() *ServiceMock {
	return &ServiceMock{repo: mock_userrepository.NewRepositoryMock()}
}

// CreateUser creates a user
func (service *ServiceMock) CreateUser(
	ctx context.Context, condition *usercommon.User) (*usercommon.User, error) {
	err := service.repo.CreateUser(ctx, condition)
	return condition, err
}

// FindFirstUser finds a user
func (service *ServiceMock) FindFirstUser(
	ctx context.Context, condition *usercommon.User, orders []string) (*usercommon.User, error) {
	return service.repo.FindFirstUser(ctx, condition, orders)
}

// FindFirstUserByName finds the user based on account.
func (service *ServiceMock) FindFirstUserByName(
	ctx context.Context, name string) (*usercommon.User, error) {
	return service.repo.FindFirstUserByName(ctx, &usercommon.User{Name: name})
}

// FindUsers find users.
func (service *ServiceMock) FindUsers(
	ctx context.Context,
	condition *usercommon.User,
	offset int,
	limit int,
	orders []string) ([]usercommon.User, error) {
	return service.repo.FindUsers(ctx, condition, offset, limit, orders)
}

// GetAllUsers returns all users.
func (service *ServiceMock) GetAllUsers(ctx context.Context) ([]usercommon.User, error) {
	return service.repo.GetAllUsers(ctx)
}

// UpdateUser updates a usercommon.
func (service *ServiceMock) UpdateUser(
	ctx context.Context, condition *usercommon.User) (*usercommon.User, error) {
	panic("Not implemented")
}

// DeleteUser deletes a usercommon.
func (service *ServiceMock) DeleteUser(
	ctx context.Context, condition *usercommon.User) error {
	return service.repo.DeleteUser(ctx, condition)
}

// AuthenticateUser authenticates a usercommon.
func (service *ServiceMock) AuthenticateUser(
	ctx context.Context, name, password string) (*usercommon.User, *usercommon.TokenInfo, error) {
	model, err := service.FindFirstUserByName(ctx, name)
	if err != nil || model.Password != password {
		return nil, nil, errors.New("Bad login")
	}

	return model, &usercommon.TokenInfo{AccessToken: "", RefreshToken: "", ExpiresIn: 10}, nil
}

// FindUserByCondition finds a user by condition.
func (service *ServiceMock) FindUserByCondition(
	ctx context.Context, condition *usercommon.Condition) ([]usercommon.User, error) {
	panic("Not implemented")
}

// ChangeUserPassword changes a user password.
func (service *ServiceMock) ChangeUserPassword(
	ctx context.Context, userID, newPassword, oldPassword string) (*usercommon.User, error) {
	userModel, err := service.FindFirstUser(ctx, &usercommon.User{ID: userID}, nil)
	if err != nil {
		return nil, err
	}

	if userModel.Password != oldPassword || !userservice.VerifyNewPassword(newPassword) {
		return nil, errors.New("Invalid password")
	}

	userModel.Password = newPassword

	err = service.repo.UpdateUser(ctx, userModel)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

// RefreshUserToken refreshes a user token.
func (service *ServiceMock) RefreshUserToken(
	ctx context.Context, refreshToken string) (*usercommon.TokenInfo, error) {
	panic("Not implemented")
}

// RevokeRefreshToken revokes a refresh token.
func (service *ServiceMock) RevokeRefreshToken(ctx context.Context, refreshToken string) error {
	panic("Not implemented")
}
