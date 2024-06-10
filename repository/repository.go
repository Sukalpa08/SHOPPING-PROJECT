package repository

import (
	"log"
	"shops/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(input *models.Item) (*models.Item, error)
}
type itemRepository struct {
	db *gorm.DB
}

func connectDB(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, config)
	return db, err
}

// NewItemRepository is a constructor for ItemRepository
func NewItemRepository(dialector gorm.Dialector, config *gorm.Config) *ItemRepository {
	db, err := connectDB(dialector, config)
	if err != nil {
		log.Fatalf("Failed to connect to the database due to error: %s", err)
		return nil
	}

	var r ItemRepository = &itemRepository{db: db}
	return &r
}
func (r *itemRepository) CreateItem(input *models.Item) (*models.Item, error) {
	f := false
	item := models.Item{
		Name:  input.Name,
		Price: input.Price,
		Sold:  &f,
	}
	if err := r.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
