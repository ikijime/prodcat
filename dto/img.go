package dto

type Image struct {
	Header string
	Body   string
	Name   string
}

func (i *Image) ToString() string {
	return "data:" + i.Header + ";base64," + i.Body
}
