package user

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbRepo := NewMockdbRepoProvider(ctrl)

	want := &Service{
		dbRepo: dbRepo,
	}
	got := NewService(dbRepo)

	assert.Equal(t, want, got)
}
