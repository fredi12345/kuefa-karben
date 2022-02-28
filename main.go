package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/fredi12345/kuefa-karben/src/auth"
	"github.com/fredi12345/kuefa-karben/src/extensions"
	"github.com/fredi12345/kuefa-karben/src/rest"
	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/fredi12345/kuefa-karben/src/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gorilla/mux"
)

func main() {
	viper.SetConfigFile(path.Join("config", "config.yaml"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	viper.SetEnvPrefix("KUEFA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	logger := initializeLogger()
	db, err := storage.NewPostgresBackend(logger)
	if err != nil {
		logger.Fatal("could not create postgres backend", zap.Error(err))
	}

	err = db.Migrate(context.Background())
	if err != nil {
		logger.Fatal("could not execute migration", zap.Error(err))
	}

	user := viper.GetString("default.user")
	password := viper.GetString("default.password")
	if user != "" && password != "" {
		err = db.CreateUser(user, password)
		if err != nil {
			logger.Fatal(fmt.Sprintf("failed to create default user %s", user), zap.Error(err))
		}
	}

	server, err := web.NewServer(db, "cookies.key")
	if err != nil {
		logger.Fatal("could not create server", zap.Error(err))
	}

	handler := createHandler(server)
	logger.Info("serving legacy application on http://localhost:8000")
	go func() {
		if err := http.ListenAndServe(":8000", handler); err != nil {
			logger.Fatal("ListenAndServe failed", zap.Error(err))
		}
	}()

	restServer := rest.NewServer(db, logger)
	logger.Info(fmt.Sprintf("server up and running: http://localhost:%d", viper.GetInt("web.port")))
	e, err := createEchoHandler(restServer, logger)
	if err != nil {
		logger.Fatal("could not create echo handler", zap.Error(err))
	}

	if err := e.Start(fmt.Sprintf(":%d", viper.GetInt("web.port"))); err != nil {
		logger.Fatal("ListenAndServe failed", zap.Error(err))
	}
}

func initializeLogger() *zap.Logger {
	var l *zap.Logger
	var err error

	if viper.GetBool("debug") {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		l, err = cfg.Build()
	} else {
		l, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("could not initialize logger: %v", err)
	}

	return l.Named("kuefa")
}

func createEchoHandler(server *rest.Server, logger *zap.Logger) (*echo.Echo, error) {
	authenticationProvider, err := auth.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("could not create authentication provider: %w", err)
	}

	e := echo.New()
	e.Binder = extensions.BinderChain(&echo.DefaultBinder{}, extensions.MultipartFileBinder())
	e.HideBanner = true
	e.HidePort = true

	e.Use(extensions.LoggerWithConfig(extensions.LoggerConfig{Logger: logger}))
	e.Use(middleware.CORS())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  viper.GetString("web.root"),
		Index: "index.html",
		HTML5: true,
	}))

	public := e.Group("public")
	public.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       viper.GetString("web.storage"),
		IgnoreBase: true,
	}))

	authAPI := e.Group("auth")
	authAPI.GET("/login", authenticationProvider.Login)
	authAPI.GET("/callback", authenticationProvider.Callback)
	authAPI.GET("/logout", authenticationProvider.Logout)

	api := e.Group("api")
	api.POST("/images", extensions.AutoBind(logger, server.UploadImage))
	api.POST("/events", extensions.AutoBind(logger, server.CreateEvent))
	return e, nil
}

func createHandler(server *web.Server) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.Dir("resources/public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public", blockDirectoryListing(fs, server)))
	r.Handle("/robots.txt", fs)

	// redirect from incomplete urls
	r.Handle("/event", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all/0", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/gallery", http.RedirectHandler("/gallery/1", http.StatusSeeOther))
	r.Handle("/gallery/0", http.RedirectHandler("/gallery/1", http.StatusSeeOther))

	// http get methods
	r.HandleFunc("/", server.WithSession(server.HandleError(server.Index))).Methods(http.MethodGet)
	r.HandleFunc("/event/create", server.WithSession(server.HandleError(server.NeedsAuthentication(server.CreateEventPage)))).Methods(http.MethodGet)
	r.HandleFunc("/event/edit", server.WithSession(server.HandleError(server.NeedsAuthentication(server.EditEventPage)))).Methods(http.MethodGet)
	r.HandleFunc("/event/{id}", server.WithSession(server.HandleError(server.EventDetail))).Methods(http.MethodGet)
	r.HandleFunc("/event/all/{page:[0-9]+}", server.WithSession(server.HandleError(server.AllEvents))).Methods(http.MethodGet)
	r.HandleFunc("/impressum", server.WithSession(server.HandleError(server.Impressum))).Methods(http.MethodGet)
	r.HandleFunc("/gallery/{page:[0-9]*}", server.WithSession(server.HandleError(server.Gallery))).Methods(http.MethodGet)

	// http post methods
	r.HandleFunc("/event/add", server.WithSession(server.HandleError(server.NeedsAuthentication(server.AddEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/event/edit", server.WithSession(server.HandleError(server.NeedsAuthentication(server.EditEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/event/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/participant/add", server.WithSession(server.HandleError(server.AddParticipant))).Methods(http.MethodPost)
	r.HandleFunc("/participant/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteParticipant)))).Methods(http.MethodPost)
	r.HandleFunc("/comment/add", server.WithSession(server.HandleError(server.AddComment))).Methods(http.MethodPost)
	r.HandleFunc("/comment/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteComment)))).Methods(http.MethodPost)
	r.HandleFunc("/image/add", server.WithSession(server.HandleError(server.NeedsAuthentication(server.AddImage)))).Methods(http.MethodPost)
	r.HandleFunc("/image/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteImage)))).Methods(http.MethodPost)
	r.HandleFunc("/user/login", server.WithSession(server.HandleError(server.Login))).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", server.WithSession(server.HandleError(server.Logout))).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(server.NotFound)
	return r
}

func blockDirectoryListing(next http.Handler, server *web.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			server.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
