package catalog

import (
	"database/sql"
	"github.com/apoberez/pc/core"
)

type Category struct {
	UID   string
	Title string
	Image string
	Slug  string
}

func (c *Category) JsonApiData() *core.JsonApiData {
	return &core.JsonApiData{
		Id: c.UID,
		Type: "category",
		Attributes: map[string]interface{}{
			"title": c.Title,
			"image": c.Image,
			"slug": c.Slug,
		},
	}
}

type CategoryRepository struct {
	db *sql.DB
}

func (r *CategoryRepository) FindAll() []*Category {
	categories := []*Category{}
	rowSql := "SELECT c.uid, c.title, c.slug, c.image FROM category c;"
	rows, err := r.db.Query(rowSql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var uid, title, slug, image string

		err = rows.Scan(&uid, &title, &slug, &image)
		categories = append(categories, &Category{uid, title, image, slug})
	}

	return categories
}
