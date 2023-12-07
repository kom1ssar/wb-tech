package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"tech-wb/api/order"
	"tech-wb/internal/config"
	"tech-wb/pkg/client/postgresql"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.connectDB,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {

	handler := a.initHttpRoutesAndMiddleware()

	a.httpServer = &http.Server{
		Addr:         a.serviceProvider.HTTPConfig().Address(),
		Handler:      handler,
		IdleTimeout:  a.serviceProvider.HTTPConfig().GetIdleTimeout(),
		ReadTimeout:  a.serviceProvider.HTTPConfig().GetTimeout(),
		WriteTimeout: a.serviceProvider.HTTPConfig().GetTimeout(),
	}

	return nil
}

func (a *App) connectDB(ctx context.Context) error {
	//db := storage.NewConnection(a.serviceProvider.DBConfig())
	//a.serviceProvider.dbService = db
	//
	//return nil
	db, err := postgresql.NewClient(ctx, a.serviceProvider.DBConfig())

	if err != nil {
		return err
	}

	a.serviceProvider.dbService = db

	return nil
}

func (a *App) initHttpRoutesAndMiddleware() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	order.RegisterRoutes(router, a.serviceProvider.OrderImpl())

	return router
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
