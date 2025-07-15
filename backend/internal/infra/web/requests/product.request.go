package requests

type GetProductDetailsRequest struct {
	Uuid string `uri:"uuid"`
}

type GetProductInteractionRequest struct {
	Uuid string `uri:"uuid"`
}

type DeleteProductRequest struct {
	Uuid string `uri:"uuid"`
}

type ListUserProductsQueryParams struct {
	Page string `form:"page"`
	// Category string `form:"category"`
}

type AddProductInteractionRequest struct {
	Uuid  string `json:"uuid"`
	Emoji string `json:"emoji"`
}
