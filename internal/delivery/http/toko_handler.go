package http

import (
	"fmt"
	"goProjectEvermos/internal/usecase"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TokoHandler struct {
	tokoUsecase usecase.TokoUsecase
}

func NewTokoHandler(tokoUsecase usecase.TokoUsecase) *TokoHandler {
	return &TokoHandler{tokoUsecase}
}

func (h *TokoHandler) GetMyToko(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	toko, err := h.tokoUsecase.GetMyToko(userID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Toko tidak ditemukan", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan info toko", toko)
}

func (h *TokoHandler) UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	tokoID, err := strconv.Atoi(c.Params("id_toko"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Toko tidak valid", nil)
	}

	var input usecase.UpdateTokoInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	file, err := c.FormFile("photo")
	var fileLocation string = ""

	if err == nil {
		fileName := fmt.Sprintf("%d-%s", time.Now().Unix())
		fileLocation = fmt.Sprintf("/public/images/%s", fileName)

		err = c.SaveFile(file, "."+fileLocation)
		if err != nil {
			return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan file", err.Error())
		}
	}

	updatedToko, err := h.tokoUsecase.UpdateToko(userID, uint(tokoID), input, fileLocation)
	if err != nil {
		return ErrorResponse(c, fiber.StatusForbidden, "Gagal memperbarui toko", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Toko berhasil diperbarui", updatedToko)
}

func (h *TokoHandler) GetAllToko(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	namaToko := c.Query("nama")

	tokos, err := h.tokoUsecase.GetAllToko(page, limit, namaToko)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mendapatkan daftar toko", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan daftar toko", tokos)
}


func (h *TokoHandler) GetTokoByID(c *fiber.Ctx) error {
	tokoID, err := strconv.Atoi(c.Params("id_toko"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Toko tidak valid", nil)
	}

	toko, err := h.tokoUsecase.GetTokoByID(uint(tokoID))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Toko tidak ditemukan", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail toko", toko)
}