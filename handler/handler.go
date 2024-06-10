package handler

import (
	"log"
	"net/http"
	"shops/models"
	"shops/services"

	"github.com/gin-gonic/gin"
)

type Itemhandler interface {
	CreateItem(c *gin.Context)
}
type itemHandler struct {
	s *services.ItemService
}

func NewItemHandler(s services.ItemService) *itemHandler {
	if s == nil {
		log.Fatal("Failed to initialize item handler, service is nil")
		return nil
	}
	var p = itemHandler{s: &s}
	return &p
}

func (h *itemHandler) SetItemService(s services.ItemService) {
	h.s = &s
}

func (h *itemHandler) GetItemService() services.ItemService {
	if h.s == nil {
		log.Fatal("Failed to get item service, it is nil")
		return nil
	}

	return *h.s
}

func (h *itemHandler) CreateItem(c *gin.Context) {

	var input models.CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an item
	s := h.GetItemService()
	item, err := s.CreateItem(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}
