package mongodb

import (
	"github.com/fatih/structs"
	"github.com/globalsign/mgo/bson"
	"github.com/qwertypomy/printers/models"
)

type PrinterDaoImpl struct {
}

// PrintSize functions
func (PrinterDaoImpl) CreatePrintSize(printSize *models.PrintSize) (err error) {
	printSize.ID = bson.NewObjectId().Hex()
	err = Db.C("printSize").Insert(&printSize)
	return
}

func (PrinterDaoImpl) UpdatePrintSize(printSize *models.PrintSize) (err error) {
	err = Db.C("printSize").UpdateId(printSize.ID, printSize)
	return
}

func (PrinterDaoImpl) DeletePrintSizeByID(id string) (err error) {
	err = Db.C("printSize").RemoveId(id)
	return
}

func (PrinterDaoImpl) GetPrintSizeByID(id string) (printSize *models.PrintSize, err error) {
	err = Db.C("printSize").FindId(bson.ObjectIdHex(id)).One(&printSize)
	return
}

func (PrinterDaoImpl) PrintSizeList() (printSizeList []models.PrintSize, err error) {
	err = Db.C("printSize").Find(bson.M{}).All(&printSizeList)
	return
}

func (PrinterDaoImpl) DeleteAllPrintSizes() (err error) {
	_, err = Db.C("printSize").RemoveAll(bson.M{})
	return
}

// Brand functions
func (PrinterDaoImpl) CreateBrand(brand *models.Brand) (err error) {
	brand.ID = bson.NewObjectId().Hex()
	err = Db.C("brand").Insert(&brand)
	return
}

func (PrinterDaoImpl) UpdateBrand(brand *models.Brand) (err error) {
	err = Db.C("brand").UpdateId(brand.ID, brand)
	return
}

func (PrinterDaoImpl) DeleteBrandByID(id string) (err error) {
	err = Db.C("brand").RemoveId(id)
	return
}

func (PrinterDaoImpl) GetBrandByID(id string) (brand *models.Brand, err error) {
	err = Db.C("brand").FindId(bson.ObjectIdHex(id)).One(&brand)
	return
}

func (PrinterDaoImpl) BrandList() (brandList []models.Brand, err error) {
	err = Db.C("brand").Find(bson.M{}).All(&brandList)
	return
}

func (PrinterDaoImpl) DeleteAllBrands() (err error) {
	_, err = Db.C("brand").RemoveAll(bson.M{})
	return
}

// PrintingTechnology functions
func (PrinterDaoImpl) CreatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error) {
	printingTechnology.ID = bson.NewObjectId().Hex()
	err = Db.C("printingTechnology").Insert(&printingTechnology)
	return
}

func (PrinterDaoImpl) UpdatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error) {
	err = Db.C("printingTechnology").UpdateId(printingTechnology.ID, printingTechnology)
	return
}

func (PrinterDaoImpl) DeletePrintingTechnologyByID(id string) (err error) {
	err = Db.C("printingTechnology").RemoveId(id)
	return
}

func (PrinterDaoImpl) GetPrintingTechnologyByID(id string) (printingTechnology *models.PrintingTechnology, err error) {
	err = Db.C("printingTechnology").FindId(bson.ObjectIdHex(id)).One(&printingTechnology)
	return
}

func (PrinterDaoImpl) PrintingTechnologyList() (printingTechnologyList []models.PrintingTechnology, err error) {
	err = Db.C("printingTechnology").Find(bson.M{}).All(&printingTechnologyList)
	return
}

func (PrinterDaoImpl) DeleteAllPrintingTechnologies() (err error) {
	_, err = Db.C("printingTechnology").RemoveAll(bson.M{})
	return
}

// FunctionType functions
func (PrinterDaoImpl) CreateFunctionType(functionType *models.FunctionType) (err error) {
	functionType.ID = bson.NewObjectId().Hex()
	err = Db.C("functionType").Insert(&functionType)
	return
}

func (PrinterDaoImpl) UpdateFunctionType(functionType *models.FunctionType) (err error) {
	err = Db.C("functionType").UpdateId(functionType.ID, functionType)
	return
}

func (PrinterDaoImpl) DeleteFunctionTypeByID(id string) (err error) {
	err = Db.C("functionType").RemoveId(id)
	return
}

func (PrinterDaoImpl) GetFunctionTypeByID(id string) (functionType *models.FunctionType, err error) {
	err = Db.C("functionType").FindId(bson.ObjectIdHex(id)).One(&functionType)
	return
}

func (PrinterDaoImpl) FunctionTypeList() (functionTypeList []models.FunctionType, err error) {
	err = Db.C("functionType").Find(bson.M{}).All(&functionTypeList)
	return
}

func (PrinterDaoImpl) DeleteAllFunctionTypes() (err error) {
	_, err = Db.C("functionType").RemoveAll(bson.M{})
	return
}

// PrintResolution functions
func (PrinterDaoImpl) CreatePrintResolution(printResolution *models.PrintResolution) (err error) {
	err = Db.C("printResolution").Insert(&printResolution)
	return
}

func (PrinterDaoImpl) DeletePrintResolution(printResolution *models.PrintResolution) (err error) {
	err = Db.C("printResolution").Remove(bson.M(structs.Map(printResolution)))
	return
}

func (PrinterDaoImpl) GetPrintResolution(printResolution *models.PrintResolution) (err error) {
	err = Db.C("brand").Find(bson.M(structs.Map(printResolution))).One(&printResolution)
	return
}

func (PrinterDaoImpl) PrintResolutionList() (printResolutionList []models.PrintResolution, err error) {
	err = Db.C("printResolution").Find(bson.M{}).All(&printResolutionList)
	return
}

func (PrinterDaoImpl) DeleteAllPrintResolutions() (err error) {
	_, err = Db.C("printResolution").RemoveAll(bson.M{})
	return
}

// Printer functions
func (PrinterDaoImpl) CreatePrinter(printer *models.Printer) (err error) {
	printer.ID = bson.NewObjectId().Hex()
	err = Db.C("printer").Insert(&printer)
	return
}

func (PrinterDaoImpl) UpdatePrinter(printer *models.Printer) (err error) {
	err = Db.C("printer").UpdateId(printer.ID, printer)
	return
}

func (PrinterDaoImpl) DeletePrinterByID(id string) (err error) {
	err = Db.C("printer").RemoveId(id)
	return
}

func (PrinterDaoImpl) GetPrinterByID(id string) (printer *models.Printer, err error) {
	err = Db.C("printer").FindId(bson.ObjectIdHex(id)).One(&printer)
	return
}

func (PrinterDaoImpl) PrinterList() (printerList []models.Printer, err error) {
	err = Db.C("printer").Find(bson.M{}).All(&printerList)
	return
}

func (PrinterDaoImpl) DeleteAllPrinters() (err error) {
	_, err = Db.C("printer").RemoveAll(bson.M{})
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
