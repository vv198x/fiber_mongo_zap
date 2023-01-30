package routes

import (
	. "fiber_mongo_zap/controllers"
	"github.com/gofiber/fiber/v2"
)

func DeclarationRoute(app *fiber.App) {
	dec := app.Group("/dec")
	dec.Get("/", GetAllDeclaration)
	dec.Post("/add", CheckXml, GetXMLFiles)

}
