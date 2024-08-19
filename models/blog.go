package models

type Blog struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	IsPublished bool   `json:"isPublished"`
}

var blogs = []Blog{
	{
		ID:          1,
		Title:       "My first blog",
		Description: "This is my first blog",
		Body:        "This is the body of my first blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          2,
		Title:       "My second blog",
		Description: "This is my second blog",
		Body:        "This is the body of my second blog",
		Author:      "Jane Doe",
		IsPublished: false,
	},
	{
		ID:          3,
		Title:       "My third blog",
		Description: "This is my third blog",
		Body:        "This is the body of my third blog",
		Author:      "John Doe",
		IsPublished: true,
	},
	{
		ID:          4,
		Title:       "My fourth blog",
		Description: "This is my fourth blog",
		Body:        "This is the body of my fourth blog",
		Author:      "Jane Doe",
		IsPublished: true,
	},
}
