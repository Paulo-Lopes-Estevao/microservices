package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/likexian/gokit/assert"
	"net/http"
	"net/http/httptest"
	"product/adapter/rest/controllers"
	"product/domain/interfaces/mocks"
	"product/usecase/dto"
	"product/usecase/system"
	"testing"
)

func TestCreateProductPOST(t *testing.T) {
	e := echo.New()

	data := dto.ProductDtoInput{
		Name:  "Manga",
		Price: 1000,
	}

	valueData, _ := json.Marshal(data)

	request := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(valueData))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	controller := gomock.NewController(t)
	defer controller.Finish()

	IproductRepository := mocks.NewMockProductInterfaceRpository(controller)
	IproductRepository.EXPECT().CreateProduct(data.Name, data.Price).Return(nil)

	UproductSystem := system.NewProductUsecaseSystem(IproductRepository)

	RproductController := controllers.NewProductController(UproductSystem)

	err := RproductController.CreateProduct(ctx)
	if err != nil {
		assert.Equal(t, http.StatusCreated, recorder.Code, "error status code != 201")
	}
}
