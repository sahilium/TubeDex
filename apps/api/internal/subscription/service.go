package subscription

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/eugene/tubedex/internal/db/sqlc"
)

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) ListSubscriptions(ctx context.Context, userID int64, limit, offset int32) ([]sqlc.ListSubscriptionsByNameRow, error) {
	subs, err := s.queries.ListSubscriptionsByName(ctx, sqlc.ListSubscriptionsByNameParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list subscriptions: %w", err)
	}
	return subs, nil
}

func (s *Service) Subscribe(ctx context.Context, userID, channelID int64) error {
	var subscribedAt pgtype.Timestamptz
	subscribedAt.Time = time.Now()
	subscribedAt.Valid = true

	_, err := s.queries.UpsertSubscription(ctx, sqlc.UpsertSubscriptionParams{
		UserID:       userID,
		ChannelID:    channelID,
		SubscribedAt: subscribedAt,
	})
	return err
}

func (s *Service) Unsubscribe(ctx context.Context, userID, channelID int64) error {
	return s.queries.DeleteSubscription(ctx, sqlc.DeleteSubscriptionParams{
		UserID:    userID,
		ChannelID: channelID,
	})
}
