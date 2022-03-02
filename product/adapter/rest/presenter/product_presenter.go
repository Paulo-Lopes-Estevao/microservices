package presenter

type ProductPresenterContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	JSONBlob(code int, b []byte) error
}
