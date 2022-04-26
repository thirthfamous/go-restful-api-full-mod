package repository

import (
	"context"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(ctx context.Context, db *gorm.DB, category *domain.Category) error
	Update(ctx context.Context, db *gorm.DB, category *domain.Category) error
	Delete(ctx context.Context, db *gorm.DB, category domain.Category) error
	FindById(ctx context.Context, db *gorm.DB, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Category, error)
}
