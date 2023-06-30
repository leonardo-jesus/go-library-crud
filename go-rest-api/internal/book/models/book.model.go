package models

type Book struct {
	Id              int      `json:"id" validate:"required,numeric,min=0"`
	Name            string   `json:"name" validate:"max=255"`
	Edition         int      `json:"edition" validate:"numeric"`
	PublicationYear int      `json:"publicationYear" validate:"numeric"`
	Authors         []string `json:"authors"`
}

type CreateBookSchema struct {
	Name            string `json:"name" validate:"max=255"`
	Edition         int    `json:"edition" validate:"numeric"`
	PublicationYear int    `json:"publicationYear" validate:"numeric"`
	Authors         []int  `json:"authors"`
}

type UpdateBookSchema struct {
	Id              int     `json:"id" validate:"required,numeric,min=0"`
	Name            *string `json:"name" validate:"max=255"`
	Edition         *int    `json:"edition" validate:"numeric"`
	PublicationYear *int    `json:"publicationYear" validate:"numeric"`
	Authors         *[]int  `json:"authors"`
}
