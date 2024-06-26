package service

import (
	"api/the_deck/Models/dto"
	"api/the_deck/Models/entity"
	"api/the_deck/Models/response"
	repository "api/the_deck/Repository"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	CategoryCreate(ctx *fiber.Ctx, input dto.RequestCategoryCreate) (*response.CategoryResponse, error)
	CategoryGetAll() ([]response.CategoryResponse, error)
	CategoryGetById(id uint) (*response.CategoryResponse, error)
	CategoryUpdate(ctx *fiber.Ctx, id uint, input dto.RequestCategoryUpdate) (*response.CategoryResponse, error)
	CategoryDelete(id uint) error
}

type categoryService struct {
	categoryService repository.CategoryRepository
}

func (c *categoryService) CategoryDelete(id uint) error {
	return c.categoryService.Delete(id)
}

func (c *categoryService) CategoryUpdate(ctx *fiber.Ctx, id uint, input dto.RequestCategoryUpdate) (*response.CategoryResponse, error) {
	category, err := c.categoryService.GetById(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		category.Name = input.Name
	}

	if input.Description != "" {
		category.Description = input.Description
	}

	updateCat, err := c.categoryService.Update(category)
	if err != nil {
		return nil, err
	}

	respon := &response.CategoryResponse{
		Id:          updateCat.Id,
		Name:        updateCat.Name,
		Description: updateCat.Description,
		AdminID:     updateCat.AdminID,
	}

	return respon, nil
}

func (c *categoryService) CategoryGetById(id uint) (*response.CategoryResponse, error) {
	category, err := c.categoryService.GetById(id)
	if err != nil {
		return nil, err
	}

	cat := &response.CategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
		AdminID:     category.AdminID,
	}

	return cat, nil
}

func (c *categoryService) CategoryGetAll() ([]response.CategoryResponse, error) {
	categories, err := c.categoryService.GetAll()
	if err != nil {
		return nil, err
	}

	var respon []response.CategoryResponse
	for _, category := range categories {
		respon = append(respon, response.CategoryResponse{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
			AdminID:     category.AdminID,
		})
	}

	return respon, nil
}

func (c *categoryService) CategoryCreate(ctx *fiber.Ctx, input dto.RequestCategoryCreate) (*response.CategoryResponse, error) {
	admin := ctx.Locals("admin").(entity.Admin)
	category := entity.Category{
		Name:        input.Name,
		Description: input.Description,
		AdminID:     admin.Id,
		Admin:       admin,
	}

	categoryCreate, err := c.categoryService.Create(category)
	if err != nil {
		return nil, err
	}

	respon := &response.CategoryResponse{
		Id:          categoryCreate.Id,
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		AdminID:     admin.Id,
	}

	return respon, nil
}

func NewCategoryService(cr repository.CategoryRepository) CategoryService {
	return &categoryService{categoryService: cr}
}
