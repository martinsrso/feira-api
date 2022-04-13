package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Market struct {
	ID         int    `csv:"ID" json:"ID,omitempty"`
	Long       int    `csv:"LONG" json:"LONG,omitempty"`
	Lat        int    `csv:"LAT" json:"LAT,omitempty"`
	SetCens    int    `csv:"SETCENS" json:"SETCENS,omitempty"`
	CodDist    string `csv:"CODDIST" json:"CODDIST,omitempty"`
	Distrito   string `csv:"DISTRITO,omitempty" json:"DISTRITO,omitempty"`
	CodSubPref int    `csv:"CODSUBPREF" json:"CODSUBPREF,omitempty"`
	SubPrefe   string `csv:"SUBPREFE" json:"SUBPREFE,omitempty"`
	Regiao5    string `csv:"REGIAO5" json:"REGIAO5,omitempty"`
	Regiao8    string `csv:"REGIAO8" json:"REGIAO8,omitempty"`
	NomeFeira  string `csv:"NOME_FEIRA" json:"NOME_FEIRA,omitempty"`
	Registro   string `csv:"REGISTRO" json:"REGISTRO,omitempty" gorm:"primaryKey"`
	Logradouro string `csv:"LOGRADOURO" json:"LOGRADOURO,omitempty"`
	Numero     string `csv:"NUMERO,omitempty" json:"NUMERO,omitempty"`
	Referencia string `csv:"REFERENCIA,omitempty" json:"REFERENCIA,omitempty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type MarketUsecase interface {
	GetByRegister(ctx context.Context, reg string) (*Market, error)
	GetByName(ctx context.Context, name string) (*Market, error)
	Update(ctx context.Context, m *Market, dtoM *Market) error
	Store(context.Context, *Market) error
	Delete(ctx context.Context, reg string) error
}

type MarketRepository interface {
	GetByRegister(ctx context.Context, reg string) (Market, error)
	GetByName(ctx context.Context, name string) (Market, error)
	Update(ctx context.Context, m *Market, dtoM *Market) error
	Store(ctx context.Context, m *Market) error
	Delete(ctx context.Context, reg string) error
}
