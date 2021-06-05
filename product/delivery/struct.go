package delivery

type addProductParam struct {
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}
