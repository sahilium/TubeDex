package channel

import (
	"context"
	"fmt"

	"github.com/eugene/tubedex/internal/db/sqlc"
	"github.com/eugene/tubedex/internal/youtube"
)

type Service struct {
	queries  *sqlc.Queries
	ytClient *youtube.Service
}

func NewService(queries *sqlc.Queries, ytClient *youtube.Service) *Service {
	return &Service{
		queries:  queries,
		ytClient: ytClient,
	}
}

type ChannelWithDetails struct {
	Channel
	Rating      *float64 `json:"rating"`
	Note        *string  `json:"note"`
	Collections []int64  `json:"collection_ids"`
}

type Channel struct {
	ID                int64  `json:"id"`
	YoutubeChannelID  string `json:"youtube_channel_id"`
	Name              string `json:"name"`
	Handle            string `json:"handle"`
	Description       string `json:"description"`
	Avatar            string `json:"avatar"`
	Banner            string `json:"banner"`
	SubscriberCount   int64  `json:"subscriber_count"`
	UploadsPlaylistID string `json:"uploads_playlist_id"`
}

func (s *Service) ListSubscribedChannels(ctx context.Context, userID int64, sort string, limit, offset int32) ([]ChannelWithDetails, error) {
	var err error
	var subs []sqlc.ListSubscriptionsByNameRow

	if sort == "recent" {
		recent, recentErr := s.queries.ListSubscriptionsByRecent(ctx, sqlc.ListSubscriptionsByRecentParams{
			UserID: userID,
			Limit:  limit,
			Offset: offset,
		})
		err = recentErr
		// Convert to the same type
		for _, r := range recent {
			subs = append(subs, sqlc.ListSubscriptionsByNameRow(r))
		}
	} else {
		subs, err = s.queries.ListSubscriptionsByName(ctx, sqlc.ListSubscriptionsByNameParams{
			UserID: userID,
			Limit:  limit,
			Offset: offset,
		})
	}

	if err != nil {
		return nil, fmt.Errorf("list subscriptions: %w", err)
	}

	var results []ChannelWithDetails
	for _, sub := range subs {
		results = append(results, ChannelWithDetails{
			Channel: Channel{
				ID:                sub.ChannelID,
				YoutubeChannelID:  sub.YoutubeChannelID,
				Name:              sub.ChannelName,
				Handle:            sub.ChannelHandle,
				Avatar:            sub.ChannelAvatar,
				Banner:            sub.ChannelBanner,
				Description:       sub.ChannelDescription,
				SubscriberCount:   sub.ChannelSubscriberCount,
				UploadsPlaylistID: sub.UploadsPlaylistID,
			},
		})
	}

	return results, nil
}

func (s *Service) GetChannelDetails(ctx context.Context, userID, channelID int64) (*ChannelWithDetails, error) {
	ch, err := s.queries.GetChannelByID(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("get channel: %w", err)
	}

	rating, err := s.queries.GetRating(ctx, sqlc.GetRatingParams{
		UserID:    userID,
		ChannelID: channelID,
	})
	var ratingVal *float64
	if err == nil {
		v := float64(rating.Rating)
		ratingVal = &v
	}

	note, err := s.queries.GetNote(ctx, sqlc.GetNoteParams{
		UserID:    userID,
		ChannelID: channelID,
	})
	var noteVal *string
	if err == nil {
		n := note.Body
		noteVal = &n
	}

	collections, err := s.queries.ListCollectionsByChannel(ctx, sqlc.ListCollectionsByChannelParams{
		UserID:    userID,
		ChannelID: channelID,
	})
	var collectionIDs []int64
	if err == nil {
		for _, c := range collections {
			collectionIDs = append(collectionIDs, c.ID)
		}
	}

	return &ChannelWithDetails{
		Channel: Channel{
			ID:                ch.ID,
			YoutubeChannelID:  ch.YoutubeChannelID,
			Name:              ch.Name,
			Handle:            ch.Handle,
			Description:       ch.Description,
			Avatar:            ch.Avatar,
			Banner:            ch.Banner,
			SubscriberCount:   ch.SubscriberCount,
			UploadsPlaylistID: ch.UploadsPlaylistID,
		},
		Rating:      ratingVal,
		Note:        noteVal,
		Collections: collectionIDs,
	}, nil
}

func (s *Service) UpsertChannel(ctx context.Context, ch *youtube.Channel) (*sqlc.Channel, error) {
	result, err := s.queries.UpsertChannel(ctx, sqlc.UpsertChannelParams{
		YoutubeChannelID:  ch.ID,
		Name:              ch.Name,
		Handle:            ch.Handle,
		Description:       ch.Description,
		Avatar:            ch.Avatar,
		Banner:            ch.Banner,
		SubscriberCount:   ch.SubscriberCount,
		UploadsPlaylistID: ch.UploadsPlaylistID,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
