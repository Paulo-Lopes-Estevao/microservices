package entities

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price,string"`
}

func NewProduct(name string, price float64) (*Product, error) {
	products := &Product{
		Name:  name,
		Price: price,
	}
	err := products.priceValid()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) priceValid() error {
	if p.Price <= 0 {
		return errors.New("price not lower the 0")
	}
	return nil
}

func (p *Product) validate() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}
