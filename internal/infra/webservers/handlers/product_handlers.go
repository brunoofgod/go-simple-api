package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brunoofgod/go-simple-api/internal/dto"
	"github.com/brunoofgod/go-simple-api/internal/entity"
	"github.com/brunoofgod/go-simple-api/internal/infra/database"
	entityPkg "github.com/brunoofgod/go-simple-api/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	productDB database.ProductInterface
}

func NewProductHandler(productDB database.ProductInterface) *ProductHandler {
	return &ProductHandler{productDB: productDB}
}

// CreateProduct godoc
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductInputDto true "Create Product"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInputDto

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	newProduct, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.productDB.Create(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary Get Product
// @Description Get Product
// @Tags Product
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200 {object} dto.ProductOutputDto
// @Failure 400 {object} Error
// @Router /products/{id} [get]
func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.ProductOutputDto

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productFound, err := p.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID = productFound.ID.String()
	product.Name = productFound.Name
	product.Price = productFound.Price

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateProduct godoc
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Param product body dto.CreateProductInputDto true "Update Product"
// @Success 200
// @Failure 400 {object} Error
// @Router /products/{id} [put]
func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInputDto

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productEntity, err := p.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productEntity.Name = product.Name
	productEntity.Price = product.Price

	err = p.productDB.Update(productEntity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200
// @Failure 400 {object} Error
// @Router /products/{id} [delete]
func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = p.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.productDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ListProducts godoc
// @Summary List Products
// @Description List Products
// @Tags Product
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Success 200 {array} dto.ProductOutputDto
// @Failure 400 {object} Error
// @Router /products [get]
func (p *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := p.productDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var productsDto []dto.ProductOutputDto
	for i := 0; i < len(products); i++ {
		productsDto = append(productsDto, dto.ProductOutputDto{
			ID:    products[i].ID.String(),
			Name:  products[i].Name,
			Price: products[i].Price,
		})
	}

	err = json.NewEncoder(w).Encode(productsDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
