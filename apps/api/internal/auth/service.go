package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	"github.com/eugene/tubedex/internal/config"
	"github.com/eugene/tubedex/internal/db/sqlc"
)

type Service struct {
	cfg      *config.Config
	queries  *sqlc.Queries
	oauthCfg *oauth2.Config
}

func NewService(cfg *config.Config, queries *sqlc.Queries) *Service {
	oauthCfg := &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleSecret,
		RedirectURL:  cfg.GoogleRedirect,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			youtube.YoutubeReadonlyScope,
		},
		Endpoint: google.Endpoint,
	}

	return &Service{
		cfg:      cfg,
		queries:  queries,
		oauthCfg: oauthCfg,
	}
}

func (s *Service) GoogleLoginURL() string {
	return s.oauthCfg.AuthCodeURL(generateState(), oauth2.AccessTypeOffline)
}

func (s *Service) HandleCallback(ctx context.Context, code string) (string, error) {
	token, err := s.oauthCfg.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("exchange code: %w", err)
	}

	client := s.oauthCfg.Client(ctx, token)
	googleUser, err := fetchGoogleUser(ctx, client)
	if err != nil {
		return "", fmt.Errorf("fetch google user: %w", err)
	}

	user, err := s.queries.UpsertUser(ctx, sqlc.UpsertUserParams{
		GoogleID:  googleUser.ID,
		Email:     googleUser.Email,
		Name:      googleUser.Name,
		AvatarUrl: googleUser.Picture,
	})
	if err != nil {
		return "", fmt.Errorf("upsert user: %w", err)
	}

	sessionID, err := generateSessionID()
	if err != nil {
		return "", fmt.Errorf("generate session: %w", err)
	}

	var expiresAt pgtype.Timestamptz
	expiresAt.Time = time.Now().Add(s.cfg.SessionMaxAge)
	expiresAt.Valid = true

	_, err = s.queries.CreateSession(ctx, sqlc.CreateSessionParams{
		ID:        sessionID,
		UserID:    user.ID,
		Data:      []byte(token.AccessToken),
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", fmt.Errorf("create session: %w", err)
	}

	return sessionID, nil
}

func (s *Service) DeleteSession(ctx context.Context, sessionID string) error {
	return s.queries.DeleteSession(ctx, sessionID)
}

type GoogleUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func fetchGoogleUser(ctx context.Context, client *http.Client) (*GoogleUser, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func generateState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func generateSessionID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
