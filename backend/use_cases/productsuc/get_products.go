package productsuc

import (
	"backend/domain/entities"
	"context"
)

type getProductsUseCase struct {
	dataStore ProductsDataStorer
}

func NewGetProductsUseCase(ds ProductsDataStorer) *getProductsUseCase {
	return &getProductsUseCase{
		dataStore: ds,
	}
}

func (uc *getProductsUseCase) GetProducts(ctx context.Context) []entities.Product {
	all := uc.dataStore.GetAll()
	return all
}
