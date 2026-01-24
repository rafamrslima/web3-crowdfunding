package dtos

type CategoryDto struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}
