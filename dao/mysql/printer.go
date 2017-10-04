package mysql

import (
	"github.com/qwertypomy/printers/models"
)

type PrinterDaoImpl struct {
}

/////////////////// function createIdValue for abstract struct IdValue ///////////////////
func (pdi PrinterDaoImpl) createIdValue(tableName string, value string) (model models.IdValue, err error) {
	model = models.IdValue{
		Value: value,
	}
	_, err = Db.Exec("insert into "+tableName+" (value) values (?)", value)
	if err != nil {
		return
	}
	err = Db.QueryRow("select LAST_INSERT_ID()").Scan(&model.Id)
	return
}

/////////////////// function createIdNameDescription for abstract struct IdNameDescription ///////////////////
func (pdi PrinterDaoImpl) createIdNameDescription(tableName string, name string, description string) (model models.IdNameDescription, err error) {
	model = models.IdNameDescription{
		Name:        name,
		Description: description,
	}
	_, err = Db.Exec("insert into "+tableName+" (name, description) values (?, ?)", name, description)
	if err != nil {
		return
	}
	err = Db.QueryRow("select LAST_INSERT_ID()").Scan(&model.Id)
	return
}

/////////////////// functions for not abstract structs based on IdValue ///////////////////

// PrinterPrintSize functions
func (pdi PrinterDaoImpl) CreatePrinterPrintSize(value string) (printSize models.PrinterPrintSize, err error) {
	printSize.IdValue, err = pdi.createIdValue("print_size", value)
	return
}

func (pdi PrinterDaoImpl) PrinterPrintSizeList() (printSizeList []models.PrinterPrintSize, err error) {
	rows, err := Db.Query("select id, name from print_size")
	if err != nil {
		return
	}
	for rows.Next() {
		var printSize models.PrinterPrintSize
		rows.Scan(&printSize.Id, &printSize.Value)
		printSizeList = append(printSizeList, printSize)
	}
	return
}

// PrinterConnectivityType functions
func (pdi PrinterDaoImpl) CreatePrinterConnectivityType(value string) (connectivityType models.PrinterConnectivityType, err error) {
	connectivityType.IdValue, err = pdi.createIdValue("connectivity_type", value)
	return
}

func (pdi PrinterDaoImpl) PrinterConnectivityTypeList() (connectivityTypeList []models.PrinterConnectivityType, err error) {
	rows, err := Db.Query("select id, name from connectivity_type")
	if err != nil {
		return
	}
	for rows.Next() {
		var connectivityType models.PrinterConnectivityType
		rows.Scan(&connectivityType.Id, &connectivityType.Value)
		connectivityTypeList = append(connectivityTypeList, connectivityType)
	}
	return
}

/////////////////// functions for not abstract structs based on IdNameDescription ///////////////////

// PrinterBrand functions
func (pdi PrinterDaoImpl) CreatePrinterBrand(name string, description string) (brand models.PrinterBrand, err error) {
	brand.IdNameDescription, err = pdi.createIdNameDescription("brand", name, description)
	return
}

func (pdi PrinterDaoImpl) PrinterBrandList() (brandList []models.PrinterBrand, err error) {
	rows, err := Db.Query("select id, name, description from brand")
	if err != nil {
		return
	}
	for rows.Next() {
		var brand models.PrinterBrand
		rows.Scan(&brand.Id, &brand.Name, &brand.Description)
		brandList = append(brandList, brand)
	}
	return
}

// PrinterPrintingTechnology functions
func (pdi PrinterDaoImpl) CreatePrinterPrintingTechnology(name string, description string) (printingTechnology models.PrinterPrintingTechnology, err error) {
	printingTechnology.IdNameDescription, err = pdi.createIdNameDescription("printing_technology", name, description)
	return
}

func (pdi PrinterDaoImpl) PrinterPrintingTechnologyList() (printingTechnologyList []models.PrinterPrintingTechnology, err error) {
	rows, err := Db.Query("select id, name, description from printing_technology")
	if err != nil {
		return
	}
	for rows.Next() {
		var printingTechnology models.PrinterPrintingTechnology
		rows.Scan(&printingTechnology.Id, &printingTechnology.Name, &printingTechnology.Description)
		printingTechnologyList = append(printingTechnologyList, printingTechnology)
	}
	return
}

// PrinterFunctionType functions
func (pdi PrinterDaoImpl) CreatePrinterFunctionType(name string, description string) (functionType models.PrinterFunctionType, err error) {
	functionType.IdNameDescription, err = pdi.createIdNameDescription("function_type", name, description)
	return
}

func (pdi PrinterDaoImpl) PrinterFunctionTypeList() (functionTypeList []models.PrinterFunctionType, err error) {
	rows, err := Db.Query("select id, name, description from function_type")
	if err != nil {
		return
	}
	for rows.Next() {
		var functionType models.PrinterFunctionType
		rows.Scan(&functionType.Id, &functionType.Name, &functionType.Description)
		functionTypeList = append(functionTypeList, functionType)
	}
	return
}

/////////////////// functions for PrinterPrintResolution ///////////////////

func (pdi PrinterDaoImpl) CreatePrinterPrintResolution(x, y int) (printResolution models.PrinterPrintResolution, err error) {
	printResolution = models.PrinterPrintResolution{
		X: x,
		Y: y,
	}
	_, err = Db.Exec("insert into print_resolution (x, y) values (?, ?)", x, y)
	if err != nil {
		return
	}
	return
}

func (pdi PrinterDaoImpl) PrinterPrintResolutionList() (printResolutionList []models.PrinterPrintResolution, err error) {
	rows, err := Db.Query("select x, y from print_resolution")
	if err != nil {
		return
	}
	for rows.Next() {
		var printResolution models.PrinterPrintResolution
		rows.Scan(&printResolution.X, &printResolution.Y)
		printResolutionList = append(printResolutionList, printResolution)
	}
	return
}

/////////////////// functions for Printer ///////////////////
func (pdi PrinterDaoImpl) CreatePrinter(
	name string,
	description string,
	pagePerMinute int,
	brand *models.PrinterBrand,
	printingTechnology *models.PrinterPrintingTechnology,
	functionType *models.PrinterFunctionType,
	printSize *models.PrinterPrintSize,
	printResolution *models.PrinterPrintResolution,
	connectivityType []models.PrinterConnectivityType,
	size string,
	weight string,
	additionalInfo string,
	number int,
	price int,
) (printer models.Printer, err error) {
	printer = models.Printer{
		Name:               name,
		Description:        description,
		PagePerMinute:      pagePerMinute,
		Brand:              brand,
		PrintingTechnology: printingTechnology,
		FunctionType:       functionType,
		PrintSize:          printSize,
		PrintResolution:    printResolution,
		ConnectivityType:   connectivityType,
		Size:               size,
		Weight:             weight,
		AdditionalInfo:     additionalInfo,
		Number:             number,
		Price:              price,
	}
	_, err = Db.Exec("insert into printer (x, y) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		printer.Name,
		printer.Description,
		printer.PagePerMinute,
		printer.Brand.Id,
		printer.PrintingTechnology.Id,
		printer.FunctionType.Id,
		printer.PrintSize.Id,
		printer.Size,
		printer.Weight,
		printer.AdditionalInfo,
		printer.Number,
		printer.Price,
		printer.PrintResolution.X,
		printer.PrintResolution.Y,
	)
	if err != nil {
		return
	}
	err = Db.QueryRow("select LAST_INSERT_ID()").Scan(&printer.Id)
	return
}

func (pdi PrinterDaoImpl) PrinterList() (printerList []models.Printer, err error) {
	rows, err := Db.Query("select id, name, pages_per_minute, brand_id, printing_technology_id, function_type_id, print_size_id, print_resolution_id, size, weight, description, additional_info, number, price, print_resolution_x, print_resolution_y from print_resolution")
	if err != nil {
		return
	}
	for rows.Next() {
		var printer models.Printer
		var brand_id, printing_technology_id, function_type_id, print_size_id, print_resolution_id, print_resolution_x, print_resolution_y int
		rows.Scan(&printer.Id, &printer.Name, &printer.PagePerMinute, &brand_id, &printing_technology_id, &function_type_id, &print_size_id, &print_resolution_id, printer.Size, printer.Weight, printer.Description, printer.AdditionalInfo, printer.Number, printer.Price, print_resolution_x, print_resolution_y)
		printerList = append(printerList, printer)

	}
	return
}
