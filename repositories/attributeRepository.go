package repositories

import (
	"context"
	"errors"
	"log"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/ent/attribute"
	"prodcat/ent/attributevaluenum"
	"prodcat/ent/attributevaluestring"
	"prodcat/ent/attributevariantnum"
	"prodcat/ent/attributevariantstring"
	"prodcat/ent/schema"
	"strconv"
)

type AttributeRepository struct {
	db *ent.Client
}

func NewAttributeRepository(db *ent.Client) *AttributeRepository {
	return &AttributeRepository{
		db: db,
	}
}

// ///////// ATRIBUTES /////////////////////////////////////////
func (ar *AttributeRepository) GetAllAttributes(ctx context.Context, l, o int, aType string) []*ent.Attribute {
	if aType != "" {
		return ar.db.Attribute.Query().Where(attribute.TypeEQ(aType)).Offset(o).Limit(l).AllX(ctx)
	}
	return ar.db.Attribute.Query().Offset(o).Limit(l).AllX(ctx)
}

func (ar *AttributeRepository) SearchAttributes(ctx context.Context, offset, limit int, query string) []*ent.Attribute {
	// @todo Check query type before search
	rows, err := ar.db.QueryContext(
		ctx,
		`SELECT id, name FROM attributes WHERE to_tsvector(name) @@ to_tsquery($1)`, query+":*",
	)

	if err != nil {
		panic(err.Error)
	}

	defer rows.Close()

	attributeids := []int{}

	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		attributeids = append(attributeids, int(id))
	}
	// This is bad. Double work for searching instead of ent.model hydration. But idk if it's even possible
	return ar.db.Attribute.Query().Where(attribute.IDIn(attributeids...)).Offset(offset).Limit(limit).AllX(ctx)
}

func (ar *AttributeRepository) GetAttribute(ctx context.Context, id int) (*ent.Attribute, error) {
	res, err := ar.db.Attribute.Query().Where(attribute.IDEQ(id)).Only(ctx)
	return res, err
}

func (ar *AttributeRepository) GetAttributeByName(ctx context.Context, name string) (*ent.Attribute, error) {
	result, err := ar.db.Attribute.Query().Where(attribute.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ar *AttributeRepository) CreateAttribute(ctx context.Context, dto dto.AttributeDTO) *ent.Attribute {
	attr := ar.db.Attribute.Create().
		SetName(dto.Name).
		SetType(dto.Type).
		SetDescription(dto.Description).
		SaveX(ctx)
	return attr
}

func (ar *AttributeRepository) UpdateAttribute(ctx context.Context, aDto dto.AttributeDTO) (attrUpdated dto.AttributeWithVariantsDTO, err error) {
	attr, _ := ar.db.Attribute.Get(ctx, aDto.ID)
	_ = ar.db.Attribute.UpdateOne(attr).
		SetName(aDto.Name).
		SetType(aDto.Type).
		SetDescription(aDto.Description).
		SaveX(ctx)

	return ar.GetAttributeWithVariants(ctx, attr.ID)
}

func (ar *AttributeRepository) GetAttributeWithVariants(ctx context.Context, id int) (dto.AttributeWithVariantsDTO, error) {
	attr, err := ar.db.Attribute.
		Query().
		WithAttributeVariantsString().
		WithAttributeVariantsNum().
		Where(attribute.IDEQ(id)).
		Only(ctx)
	return dto.NewAttrWithVariants(attr), err
}

// ///////// ATRIBUTE VARIANTS /////////////////////////////////////////
func (ar *AttributeRepository) CreateVariant(ctx context.Context, a *ent.Attribute, val interface{}) (dto.AttributeValue, error) {
	switch a.Type {
	case schema.AttrBoolType:
		return nil, errors.New("bool attribute can't have variants")
	case schema.AttrStringType:
		res, err := ar.CreateStringVariant(ctx, a, val.(string))
		if err != nil {
			return nil, err
		}
		return res, nil
	case schema.AttrNumType:
		val, err := strconv.Atoi(val.(string))
		if err != nil {
			return nil, err
		}
		res, err := ar.CreateNumVariant(ctx, a, val)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, errors.New("invalid type of attribute")
	}
}

func (ar *AttributeRepository) GetVariantByName(ctx context.Context, a *ent.Attribute, value interface{}) (dto.AttributeValue, error) {
	switch a.Type {
	case schema.AttrBoolType:
		panic("bool attribute can't have variants")
	case schema.AttrStringType:
		return ar.db.AttributeVariantString.Query().Where(attributevariantstring.ValueEQ(value.(string))).Only(ctx)
	case schema.AttrNumType:
		value, err := strconv.Atoi(value.(string))
		if err != nil {
			panic(err)
		}
		return ar.db.AttributeVariantNum.Query().Where(attributevariantnum.ValueEQ(value)).Only(ctx)
	default:
		panic("invalid type of attribute")
	}
}

func (ar *AttributeRepository) GetVariant(ctx context.Context, a *ent.Attribute, id int) (*dto.AttributeVariantDTO, error) {
	switch a.Type {
	case schema.AttrBoolType:
		panic("bool attribute can't have variants")
	case schema.AttrStringType:
		variant, err := ar.db.AttributeVariantString.Query().Where(attributevariantstring.IDEQ(id)).Only(ctx)
		if err != nil {
			return nil, err
		}
		return &dto.AttributeVariantDTO{ID: variant.ID, Value: variant.Value, Type: "string"}, nil
	case schema.AttrNumType:
		variant, err := ar.db.AttributeVariantNum.Query().Where(attributevariantnum.IDEQ(id)).Only(ctx)
		if err != nil {
			return nil, err
		}
		return &dto.AttributeVariantDTO{ID: variant.ID, Value: strconv.Itoa(variant.Value), Type: "numeric"}, nil
	default:
		panic("invalid type of attribute")
	}
}

func (ar *AttributeRepository) CreateStringVariant(ctx context.Context, a *ent.Attribute, val string) (*ent.AttributeVariantString, error) {
	res, err := ar.db.AttributeVariantString.Create().SetValue(val).SetAttribute(a).Save(ctx)
	return res, err
}

func (ar *AttributeRepository) FindVariantById(ctx context.Context, id int, vType string) (interface{}, error) {
	switch vType {
	case schema.AttrBoolType:
		return nil, errors.New("bool attribute can't have variants")
	case schema.AttrStringType:
		res, err := ar.db.AttributeVariantString.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return res, nil
	case schema.AttrNumType:
		res, err := ar.db.AttributeVariantNum.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, errors.New("invalid type of attribute")
	}
}

func (ar *AttributeRepository) FindVariantString(ctx context.Context, id int) (*ent.AttributeVariantString, error) {
	res, err := ar.db.AttributeVariantString.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ar *AttributeRepository) FindVariantNum(ctx context.Context, id int) (*ent.AttributeVariantNum, error) {
	res, err := ar.db.AttributeVariantNum.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ar *AttributeRepository) UpdateVariantNum(ctx context.Context, v *ent.AttributeVariantNum, n int) (*ent.AttributeVariantNum, error) {
	updated, err := v.Update().SetValue(n).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (ar *AttributeRepository) UpdateVariantString(ctx context.Context, v *ent.AttributeVariantString, n string) (*ent.AttributeVariantString, error) {
	updated, err := v.Update().SetValue(n).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (ar *AttributeRepository) UpdateNumVariant(ctx context.Context, v *ent.AttributeVariantNum, val int) (*ent.AttributeVariantNum, error) {
	updated, err := v.Update().SetValue(val).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (ar *AttributeRepository) CreateNumVariant(ctx context.Context, a *ent.Attribute, val int) (*ent.AttributeVariantNum, error) {
	res, err := ar.db.AttributeVariantNum.Create().SetValue(val).SetAttribute(a).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ar *AttributeRepository) AttachVariantToProduct(ctx context.Context, p *ent.Product, v dto.AttributeValue) {
	switch v := v.(type) {
	case *ent.AttributeVariantString:
		{
			ar.db.AttributeValueString.Create().SetVariant(v).SetProduct(p).SaveX(ctx)
		}
	case *ent.AttributeVariantNum:
		{
			ar.db.AttributeValueNum.Create().SetVariant(v).SetProduct(p).SaveX(ctx)
		}
	}
}

func (ar *AttributeRepository) DeleteVariant(ctx context.Context, vDto dto.AttributeVariantDTO) error {
	switch vDto.Type {
	case "string":
		v, _ := ar.FindVariantString(ctx, vDto.ID)
		ar.db.AttributeValueString.Delete().Where(attributevaluestring.VariantIDEQ(v.ID)).Exec(ctx)
		err := ar.db.AttributeVariantString.DeleteOne(v).Exec(ctx)
		if err != nil {
			panic(err.Error())
		}
	case "numeric":
		v, _ := ar.FindVariantNum(ctx, vDto.ID)
		ar.db.AttributeValueNum.Delete().Where(attributevaluenum.VariantIDEQ(v.ID)).Exec(ctx)
		err := ar.db.AttributeVariantNum.DeleteOne(v).Exec(ctx)
		if err != nil {
			panic(err.Error())
		}
	default:
		panic("invalid type of attribute")
	}
	return nil
}

// ///////// ATRIBUTE VALUES /////////////////////////////////////////
func (ar *AttributeRepository) GetAllAttributeValues(ctx context.Context) map[string]any {
	attributeValues := make(map[string]any, 3)
	attributeValues["attributeValuesString"] = ar.GetAllStringAttributeValues(ctx)
	attributeValues["attributeValuesBool"] = ar.GetAllBoolAttributeValues(ctx)
	attributeValues["attributeValuesNums"] = ar.GetAllNumAttributeValues(ctx)
	return attributeValues
}

func (ar *AttributeRepository) GetAllBoolAttributeValues(ctx context.Context) []*ent.AttributeValueBool {
	return ar.db.AttributeValueBool.Query().AllX(ctx)
}

func (ar *AttributeRepository) GetAllStringAttributeValues(ctx context.Context) []*ent.AttributeValueString {
	return ar.db.AttributeValueString.Query().AllX(ctx)
}

func (ar *AttributeRepository) GetAllNumAttributeValues(ctx context.Context) []*ent.AttributeValueNum {
	return ar.db.AttributeValueNum.Query().AllX(ctx)
}

func (ar *AttributeRepository) CreateBoolValue(ctx context.Context, a *ent.Attribute, val bool, p *ent.Product) *ent.AttributeValueBool {
	return ar.db.AttributeValueBool.Create().SetAttribute(a).SetValue(val).SetProduct(p).SaveX(ctx)
}
