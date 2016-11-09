package catalog

import (
	"net/http"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/apoberez/pc/core"
)

func GetRoutes() []*core.Route {
	return []*core.Route{
		{
			Path: "/categories",
			Method: "GET",
			Action: CategoriesList,
		}, {
			Path: "/products",
			Method: "GET",
			Action: ProductsList,
		},
	}
}

func CategoriesList(w http.ResponseWriter, r *http.Request) {
	container := r.Context().Value("container").(*core.Container)
	repo := CategoryRepository{container.DB}
	response := core.JsonResponse{}
	categories := repo.FindAll()
	d := []*core.JsonApiData{}
	for _, i := range categories {
		d = append(d, i.JsonApiData())
	}
	response.Data = d
	responseStr, err := json.Marshal(response)
	if nil != err {
		panic("Json error")
	}
	w.Write(responseStr)
}

func ProductsList(w http.ResponseWriter, r *http.Request) {
	container := r.Context().Value("container").(*core.Container)
	repo := ProductRepository{container.DB}
	result, err := json.Marshal(repo.FindAll())
	if nil != err {
		panic("Json error")
	}
	w.Write(result)
}
