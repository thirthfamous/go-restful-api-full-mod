package test

import (
	"context"
	"testing"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"
	"thirthfamous/golang-restful-api-clean-architecture/repository"

	"github.com/stretchr/testify/assert"
)

func TestConnectionSuccess(t *testing.T) {
	helper.SetupTestDB()
}

func TestSaveCategoryRepositorySuccess(t *testing.T) {
	db := helper.SetupTestDB()
	ctx := context.Background()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	err := categoryRepository.Save(ctx, db, &category)
	assert.NoError(t, err)
	assert.Equal(t, category.Id, 1)
}

func TestFindByIdCategoryRepositoryNotFound(t *testing.T) {
	db := helper.SetupTestDB()
	ctx := context.Background()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	_, err := categoryRepository.FindById(ctx, db, 100)
	assert.Error(t, err)
}

func TestUpdateCategoryRepositorySuccess(t *testing.T) {
	db := helper.SetupTestDB()
	ctx := context.Background()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)
	category.Name = "Hardware"
	err := categoryRepository.Update(ctx, db, &category)
	assert.NoError(t, err)
}

func TestDeleteCategoryRepositorySuccess(t *testing.T) {
	db := helper.SetupTestDB()
	ctx := context.Background()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	err := categoryRepository.Delete(ctx, db, category)
	assert.NoError(t, err)
}

func TestFindAllCategoryRepositorySuccess(t *testing.T) {
	db := helper.SetupTestDB()
	ctx := context.Background()
	helper.TruncateCategory(db)
	category1 := domain.Category{
		Name: "Software",
	}

	category2 := domain.Category{
		Name: "Hardware",
	}

	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category1)
	categoryRepository.Save(ctx, db, &category2)

	result, err := categoryRepository.FindAll(ctx, db)
	assert.NoError(t, err)
	assert.Equal(t, result[0].Name, category1.Name)
	assert.Equal(t, result[1].Name, category2.Name)
}
