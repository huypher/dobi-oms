package repository

import (
	"context"
	"time"

	"github.com/pghuy/dobi-oms/domain"

	"github.com/uptrace/bun"
)

type product struct {
	bun.BaseModel `bun:"products"`

	ID          int       `bun:"id"`
	Sku         string    `bun:"sku"`
	Name        string    `bun:"name"`
	Description string    `bun:"description"`
	CategoryID  int       `bun:"category_id"`
	CreatedAt   time.Time `bun:"created_at,type:timestamp with time zone"`
	UpdatedAt   time.Time `bun:"updated_at,type:timestamp with time zone"`
}

type productRepository struct {
	db *bun.DB
}

func NewProductRepository(db *bun.DB) (*productRepository, error) {
	return &productRepository{
		db: db,
	}, nil
}

func (r *productRepository) Save(ctx context.Context, p *domain.Product) error {
	if p == nil {
		return nil
	}

	model := &product{
		Sku:         p.Sku,
		Name:        p.Name,
		Description: p.Description,
		CategoryID:  p.CategoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
