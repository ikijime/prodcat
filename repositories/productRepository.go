package repositories

import (
	"context"
	"log"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/ent/product"
	"strconv"
)

type ProductRepository struct {
	db *ent.Client
}

func NewProductRepository(db *ent.Client) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetAllProducts(ctx context.Context, limit, offset int) []*ent.Product {
	return pr.db.Product.Query().Offset(offset).Limit(limit).AllX(ctx)
}

func (pr *ProductRepository) SearchProduct(ctx context.Context, offset, limit int, query string) []*ent.Product {
	// @todo Check query type before search
	rows, err := pr.db.QueryContext(
		ctx,
		`SELECT id, name FROM products WHERE to_tsvector(name) || to_tsvector(code::text) || to_tsvector(barcode::text) @@ to_tsquery($1)`, query+":*",
	)

	if err != nil {
		panic(err.Error)
	}

	defer rows.Close()

	products := []int{}

	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		products = append(products, int(id))
	}
	// This is bad. Double work for searching instead of ent.model hydration. But idk if it's even possible
	return pr.db.Product.Query().Where(product.IDIn(products...)).Offset(offset).Limit(limit).AllX(ctx)
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, dto dto.ProductDTO) (*ent.Product, error) {
	code, _ := strconv.Atoi(dto.Code)
	product, err := pr.db.Product.Create().
		SetName(dto.Name).
		SetCode(code).
		SetBarcode(dto.Barcode).
		SetDescription(dto.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) FindProductByID(ctx context.Context, id int, withAttrs bool) (*ent.Product, error) {
	if withAttrs {
		return pr.db.Product.
			Query().
			WithAttributeValuesBool(func(q *ent.AttributeValueBoolQuery) { q.WithAttribute() }).
			WithAttributeValuesString(func(q *ent.AttributeValueStringQuery) {
				q.WithVariant(
					func(q *ent.AttributeVariantStringQuery) { q.WithAttribute() })
			}).
			WithAttributeValuesNum(func(q *ent.AttributeValueNumQuery) {
				q.WithVariant(func(q *ent.AttributeVariantNumQuery) { q.WithAttribute() })
			}).
			Where(product.IDEQ(id)).
			Only(ctx)
	} else {
		return pr.db.Product.
			Query().
			Where(product.IDEQ(id)).
			Only(ctx)
	}
}

func (pr *ProductRepository) UpdateProductFromDTO(ctx context.Context, prodDTO *dto.ProductDTO) (*ent.Product, error) {
	product, err := pr.FindProductByID(ctx, prodDTO.ID, true)
	if err != nil {
		return nil, err
	}

	return product.Update().SetBarcode(prodDTO.Barcode).Save(ctx)
}
