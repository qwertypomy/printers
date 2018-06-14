package models

type Printer struct {
	ID                   string  `bson:"_id"`
	Name                 string  `db:"name"`
	Description          string  `db:"description"`
	PagePerMinute        float32 `db:"page_per_minute"`
	BrandID              string  `db:"brand_id"`
	PrintingTechnologyID string  `db:"printing_technology_id"`
	FunctionTypeID       string  `db:"function_type_id"`
	PrintSizeID          string  `db:"print_size_id"`
	PrintResolutionX     uint    `db:"print_resolution_x"`
	PrintResolutionY     uint    `db:"print_resolution_y"`
	Size                 string  `db:"size"`
	Weight               float32 `db:"weight"`
	AdditionalInfo       string  `db:"additional_info"`
	Amount               uint    `db:"amount"`
	Price                uint    `db:"price"`
}

type Brand struct {
	ID          string `bson:"_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type PrintingTechnology struct {
	ID          string `bson:"_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type FunctionType struct {
	ID          string `bson:"_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type PrintSize struct {
	ID   string `bson:"_id"`
	Name string `db:"name"`
}

type PrintResolution struct {
	X uint
	Y uint
}
