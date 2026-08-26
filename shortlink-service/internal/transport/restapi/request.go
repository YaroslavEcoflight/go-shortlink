package restapi

type CreateRequest struct {
	Url string `json:"url"`
}

type DeleteRequest struct {
	Code string `json:"code"`
}
