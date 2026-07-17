package collection

import (
	"context"
	"fmt"

	"github.com/eugene/tubedex/internal/db/sqlc"
)

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) ListCollections(ctx context.Context, userID int64) ([]sqlc.Collection, error) {
	collections, err := s.queries.ListCollectionsByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list collections: %w", err)
	}
	return collections, nil
}

func (s *Service) CreateCollection(ctx context.Context, userID int64, name, icon, color string) (sqlc.Collection, error) {
	if name == "" {
		return sqlc.Collection{}, fmt.Errorf("name is required")
	}

	collection, err := s.queries.CreateCollection(ctx, sqlc.CreateCollectionParams{
		UserID: userID,
		Name:   name,
		Icon:   icon,
		Color:  color,
	})
	if err != nil {
		return sqlc.Collection{}, fmt.Errorf("create collection: %w", err)
	}
	return collection, nil
}

func (s *Service) UpdateCollection(ctx context.Context, id, userID int64, name, icon, color string) (sqlc.Collection, error) {
	collection, err := s.queries.UpdateCollection(ctx, sqlc.UpdateCollectionParams{
		ID:     id,
		UserID: userID,
		Name:   name,
		Icon:   icon,
		Color:  color,
	})
	if err != nil {
		return sqlc.Collection{}, fmt.Errorf("update collection: %w", err)
	}
	return collection, nil
}

func (s *Service) DeleteCollection(ctx context.Context, id, userID int64) error {
	return s.queries.DeleteCollection(ctx, sqlc.DeleteCollectionParams{
		ID:     id,
		UserID: userID,
	})
}

func (s *Service) ListCollectionChannels(ctx context.Context, collectionID, userID int64) ([]sqlc.Channel, error) {
	channels, err := s.queries.ListCollectionChannels(ctx, collectionID)
	if err != nil {
		return nil, fmt.Errorf("list collection channels: %w", err)
	}
	return channels, nil
}

func (s *Service) AddChannelToCollection(ctx context.Context, collectionID, userID, channelID int64) error {
	_, err := s.queries.GetCollectionByID(ctx, sqlc.GetCollectionByIDParams{
		ID:     collectionID,
		UserID: userID,
	})
	if err != nil {
		return fmt.Errorf("collection not found")
	}

	return s.queries.AddChannelToCollection(ctx, sqlc.AddChannelToCollectionParams{
		CollectionID: collectionID,
		ChannelID:    channelID,
	})
}

func (s *Service) RemoveChannelFromCollection(ctx context.Context, collectionID, channelID int64) error {
	return s.queries.RemoveChannelFromCollection(ctx, sqlc.RemoveChannelFromCollectionParams{
		CollectionID: collectionID,
		ChannelID:    channelID,
	})
}
