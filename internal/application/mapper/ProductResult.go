package mapper

import (
	"github.com/AureoFJunior/muguiwaraGo/internal/application/command"
	"github.com/AureoFJunior/muguiwaraGo/internal/domain/entities"
)

func NewProductResultFromEntity(product *entities.ValidatedProduct) command.ProductResult {
	return command.ProductResult{
		Id:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}
