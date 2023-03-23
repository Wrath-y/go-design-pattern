package factory_method

type Factory interface {
	CreateProduct() Product
}

type Product interface {
	Operation()
}
