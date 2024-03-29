package store

import "context"

type (
	AuthRepository struct {
		store *Store
	}
)

func NewAuthRepository(store *Store) *AuthRepository {
	return &AuthRepository{
		store: store,
	}
}

func (r *AuthRepository) Find(apiKey string, ctx context.Context) (int, error) {
	var count int
	err := r.store.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM auth WHERE api_key = ?", apiKey).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
