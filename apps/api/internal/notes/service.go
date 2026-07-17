package notes

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

func (s *Service) GetNote(ctx context.Context, userID, channelID int64) (sqlc.Note, error) {
	return s.queries.GetNote(ctx, sqlc.GetNoteParams{
		UserID:    userID,
		ChannelID: channelID,
	})
}

func (s *Service) UpsertNote(ctx context.Context, userID, channelID int64, body string) (sqlc.Note, error) {
	return s.queries.UpsertNote(ctx, sqlc.UpsertNoteParams{
		UserID:    userID,
		ChannelID: channelID,
		Body:      body,
	})
}

func (s *Service) DeleteNote(ctx context.Context, userID, channelID int64) error {
	return s.queries.DeleteNote(ctx, sqlc.DeleteNoteParams{
		UserID:    userID,
		ChannelID: channelID,
	})
}
