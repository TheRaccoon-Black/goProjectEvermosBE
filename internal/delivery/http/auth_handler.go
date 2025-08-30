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

	response := fiber.Map{
		"id":      registeredUser.ID,
		"nama":    registeredUser.Nama,
		"email":   registeredUser.Email,
		"no_telp": registeredUser.NoTelp,
	}

	return SuccessResponse(c, fiber.StatusOK, "Registrasi berhasil", response)
}


func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input usecase.LoginUserInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	token, err := h.authUsecase.Login(input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusUnauthorized, "Gagal login", err.Error())
	}

	return SuccessResponse(c, fiber.StatusOK, "Login berhasil", fiber.Map{"token": token})
}