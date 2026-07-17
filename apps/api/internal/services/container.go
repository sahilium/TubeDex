package services

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/eugene/tubedex/internal/auth"
	"github.com/eugene/tubedex/internal/channel"
	"github.com/eugene/tubedex/internal/collection"
	"github.com/eugene/tubedex/internal/config"
	"github.com/eugene/tubedex/internal/db/sqlc"
	"github.com/eugene/tubedex/internal/notes"
	"github.com/eugene/tubedex/internal/ratings"
	"github.com/eugene/tubedex/internal/search"
	"github.com/eugene/tubedex/internal/subscription"
	"github.com/eugene/tubedex/internal/sync"
	"github.com/eugene/tubedex/internal/user"
	"github.com/eugene/tubedex/internal/youtube"
)

type Container struct {
	Config       *config.Config
	DB           *pgxpool.Pool
	Queries      *sqlc.Queries

	AuthService  *auth.Service
	UserService  *user.Service
	YoutubeService *youtube.Service
	ChannelService *channel.Service
	SubscriptionService *subscription.Service
	CollectionService *collection.Service
	SyncService  *sync.Service
	SearchService *search.Service
	NotesService *notes.Service
	RatingsService *ratings.Service

	AuthHandler *auth.Handler
	UserHandler *user.Handler
	ChannelHandler *channel.Handler
	SubscriptionHandler *subscription.Handler
	CollectionHandler *collection.Handler
	SyncHandler *sync.Handler
	SearchHandler *search.Handler
	NotesHandler *notes.Handler
	RatingsHandler *ratings.Handler
}

func NewContainer(cfg *config.Config, pool *pgxpool.Pool) *Container {
	q := sqlc.New(pool)

	ytService := youtube.NewService(cfg.YouTubeAPIKey)
	authService := auth.NewService(cfg, q)
	userService := user.NewService(q)
	channelService := channel.NewService(q, ytService)
	subService := subscription.NewService(q)
	colService := collection.NewService(q)
	syncService := sync.NewService(cfg, q, pool, channelService, ytService)
	searchService := search.NewService(q)
	notesService := notes.NewService(q)
	ratingsService := ratings.NewService(q)

	return &Container{
		Config:       cfg,
		DB:           pool,
		Queries:      q,

		AuthService:  authService,
		UserService:  userService,
		YoutubeService: ytService,
		ChannelService: channelService,
		SubscriptionService: subService,
		CollectionService: colService,
		SyncService:  syncService,
		SearchService: searchService,
		NotesService: notesService,
		RatingsService: ratingsService,

		AuthHandler: auth.NewHandler(authService),
		UserHandler: user.NewHandler(userService),
		ChannelHandler: channel.NewHandler(channelService),
		SubscriptionHandler: subscription.NewHandler(subService),
		CollectionHandler: collection.NewHandler(colService),
		SyncHandler: sync.NewHandler(syncService),
		SearchHandler: search.NewHandler(searchService),
		NotesHandler: notes.NewHandler(notesService),
		RatingsHandler: ratings.NewHandler(ratingsService),
	}
}
