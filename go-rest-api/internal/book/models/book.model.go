package models

type Book struct {
	Id              int      `json:"id" validate:"required,numeric,min=0"`
	Name            string   `json:"name" validate:"max=255,min=0"`
	Edition         int      `json:"edition" validate:"numeric,min=1"`
	PublicationYear int      `json:"publicationYear" validate:"numeric"`
	Authors         []string `json:"authors"`
}

type CreateBookSchema struct {
	Name            string `json:"name" validate:"max=255"`
	Edition         int    `json:"edition" validate:"numeric,min=0"`
	PublicationYear int    `json:"publicationYear" validate:"numeric,min=1"`
	Authors         []int  `json:"authors"`
}

type UpdateBookSchema struct {
	Id              int     `json:"id" validate:"required,numeric,min=0"`
	Name            *string `json:"name" validate:"omitempty,max=255"`
	Edition         *int    `json:"edition" validate:"omitempty,numeric,min=0"`
	PublicationYear *int    `json:"publicationYear" validate:"omitempty,numeric,min=1"`
	Authors         *[]int  `json:"authors"`
}

type FilteredBookSchema struct {
	Name            *string `json:"name" validate:"omitempty,max=255"`
	Edition         *int    `json:"edition" validate:"omitempty,numeric,min=0"`
	PublicationYear *int    `json:"publicationYear" validate:"omitempty,numeric,min=1"`
	Authors         *[]int  `json:"authors"`
}
