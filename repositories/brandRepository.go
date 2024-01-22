package repositories

import (
	"context"
	"prodcat/ent"
)

type BrandRepository struct {
	db *ent.Client
}

func NewBrandRepository(db *ent.Client) *BrandRepository {
	return &BrandRepository{
		db: db,
	}
}

// ///////// BRANDS /////////////////////////////////////////
func (ar *BrandRepository) GetAllBrands(ctx context.Context, l, o int) []*ent.Brand {
	return ar.db.Brand.Query().Offset(o).Limit(l).AllX(ctx)
}
