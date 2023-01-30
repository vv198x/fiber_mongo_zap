package controllers

import (
	"context"
	"fiber_mongo_zap/configs"
	"fiber_mongo_zap/logger"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var declarationCollection *mongo.Collection = configs.GetCollection(configs.DB, "declaration")

func GetAllDeclaration(c *fiber.Ctx) error {
	//Запрос умножает количество items на цену во всех документах
	pipeline := []bson.M{
		bson.M{"$unwind": "$decs"},
		bson.M{"$unwind": "$decs.items.items"},
		bson.M{
			"$group": bson.M{
				//id обязательно
				"_id": nil,
				//Названия поля будет в мапе
				"total": bson.M{"$sum": bson.M{
					"$multiply": []interface{}{
						bson.M{"$toDouble": "$decs.items.items.price"},
						"$decs.items.items.quantity",
					},
				}},
			},
		},
	}

	// Выполняем запрос
	cursor, err := declarationCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		logger.Error("mongo err", err)
	}
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			logger.Error("mongo err", err)
		}

		logger.Debug(result)
	}

	if err := cursor.Err(); err != nil {
		logger.Error("mongo err", err)
	}

	// Закрываем курсор
	cursor.Close(context.TODO())
	return c.Render("input_file_form", nil)
}

func GetXMLFiles(c *fiber.Ctx) error {
	return c.Redirect("/dec")
}
