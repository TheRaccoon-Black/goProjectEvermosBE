package http

import (
	"goProjectEvermos/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ProvCityHandler struct {
	provCityUsecase usecase.ProvCityUsecase
}

func NewProvCityHandler(provCityUsecase usecase.ProvCityUsecase) *ProvCityHandler {
	return &ProvCityHandler{provCityUsecase}
}

func (h *ProvCityHandler) GetAllProvinces(c *fiber.Ctx) error {
	provinces, err := h.provCityUsecase.GetAllProvinces()
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data provinsi", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan semua provinsi", provinces)
}

func (h *ProvCityHandler) GetCitiesByProvinceID(c *fiber.Ctx) error {
	provinceID := c.Params("prov_id")
	cities, err := h.provCityUsecase.GetCitiesByProvinceID(provinceID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data kota", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan data kota", cities)
}

func (h *ProvCityHandler) GetProvinceByID(c *fiber.Ctx) error {
	provinceID := c.Params("prov_id")
	province, err := h.provCityUsecase.GetProvinceByID(provinceID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Provinsi tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail provinsi", province)
}

func (h *ProvCityHandler) GetCityByID(c *fiber.Ctx) error {
	cityID := c.Params("city_id")
	city, err := h.provCityUsecase.GetCityByID(cityID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Kota tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail kota", city)
}