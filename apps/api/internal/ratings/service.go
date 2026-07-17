package ratings

import (
	"context"

	"github.com/eugene/tubedex/internal/db/sqlc"
)

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) GetRating(ctx context.Context, userID, channelID int64) (sqlc.Rating, error) {
	return s.queries.GetRating(ctx, sqlc.GetRatingParams{
		UserID:    userID,
		ChannelID: channelID,
	})
}

func (s *Service) UpsertRating(ctx context.Context, userID, channelID int64, rating int16) (sqlc.Rating, error) {
	return s.queries.UpsertRating(ctx, sqlc.UpsertRatingParams{
		UserID:    userID,
		ChannelID: channelID,
		Rating:    rating,
	})
}

func (s *Service) DeleteRating(ctx context.Context, userID, channelID int64) error {
	return s.queries.DeleteRating(ctx, sqlc.DeleteRatingParams{
		UserID:    userID,
		ChannelID: channelID,
	})
}
