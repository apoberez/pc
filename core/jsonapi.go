package core

type JsonApiDataAware interface {
	JsonApiData () (*JsonApiData)
}

type JsonResponse struct {
	Data interface{} `json:"data"`
}

type JsonApiData struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes map[string]interface{} `json:"attributes"`
}
