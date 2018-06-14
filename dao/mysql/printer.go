package mysql

import (
	"github.com/qwertypomy/printers/models"
)

type PrinterDaoImpl struct {
}

// PrintSize functions
func (PrinterDaoImpl) CreatePrintSize(printSize *models.PrintSize) (err error) {
	res, err := Db.NamedExec("INSERT INTO print_size (name) VALUES (:name)", &printSize)
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			printSize.ID = string(id)
		}
	}
	return
}

func (PrinterDaoImpl) UpdatePrintSize(printSize *models.PrintSize) (err error) {
	_, err = Db.NamedExec("update print_size set name=:name,  where id =:id", &printSize)
	return
}

func (PrinterDaoImpl) DeletePrintSizeByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM print_size WHERE id=?", id)
	return
}

func (PrinterDaoImpl) GetPrintSizeByID(id string) (printSize *models.PrintSize, err error) {
	err = Db.QueryRowx("SELECT * FROM print_size WHERE id=?", id).StructScan(&printSize)
	return
}

func (PrinterDaoImpl) PrintSizeList() (printSizeList []models.PrintSize, err error) {
	rows, err := Db.Queryx("SELECT * FROM print_size")
	if err != nil {
		return
	}
	for rows.Next() {
		var printSize models.PrintSize
		err = rows.StructScan(&printSize)
		if err != nil {
			return
		}
		printSizeList = append(printSizeList, printSize)
	}
	rows.Close()
	return
}

func (PrinterDaoImpl) DeleteAllPrintSizes() (err error) {
	_, err = Db.Exec("DELETE FROM print_size")
	return
}

// Brand functions
func (PrinterDaoImpl) CreateBrand(brand *models.Brand) (err error) {
	res, err := Db.NamedExec("INSERT INTO brand (name, description) VALUES (:name, :description)", &brand)
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			brand.ID = string(id)
		}
	}
	return
}

func (PrinterDaoImpl) UpdateBrand(brand *models.Brand) (err error) {
	_, err = Db.NamedExec("update brand set name=:name, description=:description,  where id =:id", &brand)
	return
}

func (PrinterDaoImpl) DeleteBrandByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM brand WHERE id=?", id)
	return
}

func (PrinterDaoImpl) GetBrandByID(id string) (brand *models.Brand, err error) {
	err = Db.QueryRowx("SELECT * FROM brand WHERE id=?", id).StructScan(&brand)
	return
}

func (PrinterDaoImpl) BrandList() (brandList []models.Brand, err error) {
	rows, err := Db.Queryx("SELECT * FROM brand")
	if err != nil {
		return
	}
	for rows.Next() {
		var brand models.Brand
		err = rows.StructScan(&brand)
		if err != nil {
			return
		}
		brandList = append(brandList, brand)
	}
	rows.Close()
	return
}

func (PrinterDaoImpl) DeleteAllBrands() (err error) {
	_, err = Db.Exec("DELETE FROM brand")
	return
}

// PrintingTechnology functions
func (PrinterDaoImpl) CreatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error) {
	res, err := Db.NamedExec("INSERT INTO printing_technology (name, description) VALUES (:name, :description)", &printingTechnology)
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			printingTechnology.ID = string(id)
		}
	}
	return
}

func (PrinterDaoImpl) UpdatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error) {
	_, err = Db.NamedExec("update printing_technology set name=:name, description=:description,  where id =:id", &printingTechnology)
	return
}

func (PrinterDaoImpl) DeletePrintingTechnologyByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM printing_technology WHERE id=?", id)
	return
}

func (PrinterDaoImpl) GetPrintingTechnologyByID(id string) (printingTechnology *models.PrintingTechnology, err error) {
	err = Db.QueryRowx("SELECT * FROM printing_technology WHERE id=?", id).StructScan(&printingTechnology)
	return
}

func (PrinterDaoImpl) PrintingTechnologyList() (printingTechnologyList []models.PrintingTechnology, err error) {
	rows, err := Db.Queryx("SELECT * FROM printing_technology")
	if err != nil {
		return
	}
	for rows.Next() {
		var printingTechnology models.PrintingTechnology
		err = rows.StructScan(&printingTechnology)
		if err != nil {
			return
		}
		printingTechnologyList = append(printingTechnologyList, printingTechnology)
	}
	rows.Close()
	return
}

func (PrinterDaoImpl) DeleteAllPrintingTechnologies() (err error) {
	_, err = Db.Exec("DELETE FROM printing_technology")
	return
}

// FunctionType functions
func (PrinterDaoImpl) CreateFunctionType(functionType *models.FunctionType) (err error) {
	res, err := Db.NamedExec("INSERT INTO function_type (name, description) VALUES (:name, :description)", &functionType)
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			functionType.ID = string(id)
		}
	}
	return
}

func (PrinterDaoImpl) UpdateFunctionType(functionType *models.FunctionType) (err error) {
	_, err = Db.NamedExec("update function_type set name=:name, description=:description,  where id =:id", &functionType)
	return
}

func (PrinterDaoImpl) DeleteFunctionTypeByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM function_type WHERE id=?", id)
	return
}

func (PrinterDaoImpl) GetFunctionTypeByID(id string) (functionType *models.FunctionType, err error) {
	err = Db.QueryRowx("SELECT * FROM function_type WHERE id=?", id).StructScan(&functionType)
	return
}

func (PrinterDaoImpl) FunctionTypeList() (functionTypeList []models.FunctionType, err error) {
	rows, err := Db.Queryx("SELECT * FROM function_type")
	if err != nil {
		return
	}
	for rows.Next() {
		var functionType models.FunctionType
		err = rows.StructScan(&functionType)
		if err != nil {
			return
		}
		functionTypeList = append(functionTypeList, functionType)
	}
	rows.Close()
	return
}

func (PrinterDaoImpl) DeleteAllFunctionTypes() (err error) {
	_, err = Db.Exec("DELETE FROM function_type")
	return
}

// PrintResolution functions
func (PrinterDaoImpl) CreatePrintResolution(printResolution *models.PrintResolution) (err error) {
	_, err = Db.NamedExec("INSERT INTO print_resolution (x, y) VALUES (:x, :y)", &printResolution)
	return
}

func (PrinterDaoImpl) DeletePrintResolution(printResolution *models.PrintResolution) (err error) {
	_, err = Db.NamedExec("DELETE FROM print_resolution WHERE x=:x, y=:y", &printResolution)
	return
}

func (PrinterDaoImpl) GetPrintResolution(printResolution *models.PrintResolution) (err error) {
	err = Db.QueryRowx("SELECT * FROM print_resolution WHERE x=?, y=?", printResolution.X, printResolution.Y).StructScan(&printResolution)
	return
}

func (PrinterDaoImpl) PrintResolutionList() (printResolutionList []models.PrintResolution, err error) {
	rows, err := Db.Queryx("SELECT * FROM print_resolution")
	if err != nil {
		return
	}
	for rows.Next() {
		var printResolution models.PrintResolution
		err = rows.StructScan(&printResolution)
		if err != nil {
			return
		}
		printResolutionList = append(printResolutionList, printResolution)
	}
	rows.Close()
	return
}

func (PrinterDaoImpl) DeleteAllPrintResolutions() (err error) {
	_, err = Db.Exec("DELETE FROM print_resolution")
	return
}

// Printer functions
func (PrinterDaoImpl) CreatePrinter(printer *models.Printer) (err error) {
	res, err := Db.NamedExec("INSERT INTO printer (name, description, page_per_minute, brand_id, printing_technology_id, function_type_id, print_size_id, print_resolution_x, print_resolution_y, size, weight, additional_info, amount, price) VALUES (:name, :description, :page_per_minute, :brand_id, :printing_technology_id, :function_type_id, :print_size_id, :print_resolution_x, :print_resolution_y, :size, :weight, :additional_info, :amount, :price)", &printer)
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			printer.ID = string(id)
		}
	}
	return
}

func (PrinterDaoImpl) UpdatePrinter(printer *models.Printer) (err error) {
	_, err = Db.NamedExec("update printer set name=:name, description=:description, page_per_minute=:page_per_minute, brand_id=:brand_id, printing_technology_id=:printing_technology_id, function_type_id=:function_type_id, print_size_id=:print_size_id, print_resolution_x=:print_resolution_x, print_resolution_y=:print_resolution_y, size=:size, weight=:weight, additional_info=:additional_info, amount=:amount, price=:price,  where id =:id", &printer)
	return
}

func (PrinterDaoImpl) DeletePrinterByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM printer WHERE id=?", id)
	return
}

func (PrinterDaoImpl) GetPrinterByID(id string) (printer *models.Printer, err error) {
	err = Db.QueryRowx("SELECT * FROM printer WHERE id=?", id).StructScan(&printer)
	return
}

func (PrinterDaoImpl) PrinterList() (printerList []models.Printer, err error) {
	rows, err := Db.Queryx("SELECT * FROM printer")
	if err != nil {
		return
	}
	for rows.Next() {
		var printer models.Printer
		err = rows.StructScan(&printer)
		if err != nil {
			return
		}
		printerList = append(printerList, printer)
	}
	rows.Close()
	return
}

func (PrinterDaoImpl) DeleteAllPrinters() (err error) {
	_, err = Db.Exec("DELETE FROM printer")
	return
}

// All Tables functions
func (pdi PrinterDaoImpl) DeleteAllData() (err error) {
	err = pdi.DeleteAllPrintSizes()
	if err != nil {
		return
	}
	err = pdi.DeleteAllBrands()
	if err != nil {
		return
	}
	err = pdi.DeleteAllPrintingTechnologies()
	if err != nil {
		return
	}
	err = pdi.DeleteAllFunctionTypes()
	if err != nil {
		return
	}
	err = pdi.DeleteAllPrintResolutions()
	if err != nil {
		return
	}
	err = pdi.DeleteAllPrinters()
	return
}
