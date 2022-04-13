package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/martinsrso/feira-api/domain"
	"github.com/martinsrso/feira-api/domain/mock"
)

func getFakeMarket() *domain.Market {
	var fakeData domain.Market
	if err := faker.FakeData(&fakeData); err != nil {
		return nil
	}
	return &fakeData
}

// TODO: make the other test

func Test_marketUsecase_GetByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMarketRepo := mock.NewMockMarketRepository(ctrl)
	mockMarket := getFakeMarket()

	t.Run("success", func(t *testing.T) {
		mockMarketRepo.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(*mockMarket, nil)
		u := NewMarketUsecase(mockMarketRepo, 5*time.Second)
		m, err := u.GetByName(context.Background(), mockMarket.NomeFeira)
		assert.Equal(t, m, mockMarket)
		assert.NoError(t, err)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockMarketRepo.EXPECT().GetByName(gomock.Any(), gomock.Any()).Return(domain.Market{}, errors.New("Unexpectde Error"))
		u := NewMarketUsecase(mockMarketRepo, 5*time.Second)
		_, err := u.GetByName(context.Background(), mockMarket.NomeFeira)
		assert.Error(t, err)
	})
}

func Test_marketUsecase_Update(t *testing.T) {
	type args struct {
		ctx       context.Context
		market    *domain.Market
		dtoMarket *domain.Market
	}
	tests := []struct {
		name    string
		m       *marketUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Update(tt.args.ctx, tt.args.market, tt.args.dtoMarket); (err != nil) != tt.wantErr {
				t.Errorf("marketUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketUsecase_Store(t *testing.T) {
	type args struct {
		ctx    context.Context
		market *domain.Market
	}
	tests := []struct {
		name    string
		m       *marketUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Store(tt.args.ctx, tt.args.market); (err != nil) != tt.wantErr {
				t.Errorf("marketUsecase.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketUsecase_GetByRegister(t *testing.T) {
	type args struct {
		ctx context.Context
		reg string
	}
	tests := []struct {
		name    string
		m       *marketUsecase
		args    args
		want    *domain.Market
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetByRegister(tt.args.ctx, tt.args.reg)
			if (err != nil) != tt.wantErr {
				t.Errorf("marketUsecase.GetByRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marketUsecase.GetByRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marketUsecase_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		reg string
	}
	tests := []struct {
		name    string
		m       *marketUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Delete(tt.args.ctx, tt.args.reg); (err != nil) != tt.wantErr {
				t.Errorf("marketUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
