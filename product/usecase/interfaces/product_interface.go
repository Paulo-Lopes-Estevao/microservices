package interfaces

import "product/usecase/dto"

type ProductSysUsecaseInterface interface {
	AddProduct(productInput *dto.ProductDtoInput) (dto.ProductDtoOutput, error)
}
