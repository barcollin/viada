package server

import (
	"flag"
	"os"
	"xaria/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

var (
	addr = flag.String("addr", ":"+os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	engine := html.New("./views", "html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)

	app.Static("/", "./assets")

	if *cert != "" {
		return app.ListenTLS(*addr, *cert, *key)
	}

	// go func() {
	// 	room.Peers.DispatchKeyFrame()
	// }()

	return app.Listen(*addr)

}

// func DispatchKeyFrame() {
// 	for range time.NewTicker(time.Second * 3).C {
// 		for _, room := range w.Rooms{
// 			room.Reers.DispatchKeyFrame()
// 		}
// 	}
// }
