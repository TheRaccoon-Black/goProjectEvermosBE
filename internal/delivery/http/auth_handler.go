package http

import (
	"goProjectEvermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input usecase.RegisterUserInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	registeredUser, err := h.authUsecase.Register(input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal melakukan registrasi", err.Error())
	}

	type userResponse struct {
		ID    uint   `json:"id"`
		Nama  string `json:"nama"`
		Email string `json:"email"`
		NoTelp string `json:"no_telp"`
	}

	response := userResponse{
		ID:    registeredUser.ID,
		Nama:  registeredUser.Nama,
		Email: registeredUser.Email,
		NoTelp: registeredUser.NoTelp,
	}

	return SuccessResponse(c, fiber.StatusOK, "Registrasi berhasil", response)
}