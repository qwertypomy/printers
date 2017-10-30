package db

import (
	"database/sql"
	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
)

func Populate(config models.Config) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	userDao := factoryDao.GetUserDaoInterface()
	printerDao := factoryDao.GetPrinterDaoInterface()

	user := models.User{Name: "Vasya", Email: "vasya@mail.ru", Password: "12345"}
	err := userDao.CreateUser(&user)
	utils.FatalError(err)
	b1 := models.Brand{Name: "HP", Description: sql.NullString{"Hewlett-Packard — одна из крупнейших американских компаний в сфере информационных технологий, существовавшая в период 1939—2015 годов, поставщик аппаратного и программного обеспечения для организаций и индивидуальных потребителей.", true}}
	err = printerDao.CreateBrand(&b1)
	utils.FatalError(err)

	pt1 := models.PrintingTechnology{Name: "Струйная печать", Description: sql.NullString{"Обладает малой скоростью печати по сравнению с лазерным принтером, но отличается высоким качеством печати полутоновых изображений, а также имеет более высокую скорость по сравнению с матричным принтером.", true}}
	err = printerDao.CreatePrintingTechnology(&pt1)
	utils.FatalError(err)
	pt2 := models.PrintingTechnology{Name: "Лазерная печать(ч/б)", Description: sql.NullString{"Подобно фотокопировальным аппаратам лазерные принтеры используют в работе процесс ксерографической печати, однако отличие состоит в том, что формирование изображения происходит путём непосредственной экспозиции (освещения) лазерным лучом фоточувствительных элементов принтера.Отпечатки, сделанные таким способом, не боятся влаги, устойчивы к истиранию и выцветанию. Качество такого изображения наиболее высокое.", true}}
	err = printerDao.CreatePrintingTechnology(&pt2)
	utils.FatalError(err)

	ps1 := models.PrintSize{Name: "A4"}
	err = printerDao.CreatePrintSize(&ps1)
	utils.FatalError(err)

	pr1 := models.PrintResolution{X: 600, Y: 600}
	err = printerDao.CreatePrintResolution(&pr1)
	utils.FatalError(err)

	ft1 := models.FunctionType{Name: "МФУ", Description: sql.NullString{"Многофункциональное устройство (МФУ) — устройство, сочетающее в себе функции принтера, сканера, факсимильного устройства, копировального модуля. Эти функции могут присутствовать в стандартной комплектации устройства или же некоторые из них могут добавляться к базовому устройству опционально.", true}}
	err = printerDao.CreateFunctionType(&ft1)
	utils.FatalError(err)

	printer1 := models.Printer{
		Name:                 "HP LaserJet Ultra M134a (G3Q66A)",
		Description:          sql.NullString{"Оцените преимущества низкой стоимости печати и высокой производительности МФУ HP LaserJet Ultra. Три черных картриджа на 2300 страниц входят в комплект поставки устройства. Создавайте документы профессионального и стабильного качества, не выходя за рамки бюджета.", true},
		PagePerMinute:        22,
		BrandID:              b1.ID,
		PrintingTechnologyID: pt2.ID,
		FunctionTypeID:       ft1.ID,
		PrintSizeID:          ps1.ID,
		PrintResolutionX:     pr1.X,
		PrintResolutionY:     pr1.Y,
		Amount:               4,
		Price:                5181,
	}
	err = printerDao.CreatePrinter(&printer1)
	utils.FatalError(err)
}
