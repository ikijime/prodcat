package controllers

import (
	"errors"
	"log"
	"net/http"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/ent/schema"
	"prodcat/repositories"
	"prodcat/services"
	"prodcat/utils"
	attributeView "prodcat/views/attribute"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type AttributeController struct {
	ar          *repositories.AttributeRepository
	authService *services.AuthService
}

func NewAttributeController(ar *repositories.AttributeRepository, as *services.AuthService) *AttributeController {
	return &AttributeController{
		ar:          ar,
		authService: as,
	}
}

// ///////// ATRIBUTES /////////////////////////////////////////
func (ac AttributeController) GetAllAttributesPage(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	aType := c.Query("type")
	if limit == 0 {
		limit = 10
		offset = 0
	}

	if !utils.StringInSlice(aType, schema.AttrTypes) {
		aType = ""
	}
	var attributes []*ent.Attribute
	searchQuery := c.Query("search")
	if searchQuery != "" {
		attributes = ac.ar.SearchAttributes(c, offset, limit, searchQuery)
		c.HTML(200, "", fbGuard(c, attributeView.Attributes(attributes)))
		return
	}

	attributes = ac.ar.GetAllAttributes(c, limit, offset, aType)
	c.HTML(200, "", fbGuard(c, attributeView.Attributes(attributes)))
}

func (ac AttributeController) SearchAttributes(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	aType := c.Query("type")
	if limit == 0 {
		limit = 1
		offset = 0
	}

	if !utils.StringInSlice(aType, schema.AttrTypes) {
		aType = ""
	}

	attributes := ac.ar.GetAllAttributes(c, limit, offset, aType)
	c.HTML(200, "", fbGuard(c, attributeView.Attributes(attributes)))
}

func (ac AttributeController) GetAttributePage(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert id")
		return
	}
	attribute, _ := ac.ar.GetAttributeWithVariants(c, id)
	c.HTML(200, "", fbGuard(c, attributeView.Attribute(attribute)))
}

func (ac AttributeController) GetAttributeEditPage(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert id")
		return
	}
	attribute, _ := ac.ar.GetAttributeWithVariants(c, id)
	c.HTML(200, "", fbGuard(c, attributeView.AttributeEdit(attribute)))
}

func (ac AttributeController) AddAttributePage(c *gin.Context) {
	attributeDTO := dto.AttributeDTO{}
	c.HTML(200, "", fbGuard(c, attributeView.AttributeAdd(attributeDTO)))
}

func (ac AttributeController) AttributesAddFormParse(c *gin.Context) {
	var attr dto.AttributeDTO
	c.ShouldBind(&attr)

	validate := validator.New()
	err := validate.Struct(&attr)
	if !dto.CheckAttributeType(attr.Type) {
		err = errors.New("Not valid type " + attr.Type)
		attr.TypeErr = err
	}

	if err != nil {
		attr.FillErrors(err)
		c.HTML(http.StatusBadRequest, "", fbGuard(c, attributeView.AttributeAdd(attr)))
		return
	}

	if attr.ID != 0 {
		WritePopupMessage(c, "success", strconv.Itoa(attr.ID))
		print(strconv.Itoa(attr.ID))
		updatedAttr, _ := ac.ar.UpdateAttribute(c, attr)
		c.HTML(
			http.StatusBadRequest, "",
			fbGuard(c, attributeView.AttributeEdit(updatedAttr)))
		return
	}

	// emptyAttr := dto.AttributeDTO{}
	newAttr := ac.ar.CreateAttribute(c, attr)
	WritePopupMessage(c, "success", "added")
	c.Redirect(302, "/attributes/"+strconv.Itoa(newAttr.ID)+"/edit")
	// c.HTML(http.StatusBadRequest, "", fbGuard(c, attributeView.AttributeAdd(emptyAttr)))
}

func (ac AttributeController) GetAllAttributes(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))
	if limit == 0 {
		limit = 10
		offset = 0
	}
	c.JSON(200, ac.ar.GetAllAttributes(c, limit, offset, ""))
}

func (ac AttributeController) GetAttribute(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert id")
		return
	}

	attr, _ := ac.ar.GetAttribute(c, id)
	c.JSON(200, attr)
}

func (ac AttributeController) AddAttribute(c *gin.Context) {
	var dto dto.AttributeDTO
	err := c.ShouldBind(&dto)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	attr := ac.ar.CreateAttribute(c, dto)
	// if err != nil {
	// 	c.JSON(200, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(200, gin.H{"data": attr})
}

// ///////// ATRIBUTE VARIANTS /////////////////////////////////////////
func (ac AttributeController) AttributeVariantEdit(c *gin.Context) {
	aId := c.Param("id")
	attributeId, err := strconv.Atoi(aId)
	if err != nil {
		log.Println("can't convert")
		c.Redirect(302, "/")
		return
	}

	vId := c.Param("vid")
	variantId, err := strconv.Atoi(vId)
	if err != nil {
		log.Println("can't convert")
		c.Redirect(302, "/")
		return
	}

	attr, err := ac.ar.GetAttribute(c, attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	variant, _ := ac.ar.GetVariant(c, attr, variantId)
	c.HTML(200, "", attributeView.AttributeVariantEdit(attr.ID, *variant))
}

func (ac AttributeController) AttributeVariantPut(c *gin.Context) {
	aId := c.Param("id")
	attributeId, err := strconv.Atoi(aId)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	var dtoVariant dto.AttributeVariantDTO
	err = c.ShouldBind(&dtoVariant)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	attr, err := ac.ar.GetAttribute(c, attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var updatedDto dto.AttributeVariantDTO
	updatedDto.AttributeID = attr.ID
	if attr.Type == schema.AttrStringType {
		variantString, err := ac.ar.FindVariantString(c, dtoVariant.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		updated, err := ac.ar.UpdateVariantString(c, variantString, dtoVariant.Value)
		if err != nil {
			switch errStr := err.Error(); {
			case strings.Contains(errStr, "duplicate"):
				entErr := err.(*ent.ConstraintError)
				pqErr := entErr.Unwrap().(*pq.Error)
				dtoVariant.Err = pqErr
				c.HTML(200, "", attributeView.AttributeVariantEdit(attr.ID, dtoVariant))
				return
			default:
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		updatedDto.ID = updated.ID
		updatedDto.Value = updated.Value
	}

	if attr.Type == schema.AttrNumType {
		variantNum, err := ac.ar.FindVariantNum(c, dtoVariant.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		num, _ := strconv.Atoi(dtoVariant.Value)
		updated, err := ac.ar.UpdateVariantNum(c, variantNum, num)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		updatedDto.ID = updated.ID
		updatedDto.Value = strconv.Itoa(updated.Value)
	}

	c.HTML(200, "", attributeView.VariantRow(attr.ID, updatedDto))
}

func (ac AttributeController) AttributeVariantAdd(c *gin.Context) {
	var dtoVariant dto.AttributeVariantDTO
	err := c.ShouldBind(&dtoVariant)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&dtoVariant)
	if err != nil {
		dtoVariant.FillErrors(err)
		c.HTML(http.StatusBadRequest, "", attributeView.VariantAddRow(dtoVariant.AttributeID, ""))
		return
	}

	attr, err := ac.ar.GetAttribute(c, dtoVariant.AttributeID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = ac.ar.CreateVariant(c, attr, dtoVariant.Value)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, "/attributes/"+"1"+"/edit")
}

func (ac AttributeController) AttributeVariantDelete(c *gin.Context) {
	aId := c.Param("id")
	attributeId, err := strconv.Atoi(aId)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	vId := c.Param("vid")
	variantId, err := strconv.Atoi(vId)
	if err != nil {
		log.Println("can't convert")
		c.Redirect(302, "/")
		return
	}

	attr, err := ac.ar.GetAttribute(c, attributeId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	variant, _ := ac.ar.GetVariant(c, attr, variantId)

	ac.ar.DeleteVariant(c, *variant)
	c.JSON(200, gin.H{})
}

// ///////// ATRIBUTE VALUES /////////////////////////////////////////
func (ac AttributeController) GetAllAttributeValues(c *gin.Context) {
	c.JSON(200, ac.ar.GetAllAttributeValues(c))
}
