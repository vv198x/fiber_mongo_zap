package main

import (
	"encoding/json"
	"fiber_mongo_zap/configs"
	"fiber_mongo_zap/logger"
	"fiber_mongo_zap/routes"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

func main() {
	//Передаем логлевел
	logger.L = logger.Zap(zapcore.DebugLevel)
	defer logger.L.Sync()

	// embedded fs
	engine := html.NewFileSystem(http.Dir("./templates"), ".htm")
	engine.Debug(true) //Для проверки как загрузились

	app := fiber.New(fiber.Config{
		ServerHeader: "Service",
		AppName:      "Service v0.0.1",
		Views:        engine,
	})

	//Логирование файбер в зап
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.L,
	}))

	configs.ConnectDB()

	//Добавляем индекс в MongoDB для быстрого выполнения запроса
	configs.CreateIndex()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/dec", http.StatusPermanentRedirect)
	})

	routes.DeclarationRoute(app)

	//Записать доступные машруты в режиме дебаг
	data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
	logger.Debug(string(data))

	logger.Error(app.Listen(os.Getenv("SERVICE_PORT")))
}
