package usecase

import (
	"backend/internal/domain"
	"context"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *userUsecase) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return u.repo.GetAll(ctx)
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	return u.repo.Update(ctx, user)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
