package youtube

import (
	"context"
	"fmt"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"github.com/rs/zerolog/log"
)

type Channel struct {
	ID                string
	Name              string
	Handle            string
	Description       string
	Avatar            string
	Banner            string
	SubscriberCount   int64
	UploadsPlaylistID string
}

type Video struct {
	ID          string
	Title       string
	Description string
	PublishedAt time.Time
	Thumbnail   string
}

type Service struct {
	apiKey string
}

func NewService(apiKey string) *Service {
	return &Service{apiKey: apiKey}
}

func (s *Service) FetchSubscriptions(ctx context.Context, accessToken string) ([]Channel, error) {
	token := &oauth2.Token{AccessToken: accessToken}
	ts := oauth2.StaticTokenSource(token)

	service, err := youtube.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, fmt.Errorf("create youtube service: %w", err)
	}

	var channels []Channel
	pageToken := ""

	for {
		call := service.Subscriptions.List([]string{"snippet"}).
			Mine(true).
			MaxResults(50)

		if pageToken != "" {
			call = call.PageToken(pageToken)
		}

		resp, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("list subscriptions: %w", err)
		}

		for _, item := range resp.Items {
			if item.Snippet == nil || item.Snippet.ResourceId == nil {
				continue
			}

			channelID := item.Snippet.ResourceId.ChannelId
			channel, err := s.FetchChannel(ctx, channelID)
			if err != nil {
				log.Warn().Err(err).Str("channel_id", channelID).Msg("failed to fetch channel details")
				continue
			}
			channels = append(channels, *channel)
		}

		if resp.NextPageToken == "" {
			break
		}
		pageToken = resp.NextPageToken
	}

	return channels, nil
}

func (s *Service) FetchChannel(ctx context.Context, channelID string) (*Channel, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		return nil, fmt.Errorf("create youtube service: %w", err)
	}

	call := service.Channels.List([]string{"snippet", "statistics", "contentDetails"}).
		Id(channelID)

	resp, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("list channels: %w", err)
	}

	if len(resp.Items) == 0 {
		return nil, fmt.Errorf("channel not found: %s", channelID)
	}

	item := resp.Items[0]
	channel := &Channel{
		ID:                item.Id,
		Name:              item.Snippet.Title,
		Description:       item.Snippet.Description,
		Avatar:            item.Snippet.Thumbnails.Default.Url,
		SubscriberCount:   int64(item.Statistics.SubscriberCount),
		UploadsPlaylistID: item.ContentDetails.RelatedPlaylists.Uploads,
	}

	if item.Snippet.CustomUrl != "" {
		channel.Handle = strings.TrimPrefix(item.Snippet.CustomUrl, "@")
	}

	return channel, nil
}

func (s *Service) FetchLatestVideos(ctx context.Context, uploadsPlaylistID string, maxResults int64) ([]Video, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		return nil, fmt.Errorf("create youtube service: %w", err)
	}

	call := service.PlaylistItems.List([]string{"snippet", "contentDetails"}).
		PlaylistId(uploadsPlaylistID).
		MaxResults(maxResults)

	resp, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("list playlist items: %w", err)
	}

	var videos []Video
	for _, item := range resp.Items {
		if item.Snippet == nil {
			continue
		}

		publishedAt, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		thumbnail := ""
		if item.Snippet.Thumbnails != nil && item.Snippet.Thumbnails.High != nil {
			thumbnail = item.Snippet.Thumbnails.High.Url
		}

		videos = append(videos, Video{
			ID:          item.ContentDetails.VideoId,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishedAt: publishedAt,
			Thumbnail:   thumbnail,
		})
	}

	return videos, nil
}
