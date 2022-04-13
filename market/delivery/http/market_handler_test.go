package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/martinsrso/feira-api/domain"
	"github.com/martinsrso/feira-api/domain/mock"
	"github.com/stretchr/testify/assert"
)

// TODO: others test

func TestMarketHandler_GetByRegister(t *testing.T) {
	var mockMarket domain.Market
	err := faker.FakeData(&mockMarket)
	assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	mockUCase := mock.NewMockMarketUsecase(ctrl)

	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		assert.NoError(t, err)

		r := gin.Default()
		req, _ := http.NewRequest("GET", "/market/10", nil)

		mockUCase.EXPECT().GetByRegister(gomock.Any(), "10").Return(&mockMarket, nil)
		NewMarketHandler(r, mockUCase)

		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("error-failed", func(t *testing.T) {
		w := httptest.NewRecorder()
		w.Code = 400

		assert.NoError(t, err)

		r := gin.Default()
		req, _ := http.NewRequest("GET", "/market/10", nil)

		mockUCase.EXPECT().GetByRegister(gomock.Any(), gomock.Any()).Return(&domain.Market{}, errors.New("error"))
		NewMarketHandler(r, mockUCase)

		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestMarketHandler_Store(t *testing.T) {
	type args struct {
		g *gin.Context
	}
	tests := []struct {
		name string
		m    *MarketHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Store(tt.args.g)
		})
	}
}

func Test_isRequestValid(t *testing.T) {
	type args struct {
		m *domain.Market
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isRequestValid(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("isRequestValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isRequestValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
