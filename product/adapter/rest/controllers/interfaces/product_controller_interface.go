package interfaces

import "product/adapter/rest/presenter"

type IproductController interface {
	CreateProduct(ctx presenter.ProductPresenterContext) error
}
