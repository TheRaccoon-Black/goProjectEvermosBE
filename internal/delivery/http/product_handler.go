package http

import (
	"fmt"
	"goProjectEvermos/internal/usecase"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	hargaReseller, _ := strconv.Atoi(c.FormValue("harga_reseller"))
	hargaKonsumen, _ := strconv.Atoi(c.FormValue("harga_konsumen"))
	stok, _ := strconv.Atoi(c.FormValue("stok"))
	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))

	input := usecase.CreateProductInput{
		NamaProduk:    c.FormValue("nama_produk"),
		HargaReseller: uint(hargaReseller),
		HargaKonsumen: uint(hargaKonsumen),
		Stok:          uint(stok),
		Deskripsi:     c.FormValue("deskripsi"),
		CategoryID:    uint(categoryID),
	}

	form, err := c.MultipartForm()
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Gagal memproses form", err.Error())
	}

	files := form.File["photos"]
	var fileLocations []string
	for _, file := range files {
		fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
		location := fmt.Sprintf("/public/images/%s", fileName)
		if err := c.SaveFile(file, "."+location); err != nil {
			return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan file", err.Error())
		}
		fileLocations = append(fileLocations, location)
	}

	_, err = h.productUsecase.CreateProduct(userID, input, fileLocations)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membuat produk", err.Error())
	}
	return SuccessResponse(c, fiber.StatusCreated, "Produk berhasil dibuat", nil)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	filter := map[string]string{
		"nama_produk": c.Query("nama_produk"),
		"category_id": c.Query("category_id"),
		"toko_id":     c.Query("toko_id"),
		"min_harga":   c.Query("min_harga"),
		"max_harga":   c.Query("max_harga"),
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	products, err := h.productUsecase.GetAllProducts(filter, page, limit)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mendapatkan produk", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan semua produk", products)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID produk tidak valid", nil)
	}
	product, err := h.productUsecase.GetProductByID(uint(id))
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, "Produk tidak ditemukan", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan detail produk", product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID produk tidak valid", nil)
	}

	hargaReseller, _ := strconv.Atoi(c.FormValue("harga_reseller"))
	hargaKonsumen, _ := strconv.Atoi(c.FormValue("harga_konsumen"))
	stok, _ := strconv.Atoi(c.FormValue("stok"))
	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))

	input := usecase.UpdateProductInput{
		NamaProduk:    c.FormValue("nama_produk"),
		HargaReseller: uint(hargaReseller),
		HargaKonsumen: uint(hargaKonsumen),
		Stok:          uint(stok),
		Deskripsi:     c.FormValue("deskripsi"),
		CategoryID:    uint(categoryID),
	}

	product, err := h.productUsecase.UpdateProduct(userID, uint(productID), input)
	if err != nil {
		return ErrorResponse(c, fiber.StatusForbidden, "Gagal memperbarui produk", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Produk berhasil diperbarui", product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "ID produk tidak valid", nil)
	}

	err = h.productUsecase.DeleteProduct(userID, uint(productID))
	if err != nil {
		return ErrorResponse(c, fiber.StatusForbidden, "Gagal menghapus produk", err.Error())
	}
	return SuccessResponse(c, fiber.StatusOK, "Produk berhasil dihapus", nil)
}