package dto

type DTO interface {
	isValid() bool
	FillErrors()
}
