package sync

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/eugene/tubedex/internal/channel"
	"github.com/eugene/tubedex/internal/config"
	"github.com/eugene/tubedex/internal/db/sqlc"
	"github.com/eugene/tubedex/internal/youtube"
)

type Service struct {
	cfg        *config.Config
	queries    *sqlc.Queries
	pool       *pgxpool.Pool
	channelSvc *channel.Service
	ytClient   *youtube.Service
}

func NewService(cfg *config.Config, queries *sqlc.Queries, pool *pgxpool.Pool, channelSvc *channel.Service, ytClient *youtube.Service) *Service {
	return &Service{
		cfg:        cfg,
		queries:    queries,
		pool:       pool,
		channelSvc: channelSvc,
		ytClient:   ytClient,
	}
}

func (s *Service) StartSync(ctx context.Context, userID int64) (*sqlc.SyncJob, error) {
	job, err := s.queries.CreateSyncJob(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("create sync job: %w", err)
	}

	go s.runSync(context.Background(), userID, job.ID, s.pool)

	return &job, nil
}

func (s *Service) GetLatestSync(ctx context.Context, userID int64) (*sqlc.SyncJob, error) {
	job, err := s.queries.GetLatestSyncJob(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get latest sync: %w", err)
	}
	return &job, nil
}

func (s *Service) runSync(ctx context.Context, userID, jobID int64, pool *pgxpool.Pool) {
	queries := sqlc.New(pool)

	log.Info().Int64("user_id", userID).Int64("job_id", jobID).Msg("starting sync")

	queries.UpdateSyncJobStatus(ctx, sqlc.UpdateSyncJobStatusParams{
		ID:     jobID,
		Status: "running",
		Error:  "",
	})

	session, err := queries.GetSessionByUserID(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get session for sync")
		queries.UpdateSyncJobStatus(ctx, sqlc.UpdateSyncJobStatusParams{
			ID:     jobID,
			Status: "failed",
			Error:  err.Error(),
		})
		return
	}

	accessToken := string(session.Data)

	channels, err := s.ytClient.FetchSubscriptions(ctx, accessToken)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch subscriptions")
		queries.UpdateSyncJobStatus(ctx, sqlc.UpdateSyncJobStatusParams{
			ID:     jobID,
			Status: "failed",
			Error:  err.Error(),
		})
		return
	}

	var subscribedAt pgtype.Timestamptz
	subscribedAt.Time = time.Now()
	subscribedAt.Valid = true

	for _, ytCh := range channels {
		dbCh, err := queries.UpsertChannel(ctx, sqlc.UpsertChannelParams{
			YoutubeChannelID:  ytCh.ID,
			Name:              ytCh.Name,
			Handle:            ytCh.Handle,
			Description:       ytCh.Description,
			Avatar:            ytCh.Avatar,
			Banner:            ytCh.Banner,
			SubscriberCount:   ytCh.SubscriberCount,
			UploadsPlaylistID: ytCh.UploadsPlaylistID,
		})
		if err != nil {
			log.Error().Err(err).Str("channel", ytCh.Name).Msg("failed to upsert channel")
			continue
		}

		queries.UpsertSubscription(ctx, sqlc.UpsertSubscriptionParams{
			UserID:       userID,
			ChannelID:    dbCh.ID,
			SubscribedAt: subscribedAt,
		})
	}

	queries.UpdateSyncJobStatus(ctx, sqlc.UpdateSyncJobStatusParams{
		ID:     jobID,
		Status: "completed",
		Error:  "",
	})

	log.Info().Int64("user_id", userID).Int("channels", len(channels)).Msg("sync completed")
}
