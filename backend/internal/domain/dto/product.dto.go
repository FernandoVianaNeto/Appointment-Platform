package dto

type Location struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type CreateProductInputDto struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       string         `json:"price"`
	Photos      []*PhotoUpload `json:"photos"`
	Category    string         `json:"category"`
}

type GetProductDetailsInputDto struct {
	Uuid string `json:"uuid"`
}

type ListUserProductsInputDto struct {
	UserUuid string `json:"user_uuid"`
	Page     string `json:"page"`
}

type AddProductInteractionInputDto struct {
	Uuid      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	Emoji     string `json:"emoji"`
	Timestamp string `json:"timestamp"`
}
