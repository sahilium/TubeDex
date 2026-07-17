package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/eugene/tubedex/internal/middleware"
	"github.com/eugene/tubedex/internal/services"
)

func setupRoutes(c *services.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(middleware.Recovery)
	r.Use(middleware.Logging)
	r.Use(middleware.CORS(middleware.AllowedOrigins(c.Config.FrontendURL)))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/login", c.AuthHandler.Login)
			r.Get("/callback", c.AuthHandler.Callback)
			r.Post("/logout", c.AuthHandler.Logout)
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthRequired(c.Queries))

			r.Get("/me", c.UserHandler.GetMe)
			r.Patch("/me", c.UserHandler.UpdateMe)

			r.Get("/subscriptions", c.SubscriptionHandler.List)
			r.Delete("/subscriptions", c.SubscriptionHandler.Delete)

			r.Get("/channels", c.ChannelHandler.List)
			r.Get("/channels/{id}", c.ChannelHandler.Get)

			r.Get("/collections", c.CollectionHandler.List)
			r.Post("/collections", c.CollectionHandler.Create)
			r.Patch("/collections/{id}", c.CollectionHandler.Update)
			r.Delete("/collections/{id}", c.CollectionHandler.Delete)
			r.Get("/collections/{id}/channels", c.CollectionHandler.ListChannels)
			r.Post("/collections/{id}/channels", c.CollectionHandler.AddChannel)
			r.Delete("/collections/{id}/channels", c.CollectionHandler.RemoveChannel)

			r.Post("/sync", c.SyncHandler.Sync)
			r.Get("/sync/status", c.SyncHandler.Status)

			r.Get("/search", c.SearchHandler.Search)

			r.Get("/notes", c.NotesHandler.Get)
			r.Put("/notes", c.NotesHandler.Upsert)
			r.Delete("/notes", c.NotesHandler.Delete)

			r.Get("/ratings", c.RatingsHandler.Get)
			r.Put("/ratings", c.RatingsHandler.Upsert)
			r.Delete("/ratings", c.RatingsHandler.Delete)
		})
	})

	return r
}
