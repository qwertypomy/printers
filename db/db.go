package db

import (
	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers/models"
)

func Populate(config models.Config) {
	factoryDao := factory.FactoryDao{config.Engine}
	userDao := factoryDao.GetUserDaoInterface()
	printerDao := factoryDao.GetPrinterDaoInterface()

	user := models.User{Name: "Vasya", Email: "vasya@mail.ru", Password: "12345"}
	userDao.CreateUser(&user)

	b1 := models.Brand{Name: "HP", Description: "Hewlett-Packard — одна из крупнейших американских компаний в сфере информационных технологий, существовавшая в период 1939—2015 годов, поставщик аппаратного и программного обеспечения для организаций и индивидуальных потребителей."}
	printerDao.Create(&b1)

	pt1 := models.PrintingTechnology{Name: "Струйная печать", Description: "Обладает малой скоростью печати по сравнению с лазерным принтером, но отличается высоким качеством печати полутоновых изображений, а также имеет более высокую скорость по сравнению с матричным принтером."}
	printerDao.Create(&pt1)
	pt2 := models.PrintingTechnology{Name: "Лазерная печать(ч/б)", Description: "Подобно фотокопировальным аппаратам лазерные принтеры используют в работе процесс ксерографической печати, однако отличие состоит в том, что формирование изображения происходит путём непосредственной экспозиции (освещения) лазерным лучом фоточувствительных элементов принтера.Отпечатки, сделанные таким способом, не боятся влаги, устойчивы к истиранию и выцветанию. Качество такого изображения наиболее высокое."}
	printerDao.Create(&pt2)

	ps1 := models.PrintSize{Name: "A4"}
	printerDao.Create(&ps1)

	pr1 := models.PrintResolution{X: 600, Y: 600}
	printerDao.Create(&pr1)

	ft1 := models.FunctionType{Name: "МФУ", Description: "Многофункциональное устройство (МФУ) — устройство, сочетающее в себе функции принтера, сканера, факсимильного устройства, копировального модуля. Эти функции могут присутствовать в стандартной комплектации устройства или же некоторые из них могут добавляться к базовому устройству опционально."}
	printerDao.Create(&ft1)

	ct1 := models.ConnectivityType{Name: "Ethernet"}
	printerDao.Create(&ct1)
	ct2 := models.ConnectivityType{Name: "Wi-Fi"}
	printerDao.Create(&ct2)

	printer1 := models.Printer{
		Name:                 "HP LaserJet Ultra M134a (G3Q66A)",
		Description:          "Оцените преимущества низкой стоимости печати и высокой производительности МФУ HP LaserJet Ultra. Три черных картриджа на 2300 страниц входят в комплект поставки устройства. Создавайте документы профессионального и стабильного качества, не выходя за рамки бюджета.",
		PagePerMinute:        22,
		BrandID:              b1.ID,
		PrintingTechnologyID: pt2.ID,
		FunctionTypeID:       ft1.ID,
		PrintSizeID:          ps1.ID,
		PrintResolutionX:     pr1.X,
		PrintResolutionY:     pr1.Y,
		ConnectivityTypes:    []models.ConnectivityType{ct1, ct2},
		Number:               4,
		Price:                5181,
	}
	printerDao.Create(&printer1)

	userDao.DeleteAllUsers()
	printerDao.DeleteAllData()
}
