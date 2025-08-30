package usecase

import (
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
)

type CreateCategoryInput struct {
	NamaCategory string `json:"nama_category"`
}

type CategoryUsecase interface {
	CreateCategory(input CreateCategoryInput) (domain.Category, error)
	GetAllCategories() ([]domain.Category, error)
	GetCategoryByID(id uint) (domain.Category, error)
	UpdateCategory(id uint, input CreateCategoryInput) (domain.Category, error)
	DeleteCategory(id uint) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo}
}

func (uc *categoryUsecase) CreateCategory(input CreateCategoryInput) (domain.Category, error) {
	category := domain.Category{
		NamaCategory: input.NamaCategory,
	}
	return uc.categoryRepo.Create(category)
}

func (uc *categoryUsecase) GetAllCategories() ([]domain.Category, error) {
	return uc.categoryRepo.FindAll()
}

func (uc *categoryUsecase) GetCategoryByID(id uint) (domain.Category, error) {
	return uc.categoryRepo.FindByID(id)
}

func (uc *categoryUsecase) UpdateCategory(id uint, input CreateCategoryInput) (domain.Category, error) {
	category, err := uc.categoryRepo.FindByID(id)
	if err != nil {
		return domain.Category{}, err
	}
	category.NamaCategory = input.NamaCategory
	return uc.categoryRepo.Update(category)
}

func (uc *categoryUsecase) DeleteCategory(id uint) error {
	category, err := uc.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.categoryRepo.Delete(category)
}