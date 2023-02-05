package mariadb

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alanpryoga/blog-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestRepository_FindUserByID(t *testing.T) {
	type mockFields struct {
		db sqlmock.Sqlmock
	}
	type args struct {
		ctx context.Context
		id  int32
	}
	tests := []struct {
		name    string
		args    args
		mock    func(mf mockFields)
		want    entity.User
		wantErr error
	}{
		{
			name: "Given no error and row exist " +
				"Then it should return non-empty struct and nil error",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func(mf mockFields) {
				mf.db.ExpectQuery(findUserByIDQuery).
					WithArgs(int32(1)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "hashed_password"}).
						AddRow(int32(1), "user@example.test", "secret"))
			},
			want: entity.User{
				ID:             1,
				Email:          "user@example.test",
				HashedPassword: "secret",
			},
			wantErr: nil,
		},
		{
			name: "Given no error and no row exist " +
				"Then it should return empty struct and non-nil error",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func(mf mockFields) {
				mf.db.ExpectQuery(findUserByIDQuery).
					WithArgs(int32(1)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "hashed_password"}))
			},
			want:    entity.User{},
			wantErr: entity.ErrUserNotFound,
		},
		{
			name: "Given an error " +
				"Then it should return empty struct and non-nil error",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func(mf mockFields) {
				mf.db.ExpectQuery(findUserByIDQuery).
					WithArgs(int32(1)).
					WillReturnError(assert.AnError)
			},
			want:    entity.User{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mf := mockFields{
				db: mock,
			}
			tt.mock(mf)
			repo := &Repository{
				db: db,
			}
			got, err := repo.FindUserByID(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
