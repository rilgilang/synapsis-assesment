package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"synapsis-challenge/bootstrap"
	"synapsis-challenge/config/yaml"
	"synapsis-challenge/internal/api/routes"
	"synapsis-challenge/internal/middlewares/jwt"
	"synapsis-challenge/internal/repositories"
	"synapsis-challenge/internal/service"
	"synapsis-challenge/migrations"
)

func main() {

	//environment vairable that im using is app.yaml
	//no specific reason only I usually use yaml like this because it looks cleaner and readable
	cfg, err := yaml.NewConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf(`read cfg yaml got error : %v`, err))
	}

	//made connection to our database
	db, err := bootstrap.DatabaseConnection(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf(`db connection error got : %v`, err))
	}

	fmt.Println("Database connection success!")

	migrations.AutoMigration(db)

	if err != nil {
		log.Fatal(fmt.Sprintf(`error auto migrate got : %v`, err))
	}

	fmt.Println("Migration success!")

	//repository to interact with database
	userRepo := repositories.NewUserRepo(db)
	productRepo := repositories.NewProductRepo(db)
	cartRepo := repositories.NewCartRepo(db)
	cartProductRepo := repositories.NewCartProductRepo(db)
	orderRepo := repositories.NewOrdersRepo(db)
	transactionRepo := repositories.NewTransactionsRepo(db)

	//middleware for validation authorization
	middleware := jwt.NewAuthMiddleware(userRepo, cfg)

	//setup the service / bussines logic that we gonna used later
	userService := service.NewAuthService(middleware, userRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(productRepo, cartRepo, cartProductRepo)
	checkoutService := service.NewCheckoutService(transactionRepo, cartRepo, orderRepo, cartProductRepo)
	transactionsService := service.NewTransactionsService(transactionRepo)

	app := fiber.New()
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Mwehehe"))
	})

	//setup the routers
	api := app.Group("/api")
	routes.LoginRouter(api, userService)
	routes.ProductRouter(api, productService)
	routes.CartRouter(api, middleware, cartService)
	routes.CheckoutRouter(api, middleware, checkoutService)
	routes.TransactionsRouter(api, middleware, transactionsService)

	log.Fatal(app.Listen(fmt.Sprintf(`:%s`, cfg.App.Port)))
}
