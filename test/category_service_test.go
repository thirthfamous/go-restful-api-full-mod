package test

import (
	"context"
	"testing"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"
	"thirthfamous/golang-restful-api-clean-architecture/model/web"
	"thirthfamous/golang-restful-api-clean-architecture/repository"
	"thirthfamous/golang-restful-api-clean-architecture/service"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCreateServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	validate := validator.New()
	request := web.CategoryCreateRequest{
		Name: "Komputer",
	}
	response, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)
	assert.Equal(t, response.Name, request.Name)
	assert.NoError(t, err)
}

func TestCreateServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	validate := validator.New()
	request := web.CategoryCreateRequest{
		Name: "",
	}
	_, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)
	assert.Error(t, err)

}

func TestCreateDuplicateServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	request := web.CategoryCreateRequest{
		Name: "Software",
	}
	_, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)
	assert.Error(t, err)
}

func TestUpdateCategoryServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	request := web.CategoryUpdateRequest{
		Id:   category.Id,
		Name: "Hardware",
	}
	response, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Update(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, request.Name, response.Name)
}

func TestUpdateCategoryServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	request := web.CategoryUpdateRequest{
		Id:   category.Id,
		Name: "",
	}
	response, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Update(ctx, request)
	assert.Error(t, err)
	assert.Equal(t, request.Name, response.Name)
}

func TestUpdateCategoryServiceFailedIdNotFound(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	request := web.CategoryUpdateRequest{
		Id:   12345,
		Name: "Hardware",
	}
	_, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Update(ctx, request)
	assert.Error(t, err)
}

func TestDeleteCategoryServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Delete(ctx, category.Id)
	assert.NoError(t, err)
}

func TestDeleteCategoryServiceNotFound(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Delete(ctx, 12345)
	assert.Error(t, err)
}

func TestFindByIdCategoryServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	_, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).FindById(ctx, category.Id)
	assert.NoError(t, err)
}

func TestFindByIdCategoryServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)

	validate := validator.New()
	_, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).FindById(ctx, 100)
	assert.Error(t, err)
}

func TestFindAllCategoryServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	category := domain.Category{
		Name: "Software",
	}
	category2 := domain.Category{
		Name: "Hardware",
	}
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(ctx, db, &category)
	categoryRepository.Save(ctx, db, &category2)

	validate := validator.New()
	result, err := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).FindAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, result[0].Name, category.Name)
	assert.Equal(t, result[1].Name, category2.Name)
}
