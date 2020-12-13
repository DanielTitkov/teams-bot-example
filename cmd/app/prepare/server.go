package prepare

// import (
// 	"github.com/DanielTitkov/teams-bot-example/internal/api/handler"
// 	"github.com/DanielTitkov/teams-bot-example/internal/app"
// 	"github.com/DanielTitkov/teams-bot-example/internal/configs"
// 	"github.com/DanielTitkov/teams-bot-example/internal/logger"
// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// )

// func NewServer(cfg configs.Config, logger *logger.Logger, app *app.App) *echo.Echo {
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.CORS())
// 	if cfg.Env != "dev" {
// 		e.Use(middleware.Recover())
// 	}
// 	handler.NewHandler(e, cfg, logger, app)
// 	return e
// }
