package postgres

import (
	"github.com/AureoFJunior/muguiwaraGo/internal/domain/entities"
	"github.com/AureoFJunior/muguiwaraGo/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &GormProductRepository{db: db}
}

func (repo *GormProductRepository) Create(product *entities.ValidatedProduct) error {
	dbProduct := ToDBProduct(product)

	if err := repo.db.Create(dbProduct).Error; err != nil {
		return err
	}

	storedProduct, err := repo.FindByID(dbProduct.ID)
	if err != nil {
		return err
	}

	*product = *storedProduct

	return nil
}

func (repo *GormProductRepository) FindByID(id uuid.UUID) (*entities.ValidatedProduct, error) {
	var dbProduct Product

	return FromDBProduct(&dbProduct)
}

func (repo *GormProductRepository) GetAll() ([]*entities.ValidatedProduct, error) {
	var dbProducts []Product
	var err error

	products := make([]*entities.ValidatedProduct, len(dbProducts))
	for i, dbProduct := range dbProducts {
		products[i], err = FromDBProduct(&dbProduct)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func (repo *GormProductRepository) Update(product *entities.ValidatedProduct) error {
	dbProduct := ToDBProduct(product)
	err := repo.db.Model(&Product{}).Where("id = ?", dbProduct.ID).Updates(dbProduct).Error
	if err != nil {
		return err
	}

	storedProduct, err := repo.FindByID(dbProduct.ID)
	if err != nil {
		return err
	}

	*product = *storedProduct

	return nil
}

func (repo *GormProductRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&Product{}, id).Error
}
