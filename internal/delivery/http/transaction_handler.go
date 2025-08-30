package http

import (
	"goProjectEvermos/internal/usecase"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	trxUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(trxUsecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{trxUsecase}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var input usecase.CreateTransactionInput
	if err := c.BodyParser(&input); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses request", err.Error())
	}

	newTrx, err := h.trxUsecase.CreateTransaction(userID, input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal membuat transaksi", err.Error())
	}

	return SuccessResponse(c, fiber.StatusCreated, "Transaksi berhasil dibuat", newTrx)
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	transactions, err := h.trxUsecase.GetAllTransactions(userID)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mendapatkan riwayat transaksi", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan riwayat transaksi", transactions)
}

func (h *TransactionHandler) GetTransactionByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	trxID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID Transaksi tidak valid", nil)
	}

	transaction, err := h.trxUsecase.GetTransactionByID(userID, uint(trxID))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Transaksi tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail transaksi", transaction)
}