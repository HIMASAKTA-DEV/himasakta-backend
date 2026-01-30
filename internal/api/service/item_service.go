package service

import (
	"context"

	"github.com/azkaazkun/be-samarta/internal/api/repository"
	"github.com/azkaazkun/be-samarta/internal/dto"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"gorm.io/gorm"
)

type (
	ItemService interface {
		Create(ctx context.Context, req dto.CreateItemRequest) (entity.Item, error)
		GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Item, meta.Meta, error)
		GetById(ctx context.Context, id string) (entity.Item, error)
		Update(ctx context.Context, req dto.UpdateItemRequest) (entity.Item, error)
		Delete(ctx context.Context, id string) error
	}

	itemService struct {
		itemRepository repository.ItemRepository
		db             *gorm.DB
	}
)

func NewItem(itemRepository repository.ItemRepository, db *gorm.DB) ItemService {
	return &itemService{
		itemRepository: itemRepository,
		db:             db,
	}
}

func (s *itemService) Create(ctx context.Context, req dto.CreateItemRequest) (entity.Item, error) {
	newItem, err := s.itemRepository.Create(ctx, nil, entity.Item{
		CategoryCode: req.CategoryCode,
		CategoryName: req.CategoryName,
		Group:        req.Group,
		AccountCode:  req.AccountCode,
		Description:  req.Description,
	})
	if err != nil {
		return entity.Item{}, err
	}

	return newItem, nil
}

func (s *itemService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Item, meta.Meta, error) {
	itemList, newMeta, err := s.itemRepository.GetAll(ctx, nil, metaReq)
	if err != nil {
		return nil, metaReq, err
	}

	return itemList, newMeta, nil
}

func (s *itemService) GetById(ctx context.Context, id string) (entity.Item, error) {
	item, err := s.itemRepository.GetById(ctx, nil, id)
	if err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (s *itemService) Update(ctx context.Context, req dto.UpdateItemRequest) (entity.Item, error) {
	item, err := s.itemRepository.GetById(ctx, nil, req.ID)
	if err != nil {
		return entity.Item{}, err
	}

	if req.CategoryCode != "" {
		item.CategoryCode = req.CategoryCode
	}
	if req.CategoryName != "" {
		item.CategoryName = req.CategoryName
	}
	if req.Group != "" {
		item.Group = req.Group
	}
	if req.AccountCode != "" {
		item.AccountCode = req.AccountCode
	}
	if req.Description != "" {
		item.Description = req.Description
	}

	updatedItem, err := s.itemRepository.Update(ctx, nil, item)
	if err != nil {
		return entity.Item{}, err
	}

	return updatedItem, nil
}

func (s *itemService) Delete(ctx context.Context, id string) error {
	item, err := s.itemRepository.GetById(ctx, nil, id)
	if err != nil {
		return err
	}

	if err := s.itemRepository.Delete(ctx, nil, item); err != nil {
		return err
	}

	return nil
}
