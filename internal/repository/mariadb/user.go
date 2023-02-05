package mariadb

import (
	"context"
	"database/sql"

	"github.com/alanpryoga/blog-api/internal/entity"
)

func (repo *Repository) FindUserByID(ctx context.Context, id int32) (entity.User, error) {
	var user entity.User

	err := repo.db.QueryRowContext(ctx, findUserByIDQuery, id).Scan(&user.ID, &user.Email, &user.HashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, entity.ErrUserNotFound
		}
		return user, err
	}

	return user, nil
}
