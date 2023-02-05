package redis

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	db, _ := redismock.NewClientMock()

	want := &Repository{
		db: db,
	}
	got := NewRepository(db)

	assert.Equal(t, want, got)
}
