package http

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userSvc := NewMockuserServiceProvider(ctrl)

	want := &Controller{
		userSvc: userSvc,
	}
	got := NewController(userSvc)

	assert.Equal(t, want, got)
}
