package services

import (
	"log"
	"shops/models"
	"shops/repository"
)

type ItemService interface {
	CreateItem(input models.CreateItemInput) (*models.Item, error)
}

type itemService struct {
	r *repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) *ItemService {
	if r == nil {
		log.Fatal("Failed to initialize item service, repository is nil")
		return nil
	}
	var s ItemService = &itemService{r: &r}
	return &s
}

func (s *itemService) GetItemRepository() repository.ItemRepository {
	if s.r == nil {
		log.Fatal("Failed to get item repository, it is nil")
		return nil
	}

	return *s.r
}
func (s *itemService) CreateItem(input models.CreateItemInput) (*models.Item, error) {
	r := s.GetItemRepository()

	sold := false
	data := &models.Item{
		Name:  &input.Name,
		Price: &input.Price,
		Sold:  &sold,
	}
	item, err := r.CreateItem(data)
	return item, err
}
