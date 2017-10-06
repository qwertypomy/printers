package models

type Printer struct {
	ID                   uint
	Name                 string             `gorm:"size:45;unique;not null"`
	Description          string             `gorm:"size:2000"`
	PagePerMinute        uint               `gorm:"not null"`
	BrandID              uint               `gorm:"not null"`
	PrintingTechnologyID uint               `gorm:"not null"`
	FunctionTypeID       uint               `gorm:"not null"`
	PrintSizeID          uint               `gorm:"not null"`
	PrintResolutionX     uint               `gorm:"not null"`
	PrintResolutionY     uint               `gorm:"not null"`
	ConnectivityTypes    []ConnectivityType `gorm:"many2many:printer_connectivity_type"`
	Size                 string             `gorm:"size:16"`
	Weight               string             `gorm:"size:16"`
	AdditionalInfo       string             `gorm:"size:500"`
	Number               uint               `gorm:"size:8;not null"`
	Price                uint               `gorm:"size:8;not null"`
}

type Brand struct {
	ID          uint
	Name        string `gorm:"size:16;unique;not null"`
	Description string `gorm:"size:2000"`
	Printers    []Printer
}

type PrintingTechnology struct {
	ID          uint
	Name        string `gorm:"size:16;unique;not null"`
	Description string `gorm:"size:2000"`
	Printers    []Printer
}

type FunctionType struct {
	ID          uint
	Name        string `gorm:"size:45;unique;not null"`
	Description string `gorm:"size:2000"`
	Printers    []Printer
}

type PrintSize struct {
	ID       uint
	Name     string `gorm:"size:3;unique;not null"`
	Printers []Printer
}

type PrintResolution struct {
	X        uint `gorm:"not null"`
	Y        uint `gorm:"not null"`
	Printers []Printer
}

type ConnectivityType struct {
	ID       uint
	Name     string    `gorm:"size:16;unique;not null"`
	Printers []Printer `gorm:"many2many:printer_connectivity_type;"`
}
