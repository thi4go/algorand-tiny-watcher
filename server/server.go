package server

import (
	"algorand-tiny-watcher/config"
	"algorand-tiny-watcher/pkg/watcher"
	"algorand-tiny-watcher/server/api"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TinyWatcherServer struct {
	watcher *watcher.Watcher
	config  *config.WatcherConfig
}

func New() (*TinyWatcherServer, error) {
	config := config.NewWatcherConfig()
	watcher := watcher.NewWatcher(config)

	return &TinyWatcherServer{
		watcher: watcher,
		config:  config,
	}, nil
}

func (tws *TinyWatcherServer) Start() {
	initLogger()

	go tws.watcher.Run()

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var method string

			// Stop health-check log spam
			if c.Request().RequestURI == "/health" && v.Status == 200 {
				return nil
			}

			method = fmt.Sprintf("[%s: %d]", c.Request().Method, v.Status)
			logger.Info().Msg(fmt.Sprintf("%s %dms : %s",
				method, v.Latency.Milliseconds(), v.URI))

			b, err := io.ReadAll(c.Request().Body)
			if err != nil {
				logger.Err(err)
			}

			if len(b) > 0 {
				logger.Info().Msg(string(b))
			}

			if v.Error != nil {
				logger.Err(err)
			}

			return nil
		},
	}))

	e.GET(api.RouteWatchState, tws.handleWatchState)
	e.POST(api.RouteWatchAddress, tws.handleWatchAddress)

	e.Logger.Fatal(e.Start(":" + tws.config.Port))
}
