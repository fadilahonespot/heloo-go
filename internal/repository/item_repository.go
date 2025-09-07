package repository

import (
	"gorm.io/gorm"

	"heloo-go/internal/domain"
)

type ItemRepository interface {
	AutoMigrate() error
	List() ([]domain.Item, error)
	GetByID(id string) (domain.Item, error)
	Create(in domain.Item) (domain.Item, error)
	Update(in domain.Item) (domain.Item, error)
	Delete(id string) (int64, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) AutoMigrate() error { return r.db.AutoMigrate(&domain.Item{}) }

func (r *itemRepository) List() ([]domain.Item, error) {
	var items []domain.Item
	if err := r.db.Order("id desc").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *itemRepository) GetByID(id string) (domain.Item, error) {
	var item domain.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return domain.Item{}, err
	}
	return item, nil
}

func (r *itemRepository) Create(in domain.Item) (domain.Item, error) {
	if err := r.db.Create(&in).Error; err != nil {
		return domain.Item{}, err
	}
	return in, nil
}

func (r *itemRepository) Update(in domain.Item) (domain.Item, error) {
	if err := r.db.Model(&domain.Item{}).Where("id = ?", in.ID).Updates(map[string]any{
		"title":       in.Title,
		"description": in.Description,
	}).Error; err != nil {
		return domain.Item{}, err
	}
	return in, nil
}

func (r *itemRepository) Delete(id string) (int64, error) {
	res := r.db.Delete(&domain.Item{}, id)
	return res.RowsAffected, res.Error
}
