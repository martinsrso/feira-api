package postgres

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/martinsrso/feira-api/domain"
	"gorm.io/gorm"
)

type postgresMarketRepository struct {
	DB *gorm.DB
}

func NewPostgresMarketRepository(DB *gorm.DB) domain.MarketRepository {
	return &postgresMarketRepository{DB}
}

func (p *postgresMarketRepository) Update(ctx context.Context, m *domain.Market, dtoM *domain.Market) error {
	p.DB.WithContext(ctx)

	var updMap map[string]interface{}
	data, _ := json.Marshal(dtoM)
	json.Unmarshal(data, &updMap)

	result := p.DB.Model(&m).Updates(dtoM)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *postgresMarketRepository) Store(ctx context.Context, m *domain.Market) error {
	p.DB.WithContext(ctx)

	result := p.DB.Create(&m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *postgresMarketRepository) GetByName(ctx context.Context, name string) (domain.Market, error) {
	p.DB.WithContext(ctx)

	var market domain.Market

	result := p.DB.First(&market, "nome_feira = ?", name).Limit(1)

	if result.Error != nil {
		return market, result.Error
	}

	return market, nil
}

func (p *postgresMarketRepository) GetByRegister(ctx context.Context, reg string) (domain.Market, error) {
	p.DB.WithContext(ctx)

	var market domain.Market

	result := p.DB.First(&market, "registro = ?", reg).Limit(1)

	if result.Error != nil {
		return market, result.Error
	}

	return market, nil
}

func (p *postgresMarketRepository) Delete(ctx context.Context, reg string) error {
	p.DB.WithContext(ctx)

	var market domain.Market

	result := p.DB.Delete(&market, "registro = ?", reg)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("the market cannot be deleted")
	}

	return nil
}
