package controllers

import (
	"errors"
	"net/http"
	"product/adapter/rest/controllers/interfaces"
	"product/adapter/rest/presenter"
	"product/usecase/dto"
	IproductUseSystem "product/usecase/interfaces"
)

type productController struct {
	productUsecaseInterface IproductUseSystem.ProductSysUsecaseInterface
}

func NewProductController(Iproduct IproductUseSystem.ProductSysUsecaseInterface) interfaces.IproductController {
	return &productController{
		productUsecaseInterface: Iproduct,
	}
}

func (p productController) CreateProduct(ctx presenter.ProductPresenterContext) error {
	var input dto.ProductDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	usecaseProductOutput, err := p.productUsecaseInterface.AddProduct(&input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, usecaseProductOutput)
}
