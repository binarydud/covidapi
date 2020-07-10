package router

import (
	"net/http"
	"time"

	"github.com/binarydud/covidapi/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func NewRouter(log zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	// r.Use(mw.LogCtxHandler(log.Logger))
	r.Use(hlog.NewHandler(log))
	r.Use(hlog.RemoteAddrHandler("ip"))
	r.Use(hlog.UserAgentHandler("user_agent"))
	r.Use(hlog.RefererHandler("referer"))
	r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("host", r.Host).
			Str("uri", r.RequestURI).
			Str("url", r.URL.String()).
			Str("path", r.URL.Path).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	// r.Use(mw.ZeroLogger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(handlers.DBMiddleware)
	r.Get("/health", handlers.HealthHandler)
	r.Post("/slack", handlers.CommandHandler)
	r.Get("/authorize", handlers.AuthorizeHandler)
	r.Get("/us/current", handlers.USHandler)
	r.Get("/us/daily", handlers.USHistoricalHandler)
	r.Get("/states/{state}/current", handlers.StateHandler)
	r.Get("/states/{state}/daily", handlers.StateHistoricalHandler)
	return r
}
