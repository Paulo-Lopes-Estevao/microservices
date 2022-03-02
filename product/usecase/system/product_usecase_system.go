package system

import (
	"errors"
	"product/domain/entities"
	"product/domain/interfaces/repository"
	"product/usecase/dto"
	"product/usecase/interfaces"
)

type productUsecaseSystem struct {
	productRepository repository.ProductInterfaceRpository
}

func NewProductUsecaseSystem(pRepository repository.ProductInterfaceRpository) interfaces.ProductSysUsecaseInterface {
	return &productUsecaseSystem{
		productRepository: pRepository,
	}
}

func (p productUsecaseSystem) AddProduct(productInput *dto.ProductDtoInput) (dto.ProductDtoOutput, error) {
	in_product, err := entities.NewProduct(productInput.Name, productInput.Price)
	if err != nil {
		return responseDataError(err)
	}
	errProd := p.productRepository.CreateProduct(in_product.Name, in_product.Price)
	if errProd != nil {
		return responseDataError(errProd)
	}

	return responseData(in_product)
}

func responseDataError(dataError error) (dto.ProductDtoOutput, error) {
	output := dto.ProductDtoOutput{
		ErrorMessage: dataError.Error(),
	}
	return dto.ProductDtoOutput{}, errors.New(output.ErrorMessage)
}

func responseData(productInput *entities.Product) (dto.ProductDtoOutput, error) {
	return dto.ProductDtoOutput{
			Data: dto.ProductDtoInput{
				Name:  productInput.Name,
				Price: productInput.Price},
			ErrorMessage: "",
		},
		nil
}
