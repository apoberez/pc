package catalog

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type ProductAttributes map[string]interface{}

type Product struct {
	UID           string
	Title         string
	Specification ProductAttributes
	Manufacturer  string
}

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) FindAll() []*Product {
	products := []*Product{}
	rowSql := `SELECT p.uid, p.title, p.specification, m.title
			FROM product p
			LEFT JOIN manufacturer m ON m.uid = p.manufacturer;`
	rows, err := r.db.Query(rowSql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var uid, title, specStr, manufacturer string

		err = rows.Scan(&uid, &title, &specStr, &manufacturer)
		if err != nil {
			panic("Sql error")
		}
		spec := make(ProductAttributes)
		err = json.Unmarshal([]byte(specStr), &spec)
		if err != nil {
			fmt.Print(err)
			panic("Invalid json")
		}
		products = append(products, &Product{uid, title, spec, manufacturer})
	}

	return products
}