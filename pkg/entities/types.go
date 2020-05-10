package entities

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PriceInCents uint64 `json:"priceInCents"` // Oversimplified: we could use a type for the Amount or something similar
	Inventory    uint32 `json:"inventory"`
}

type Catalog struct {
	Products []Product `json:"products"`
}

type CatalogFilterOptions struct {
	ProductNameContains string
}

type Account struct {
	ID    string
	Email string
}

type Order struct {
	AccountID string
	ProductID string

	TotalInCents uint64 // Oversimplified: no shipping costs, no several items, no product freeze, etc
}
