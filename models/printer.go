package models

import (
	"database/sql"
)

type Printer struct {
	ID                   uint
	Name                 string         `db:"name"`
	Description          sql.NullString `db:"description"`
	PagePerMinute        float32        `db:"page_per_minute"`
	BrandID              uint           `db:"brand_id"`
	PrintingTechnologyID uint           `db:"printing_technology_id"`
	FunctionTypeID       uint           `db:"function_type_id"`
	PrintSizeID          uint           `db:"print_size_id"`
	PrintResolutionX     uint           `db:"print_resolution_x"`
	PrintResolutionY     uint           `db:"print_resolution_y"`
	Size                 sql.NullString `db:"size"`
	Weight               float32        `db:"weight"`
	AdditionalInfo       sql.NullString `db:"additional_info"`
	Amount               uint           `db:"amount"`
	Price                uint           `db:"price"`
}

type Brand struct {
	ID          uint
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}

type PrintingTechnology struct {
	ID          uint
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}

type FunctionType struct {
	ID          uint
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}

type PrintSize struct {
	ID   uint
	Name string `db:"name"`
}

type PrintResolution struct {
	X uint
	Y uint
}
