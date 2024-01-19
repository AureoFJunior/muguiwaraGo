package services

import (
	"github.com/AureoFJunior/muguiwaraGo/internal/application/command"
	"github.com/AureoFJunior/muguiwaraGo/internal/application/mapper"
	"github.com/AureoFJunior/muguiwaraGo/internal/domain/entities"
	"github.com/AureoFJunior/muguiwaraGo/internal/domain/repositories"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepository repositories.ProductRepository
	sellerRepository  repositories.SellerRepository
}

func NewProductService(
	productRepository repositories.ProductRepository,
	sellerRepository repositories.SellerRepository,
) *ProductService {
	return &ProductService{productRepository: productRepository, sellerRepository: sellerRepository}
}

func (s *ProductService) CreateProduct(productCommand *command.CreateProductCommand) (*command.CreateProductCommandResult, error) {
	storedSeller, err := s.sellerRepository.FindByID(productCommand.SellerID)
	if err != nil {
		return nil, err
	}

	var newProduct = entities.NewProduct(
		productCommand.Name,
		productCommand.Price,
		*storedSeller,
	)

	validatedProduct, err := entities.NewValidatedProduct(newProduct)
	if err != nil {
		return nil, err
	}

	err = s.productRepository.Create(validatedProduct)
	if err != nil {
		return nil, err
	}

	var result command.CreateProductCommandResult
	result.Result = mapper.NewProductResultFromEntity(validatedProduct)

	return &result, nil
}

func (s *ProductService) GetAllProducts() ([]*entities.ValidatedProduct, error) {
	return s.productRepository.GetAll()
}

func (s *ProductService) FindProductByID(id uuid.UUID) (*entities.ValidatedProduct, error) {
	return s.productRepository.FindByID(id)
}
