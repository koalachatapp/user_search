package main

import (
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/koalachatapp/usersearch/cmd/rest/handler"
	"github.com/koalachatapp/usersearch/internal/core/service"
	"github.com/koalachatapp/usersearch/internal/repository"
)

func main() {
	repo := repository.NewUsersearchRepository()
	service := service.NewUsersearchService(repo)

	// handler
	resthandler := handler.NewRestHandler(service)
	app := fiber.New(fiber.Config{
		Prefork:           true,
		CaseSensitive:     true,
		UnescapePath:      true,
		ReduceMemoryUsage: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
	})
	app.Use(recover.New())
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	app.Get("/", resthandler.Get)

	app.Listen(":3020")
}
