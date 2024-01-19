package command

import "github.com/google/uuid"

type CreateProductCommand struct {
	ID    uuid.UUID
	Name  string
	Price float64
}
