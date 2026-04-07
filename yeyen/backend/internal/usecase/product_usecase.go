package usecase

import (
	"backend/internal/domain"
	"context"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) CreateProduct(ctx context.Context, product *domain.Product) error {
	return u.repo.Create(ctx, product)
}

func (u *productUsecase) GetProduct(ctx context.Context, id int) (*domain.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *productUsecase) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	return u.repo.GetAll(ctx)
}

func (u *productUsecase) UpdateProduct(ctx context.Context, product *domain.Product) error {
	return u.repo.Update(ctx, product)
}

func (u *productUsecase) DeleteProduct(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
