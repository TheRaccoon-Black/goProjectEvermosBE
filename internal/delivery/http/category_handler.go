package http

import (
	"goProjectEvermos/internal/usecase"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var input usecase.CreateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	category, err := h.categoryUsecase.CreateCategory(input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membuat kategori", err.Error())
	}
	return SuccessResponse(c, fiber.StatusCreated, "Kategori berhasil dibuat", category)
}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.categoryUsecase.GetAllCategories()
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mendapatkan kategori", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan semua kategori", categories)
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Kategori tidak valid", nil)
	}
	category, err := h.categoryUsecase.GetCategoryByID(uint(id))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Kategori tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail kategori", category)
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Kategori tidak valid", nil)
	}
	var input usecase.CreateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}
	category, err := h.categoryUsecase.UpdateCategory(uint(id), input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Gagal memperbarui, kategori tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Kategori berhasil diperbarui", category)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Kategori tidak valid", nil)
	}
	err = h.categoryUsecase.DeleteCategory(uint(id))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Gagal menghapus, kategori tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Kategori berhasil dihapus", nil)
}