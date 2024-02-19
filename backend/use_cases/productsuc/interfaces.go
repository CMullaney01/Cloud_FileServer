package productsuc

import (
	"backend/domain/entities"
)

type ProductsDataStorer interface {
	GetAll() []entities.Product
	Create(product *entities.Product) error
}
