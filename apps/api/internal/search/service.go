package search

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/eugene/tubedex/internal/db/sqlc"
)

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

type SearchResult struct {
	Type        string `json:"type"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Handle      string `json:"handle"`
	Description string `json:"description"`
}

func (s *Service) GlobalSearch(ctx context.Context, userID int64, query string, limit int32) ([]SearchResult, error) {
	var col2 pgtype.Text
	col2.String = query
	col2.Valid = true

	rows, err := s.queries.GlobalSearch(ctx, sqlc.GlobalSearchParams{
		UserID:  userID,
		Column2: col2,
		Limit:   limit,
	})
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	for _, row := range rows {
		results = append(results, SearchResult{
			Type:        row.ResultType,
			ID:          row.ResultID,
			Name:        row.ResultName,
			Avatar:      row.ResultAvatar,
			Handle:      row.ResultHandle,
			Description: row.ResultDescription,
		})
	}

	return results, nil
}
