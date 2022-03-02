package repository

type ProductInterfaceRpository interface {
	CreateProduct(name string, price float64) error
}
