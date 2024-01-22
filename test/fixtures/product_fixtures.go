package fixtures

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/repositories"
	"strconv"

	_ "github.com/xiaoqidun/entps"
)

func ProductFixtures(ctx context.Context, db *ent.Client) {

	pr := repositories.NewProductRepository(db)
	ar := repositories.NewAttributeRepository(db)
	csvFile, err := os.Open("./test/fixtures/product_fixtures.csv")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
	}

	products, err := NewProducts(ctx, csvFile, pr, ar)
	if err != nil {
		fmt.Println(fmt.Errorf("error generating from file: %v", err))
	}

	keyboard1 := products[0]

	// attrKeys, _ := ar.CreateAttribute(ctx, schema.AttrNumType, "Keys", "Number of keys")

	keyboard1Restored := pr.FindProductByID(ctx, keyboard1.ID, true)
	println(keyboard1Restored.ID)

	// productDTO = repositories.ProductDTO{Name: "Mouse", Description: "Cheapest mouse"}
	// product, _ := pr.CreateProduct(ctx, productDTO)
	// attrString, _ := ar.CreateAttribute(ctx, schema.AttrStringType, "Color", "Main poduct color")
	// attrNum, _ := ar.CreateAttribute(ctx, schema.AttrNumType, "Buttons", "Number of buttons")

	// ar.CreateValue(ctx, attrString, "White", product)
	// ar.CreateValue(ctx, attrNum, 3, product)
}

func NewProducts(
	ctx context.Context,
	csvFile io.Reader,
	pr *repositories.ProductRepository,
	ar *repositories.AttributeRepository,
) ([]*ent.Product, error) {
	var products []*ent.Product
	reader := csv.NewReader(csvFile)
	var product *ent.Product

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return products, err
		}

		if len(line) < 1 {
			return products, fmt.Errorf("invalid file structure")
		}

		name := line[0]
		description := line[1]
		code := line[2]
		barcode := line[3]
		atype := line[4]
		attName := line[5]
		attDesc := line[6]

		if name != "@" && description != "@" && barcode != "@" {
			prodDTO := dto.ProductDTO{
				Name:        line[0],
				Description: line[1],
				Code:        code,
				Barcode:     line[3],
			}
			product, _ = pr.CreateProduct(ctx, prodDTO)
		}

		products = append(products, product)

		attribute, err := ar.GetAttributeByName(ctx, attName)
		if err != nil {

			attribute = ar.CreateAttribute(ctx, dto.AttributeDTO{Name: attName, Type: atype, Description: attDesc})
		}

		switch atype {
		case "string", "numeric":
			{
				attributeVariant, err := ar.GetVariantByName(ctx, attribute, line[7])
				if err != nil {
					attributeVariant, _ = ar.CreateVariant(ctx, attribute, line[7])
				}

				ar.AttachVariantToProduct(ctx, product, attributeVariant)
			}
		case "bool":
			{
				attributValueBool, _ := strconv.ParseBool(line[7])
				ar.CreateBoolValue(ctx, attribute, attributValueBool, product)
			}
		default:
			{
				println(atype)
				panic("fixtures not valid type")
			}
		}

	}

	return products, nil
}
