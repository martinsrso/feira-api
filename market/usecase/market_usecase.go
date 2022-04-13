package usecase

import (
	"context"
	"time"

	"github.com/martinsrso/feira-api/domain"
)

type marketUsecase struct {
	marketRepo     domain.MarketRepository
	contextTimeout time.Duration
}

func NewMarketUsecase(m domain.MarketRepository, timeout time.Duration) domain.MarketUsecase {
	return &marketUsecase{
		marketRepo:     m,
		contextTimeout: timeout,
	}
}

func (m *marketUsecase) GetByName(ctx context.Context, name string) (*domain.Market, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, err := m.marketRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *marketUsecase) Update(ctx context.Context, market *domain.Market, dtoMarket *domain.Market) error {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	// check if exists market

	err := m.marketRepo.Update(ctx, market, dtoMarket)
	if err != nil {
		return err
	}

	return nil
}

func (m *marketUsecase) Store(ctx context.Context, market *domain.Market) error {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	// check if exists market

	err := m.marketRepo.Store(ctx, market)
	if err != nil {
		return err
	}

	return nil
}

func (m *marketUsecase) GetByRegister(ctx context.Context, reg string) (*domain.Market, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, err := m.marketRepo.GetByRegister(ctx, reg)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *marketUsecase) Delete(ctx context.Context, reg string) error {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res := m.marketRepo.Delete(ctx, reg)
	if res != nil {
		return res
	}

	return nil
}
