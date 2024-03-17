package mux

import (
	"net/http"

	"filmlib/internal/config"
	"filmlib/internal/handlers"
	"filmlib/internal/logging"
	"filmlib/internal/mux/middleware"

	_ "filmlib/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupMux(handlers *handlers.Handlers, config *config.Config, logger *logging.LogrusLogger) *http.Handler {
	mux := http.NewServeMux()

	baseUrl := config.API.BaseUrl

	actorsBaseUrl := baseUrl + "actors/"
	actorsSpecificUrl := actorsBaseUrl + "{id}/"

	moviesBaseUrl := baseUrl + "movies/"
	moviesSpecificUrl := moviesBaseUrl + "{id}/"

	searchBaseUrl := baseUrl + "search/"
	searchMovieUrl := searchBaseUrl + "movies/"

	swaggerUrl := "/swagger/*"

	mux.Handle("POST "+actorsBaseUrl, middleware.Stack(
		wrapHandleFunc(handlers.ActorHandler.CreateActor),
		middleware.Auth,
	))
	mux.HandleFunc("GET "+actorsBaseUrl, handlers.ActorHandler.GetActors)
	mux.Handle("GET "+actorsSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.ActorHandler.ReadActor),
		middleware.ExtractID,
	))
	mux.Handle("PATCH "+actorsSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.ActorHandler.UpdateActor),
		middleware.ExtractID,
		middleware.Auth,
	))
	mux.Handle("DELETE "+actorsSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.ActorHandler.DeleteActor),
		middleware.ExtractID,
		middleware.Auth,
	))

	mux.Handle("GET "+moviesBaseUrl, middleware.Stack(
		wrapHandleFunc(handlers.MovieHandler.GetMovies),
		middleware.ExtractSortParams,
	))
	mux.Handle("POST "+moviesBaseUrl, middleware.Stack(
		wrapHandleFunc(handlers.MovieHandler.CreateMovie),
		middleware.Auth,
	))
	mux.Handle("GET "+moviesSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.MovieHandler.ReadMovie),
		middleware.ExtractID,
	))
	mux.Handle("PATCH "+moviesSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.MovieHandler.UpdateMovie),
		middleware.ExtractID,
		middleware.Auth,
	))
	mux.Handle("DELETE "+moviesSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.MovieHandler.DeleteMovie),
		middleware.ExtractID,
		middleware.Auth,
	))

	mux.Handle("GET "+searchMovieUrl, middleware.Stack(
		wrapHandleFunc(handlers.SearchHandler.SearchMovies),
		middleware.ExtractQuery,
		middleware.Auth,
	))

	mwStack := middleware.Stack(mux,
		middleware.JsonHeader, middleware.NewLogger(logger),
		middleware.RequestID, middleware.PanicRecovery,
	)

	mux.HandleFunc("GET "+swaggerUrl, httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
	))

	return &mwStack
}

func wrapHandleFunc(hf http.HandlerFunc) http.Handler {
	return hf
}
