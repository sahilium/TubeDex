package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/eugene/tubedex/internal/db/sqlc"
)

type contextKey string

const UserKey contextKey = "user"

func AuthRequired(queries *sqlc.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID, err := r.Cookie("session_id")
			if err != nil {
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}

			session, err := queries.GetSessionByID(r.Context(), sessionID.Value)
			if err != nil {
				http.Error(w, `{"error":"invalid session"}`, http.StatusUnauthorized)
				return
			}

			if session.ExpiresAt.Valid && session.ExpiresAt.Time.Before(time.Now()) {
				http.Error(w, `{"error":"session expired"}`, http.StatusUnauthorized)
				return
			}

			user, err := queries.GetUserByID(r.Context(), session.UserID)
			if err != nil {
				http.Error(w, `{"error":"user not found"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, &user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUser(r *http.Request) *sqlc.User {
	user, ok := r.Context().Value(UserKey).(*sqlc.User)
	if !ok {
		return nil
	}
	return user
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote", r.RemoteAddr).
			Msg("request")
		next.ServeHTTP(w, r)
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Interface("panic", err).Msg("recovered from panic")
				http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
