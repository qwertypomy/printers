package db

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
	"os"
	"strconv"
)

var brandN, technologyN, functionTypeN, printSizeN, resolutionXN, resolutionYN, nameN, descriptionN,
	additionalInfoN, amountN, weightN, pagePerMinuteN, priceN, sizeN int

func Populate(config models.Config) (err error) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	printerDao := factoryDao.GetPrinterDaoInterface()

	arr, err := ReadCSV()
	if err != nil {
		return err
	}

	brands, technologies, functionTypes, sizes, resolutions := ArrToObjects(arr)

	for _, brand := range brands {
		err = printerDao.CreateBrand(&brand)
		utils.FatalError(err)
	}
	for _, technology := range technologies {
		err = printerDao.CreatePrintingTechnology(&technology)
		utils.FatalError(err)
	}
	for _, functionType := range functionTypes {
		err = printerDao.CreateFunctionType(&functionType)
		utils.FatalError(err)
	}
	for _, size := range sizes {
		err = printerDao.CreatePrintSize(&size)
		utils.FatalError(err)
	}
	for _, resolution := range resolutions {
		err = printerDao.CreatePrintResolution(&resolution)
		utils.FatalError(err)
	}

	printers := ArrToPrinters(arr, config)

	for _, printer := range printers {
		err = printerDao.CreatePrinter(&printer)
		utils.FatalError(err)
	}

	return
}

func ReadCSV() (arr [][]string, err error) {
	f, err := os.Open("./printers.csv")
	if err != nil {
		return
	}
	r := csv.NewReader(bufio.NewReader(f))
	arr, err = r.ReadAll()
	if err != nil {
		return
	}

	for i, name := range arr[0] {
		switch name {
		case "Brand":
			brandN = i
		case "PrintingTechnology":
			technologyN = i
		case "FunctionType":
			functionTypeN = i
		case "PrintSize":
			printSizeN = i
		case "PrintResolutionX":
			resolutionXN = i
		case "PrintResolutionY":
			resolutionYN = i
		case "Name":
			nameN = i
		case "Description":
			descriptionN = i
		case "AdditionalInfo":
			additionalInfoN = i
		case "Amount":
			amountN = i
		case "Weight":
			weightN = i
		case "PagePerMinute":
			pagePerMinuteN = i
		case "Price":
			priceN = i
		case "Size":
			sizeN = i
		}
	}

	return
}

func ArrToObjects(arr [][]string) (brands []models.Brand,
	technologies []models.PrintingTechnology,
	functionTypes []models.FunctionType,
	sizes []models.PrintSize,
	resolutions []models.PrintResolution) {

	// Fill brands, technologies, functionTypes, sizes, resolutions
	for _, printer := range arr[1:] {
		unique := true

		temp := printer[brandN]
		for _, v := range brands {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			brands = append(brands, models.Brand{Name: temp})
		}
		unique = true

		temp = printer[technologyN]
		for _, v := range technologies {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			technologies = append(technologies, models.PrintingTechnology{Name: temp})
		}
		unique = true

		temp = printer[functionTypeN]
		for _, v := range functionTypes {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			functionTypes = append(functionTypes, models.FunctionType{Name: temp})
		}
		unique = true

		temp = printer[printSizeN]
		for _, v := range sizes {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			sizes = append(sizes, models.PrintSize{Name: temp})
		}
		unique = true
		tempx, err := strconv.Atoi(printer[resolutionXN])
		utils.FatalError(err)
		tempX := uint(tempx)
		tempy, err := strconv.Atoi(printer[resolutionYN])
		utils.FatalError(err)
		tempY := uint(tempy)
		for _, v := range resolutions {
			if v.X == tempX && v.Y == tempY {
				unique = false
				break
			}
		}
		if unique {
			resolutions = append(resolutions, models.PrintResolution{X: tempX, Y: tempY})
		}
		unique = true
	}
	return
}

func ArrToPrinters(arr [][]string, config models.Config) (printers []models.Printer) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	printerDao := factoryDao.GetPrinterDaoInterface()

	brands := printerDao.BrandList()
	technologies := printerDao.PrintingTechnologyList()
	functionTypes := printerDao.FunctionTypeList()
	printSizes := printerDao.PrintSizeList()
	resolutions := printerDao.PrintResolutionList()

	for _, printer := range arr[1:] {
		amount, err := strconv.Atoi(printer[amountN])
		utils.FatalError(err)
		weight, err := strconv.ParseFloat(printer[weightN], 32)
		utils.FatalError(err)
		ppm, err := strconv.ParseFloat(printer[pagePerMinuteN], 32)
		utils.FatalError(err)
		price, err := strconv.Atoi(printer[priceN])
		utils.FatalError(err)

		var brand, technology, functionType, printSize, resolutionX, resolutionY uint
		for _, v := range brands {
			if v.Name == printer[brandN] {
				brand = v.ID
			}
		}
		for _, v := range technologies {
			if v.Name == printer[technologyN] {
				technology = v.ID
			}
		}
		for _, v := range functionTypes {
			if v.Name == printer[functionTypeN] {
				functionType = v.ID
			}
		}
		for _, v := range printSizes {
			if v.Name == printer[printSizeN] {
				printSize = v.ID
			}
		}
		for _, v := range resolutions {
			tempx, err := strconv.Atoi(printer[resolutionXN])
			utils.FatalError(err)
			tempX := uint(tempx)
			tempy, err := strconv.Atoi(printer[resolutionYN])
			utils.FatalError(err)
			tempY := uint(tempy)
			if v.X == tempX && v.Y == tempY {
				resolutionX = v.X
				resolutionY = v.Y
			}
		}

		printers = append(printers, models.Printer{
			Name:                 printer[nameN],
			Description:          sql.NullString{printer[descriptionN], true},
			AdditionalInfo:       sql.NullString{printer[additionalInfoN], true},
			Amount:               uint(amount),
			Weight:               float32(weight),
			PagePerMinute:        float32(ppm),
			Price:                uint(price),
			Size:                 sql.NullString{printer[sizeN], true},
			BrandID:              brand,
			PrintingTechnologyID: technology,
			FunctionTypeID:       functionType,
			PrintSizeID:          printSize,
			PrintResolutionX:     resolutionX,
			PrintResolutionY:     resolutionY,
		})
	}
	return
}
