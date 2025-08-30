package http

import (
	"goProjectEvermos/internal/usecase"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type AlamatHandler struct {
	alamatUsecase usecase.AlamatUsecase
}

func NewAlamatHandler(alamatUsecase usecase.AlamatUsecase) *AlamatHandler {
	return &AlamatHandler{alamatUsecase}
}

func (h *AlamatHandler) CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var input usecase.CreateAlamatInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	newAlamat, err := h.alamatUsecase.CreateAlamat(userID, input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan alamat", err.Error())
	}

	return SuccessResponse(c, fiber.StatusCreated, "Alamat berhasil ditambahkan", newAlamat)
}

func (h *AlamatHandler) GetAllAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	alamats, err := h.alamatUsecase.GetAllAlamat(userID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mendapatkan alamat", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan semua alamat", alamats)
}

// Tambahkan handler baru GetAlamatByID
func (h *AlamatHandler) GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// Ambil ID dari parameter URL dan konversi ke integer
	alamatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Alamat tidak valid", err.Error())
	}

	alamat, err := h.alamatUsecase.GetAlamatByID(userID, uint(alamatID))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Alamat tidak ditemukan", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail alamat", alamat)
}

func (h *AlamatHandler) UpdateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	alamatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Alamat tidak valid", nil)
	}

	var input usecase.UpdateAlamatInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	updatedAlamat, err := h.alamatUsecase.UpdateAlamat(userID, uint(alamatID), input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Gagal memperbarui, alamat tidak ditemukan", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Alamat berhasil diperbarui", updatedAlamat)
}


// Tambahkan handler baru DeleteAlamat
func (h *AlamatHandler) DeleteAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	alamatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Alamat tidak valid", nil)
	}

	err = h.alamatUsecase.DeleteAlamat(userID, uint(alamatID))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Gagal menghapus, alamat tidak ditemukan", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Alamat berhasil dihapus", nil)
}