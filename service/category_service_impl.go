package service

import (
	"context"
	"errors"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"
	"thirthfamous/golang-restful-api-clean-architecture/model/web"
	"thirthfamous/golang-restful-api-clean-architecture/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *gorm.DB, validated *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validated,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	db := service.DB
	if err != nil {
		return web.CategoryResponse{}, err
	}

	category := domain.Category{
		Name: request.Name,
	}

	err = service.CategoryRepository.Save(ctx, db, &category)

	return helper.ToCategoryResponse(category), err
}
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	db := service.DB
	if err != nil {
		return web.CategoryResponse{}, err
	}

	_, err = repository.NewCategoryRepository().FindById(ctx, db, request.Id)

	if err != nil {
		return web.CategoryResponse{}, err
	}

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	err = service.CategoryRepository.Update(ctx, db, &category)

	return helper.ToCategoryResponse(category), err

}
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) error {

	db := service.DB
	if categoryId == 0 {
		return errors.New("category id is required for this request")
	}
	_, exist := service.CategoryRepository.FindById(ctx, db, categoryId)
	if exist != nil {
		return exist
	}
	err := service.CategoryRepository.Delete(ctx, db, domain.Category{Id: categoryId})

	return err
}
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error) {
	db := service.DB
	if categoryId == 0 {
		return web.CategoryResponse{}, errors.New("category id is required for this request")
	}
	category, exist := service.CategoryRepository.FindById(ctx, db, categoryId)
	return web.CategoryResponse{Id: category.Id, Name: category.Name}, exist

}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) ([]web.CategoryResponse, error) {
	db := service.DB
	category, err := service.CategoryRepository.FindAll(ctx, db)
	return helper.ToCategoryResponses(category), err
}
