package repository

import (
	"context"
	"errors"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, db *gorm.DB, category *domain.Category) error {
	result := db.WithContext(ctx).Create(&category)
	return result.Error
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, db *gorm.DB, category *domain.Category) error {
	return db.WithContext(ctx).Model(&category).Update("Name", category.Name).Error
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, category domain.Category) error {
	return db.WithContext(ctx).Delete(&category, category.Id).Error
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, categoryId int) (domain.Category, error) {
	category := domain.Category{}
	result := db.WithContext(ctx).First(&category, categoryId)
	if result.RowsAffected == 0 {
		return category, errors.New("NOT FOUND")
	}

	return category, result.Error
}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Category, error) {
	categories := []domain.Category{}
	result := db.WithContext(ctx).Order("id").Find(&categories)

	if result.RowsAffected == 0 {
		return categories, errors.New("NOT FOUND")
	}

	return categories, result.Error
}
