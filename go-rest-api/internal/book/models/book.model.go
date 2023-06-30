package models

type BookDB struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Edition         int      `json:"edition"`
	PublicationYear int      `json:"publicationYear"`
	Authors         []string `json:"authors"`
}

type BookAPI struct {
	Name            string `json:"name"`
	Edition         int    `json:"edition"`
	PublicationYear int    `json:"publicationYear"`
	Authors         []int  `json:"authors"`
}
