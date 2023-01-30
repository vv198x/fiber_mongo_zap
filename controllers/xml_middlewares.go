package controllers

import (
	"context"
	"encoding/xml"
	"fiber_mongo_zap/logger"
	"fiber_mongo_zap/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Проверяю входящие xml
func CheckXml(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["files"]
		if len(files) == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Выберете файл")
		}

		pos := make([]models.PurchaseOrder, len(files)-1)
		for _, file := range files {
			var po models.PurchaseOrder
			f, _ := file.Open()
			defer f.Close()
			err := xml.NewDecoder(f).Decode(&po)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Неверный формат xml")
			}
			pos = append(pos, po)
			logger.Info(fmt.Sprintf("Purchase order number: %s, order date: %s\n", po.PurchaseOrderNumber, po.OrderDate))
		}

		c.Locals("dec", &pos)
	}

	//Валидация
	return func(c *fiber.Ctx) error {
		pos := c.Locals("dec").(*[]models.PurchaseOrder)
		validator := validator.New()
		for _, po := range *pos {
			if err := validator.Struct(po); err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Неверный формат xml")
			}
		}

		//Записываю в базу
		return func(c *fiber.Ctx, pos *[]models.PurchaseOrder) error {
			//добавить ожидание
			mongo := models.Mongo{
				ID:   primitive.NewObjectID(),
				Decs: *pos,
			}
			result, err := declarationCollection.InsertOne(context.Background(), mongo)
			if err != nil {
				logger.Error("")
			}
			logger.Debug("Document added: ", result.InsertedID)

			return c.Next()
		}(c, pos)
	}(c)
}
