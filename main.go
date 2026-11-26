package main

import (
	"log"
	"API/store"
	"API/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
	AllowOrigins: []string{"http://localhost:5173"},
	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders: []string{"Content-Type"},
	}))
	st := store.NewStore("data.json")
	h := handlers.NewHandler(st)

	app.Get("/artists", h.ListArtists)
	app.Get("/artists/:id", h.GetArtist)
	app.Post("/artists", h.CreateArtist)
	app.Put("/artists/:id", h.UpdateArtist)
	app.Delete("/artists/:id", h.DeleteArtist)

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":8080"))
}