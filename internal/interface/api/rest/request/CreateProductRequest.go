package request

import (
	"github.com/AureoFJunior/muguiwaraGo/internal/application/command"
)

type CreateProductRequest struct {
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

func (req *CreateProductRequest) ToCreateProductCommand() (*command.CreateProductCommand, error) {
	return &command.CreateProductCommand{
		Name:  req.Name,
		Price: req.Price,
	}, nil
}
