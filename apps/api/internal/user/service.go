package user

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

func (s *Service) UpdateUser(ctx context.Context, userID int64, name string) (sqlc.User, error) {
	user, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		return sqlc.User{}, err
	}

	return s.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        userID,
		Name:      name,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
	})
}
