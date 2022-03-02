package system_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"product/domain/interfaces/mocks"
	"product/usecase/dto"
	"product/usecase/system"
	"testing"
)

func TestAddProduct(t *testing.T) {
	input := dto.ProductDtoInput{
		Name:  faker.Name(),
		Price: 1000.0,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	IproductRepository := mocks.NewMockProductInterfaceRpository(controller)
	IproductRepository.EXPECT().CreateProduct(input.Name, input.Price).Return(nil)

	UproductSystem := system.NewProductUsecaseSystem(IproductRepository)

	output, err := UproductSystem.AddProduct(&input)
	assert.Nil(t, err)
	assert.Equal(t, output.Data.Price, 1000.0)
}
