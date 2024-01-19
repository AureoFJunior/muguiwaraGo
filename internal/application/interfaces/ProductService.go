package interfaces

import (
	"github.com/AureoFJunior/muguiwaraGo/internal/application/command"
	"github.com/AureoFJunior/muguiwaraGo/internal/domain/entities"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(productCommand *command.CreateProductCommand) (*command.CreateProductCommandResult, error)
	GetAllProducts() ([]*entities.ValidatedProduct, error)
	FindProductByID(id uuid.UUID) (*entities.ValidatedProduct, error)
}
