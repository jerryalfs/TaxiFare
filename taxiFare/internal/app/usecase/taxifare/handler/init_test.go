package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_taxiFare "taxiFare/internal/app/repository/taxiFare/mock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTaxiFareRedisRepo := mock_taxiFare.NewMockRepository(ctrl)
	mockDetailFareRedisRepo := mock_taxiFare.NewMockRepository(ctrl)
	ucMock := New(mockTaxiFareRedisRepo,mockDetailFareRedisRepo)
	assert.NotNil(t, ucMock)
}
