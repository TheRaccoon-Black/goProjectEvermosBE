package http

import (
	"goProjectEvermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	user, err := h.userUsecase.GetProfile(userID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "User tidak ditemukan", err.Error())
	}

	type userResponse struct {
		ID    uint   `json:"id"`
		Nama  string `json:"nama"`
		Email string `json:"email"`
		NoTelp string `json:"no_telp"`
		Role string `json:"role"`
	}

	response := userResponse{
		ID:    user.ID,
		Nama:  user.Nama,
		Email: user.Email,
		NoTelp: user.NoTelp,
		Role: user.Role,
	}

	return SuccessResponse(c, fiber.StatusOK, "Profil berhasil didapatkan", response)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	// Ambil userID dari context
	userID := c.Locals("userID").(uint)

	// Parse request body
	var input usecase.UpdateProfileInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	// Panggil usecase untuk update
	updatedUser, err := h.userUsecase.UpdateProfile(userID, input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui profil", err.Error())
	}

	// Buat DTO response agar konsisten
	type userResponse struct {
		ID        uint   `json:"id"`
		Nama      string `json:"nama"`
		Email     string `json:"email"`
		NoTelp    string `json:"no_telp"`
		Tentang   string `json:"tentang"`
		Pekerjaan string `json:"pekerjaan"`
	}

	response := userResponse{
		ID:        updatedUser.ID,
		Nama:      updatedUser.Nama,
		Email:     updatedUser.Email,
		NoTelp:    updatedUser.NoTelp,
		Tentang:   updatedUser.Tentang,
		Pekerjaan: updatedUser.Pekerjaan,
	}

	return SuccessResponse(c, fiber.StatusOK, "Profil berhasil diperbarui", response)
}
