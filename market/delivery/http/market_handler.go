package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martinsrso/feira-api/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type MarketHandler struct {
	MarketUsecase domain.MarketUsecase
}

func NewMarketHandler(g *gin.Engine, mus domain.MarketUsecase) {
	handler := &MarketHandler{
		MarketUsecase: mus,
	}

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Feira API SP"})
	})
	g.POST("/market", handler.Store)
	g.GET("/market/:register", handler.GetByRegister)
	g.DELETE("/market/:register", handler.Delete)
	g.PUT("/market/:register", handler.Update)
	g.GET("/market", handler.GetByName)
}

func (m *MarketHandler) GetByName(g *gin.Context) {
	name := g.Query("nome")
	if name == "" {
		g.JSON(http.StatusNotFound, ResponseError{Message: domain.ErrNotFound.Error()})
		return
	}

	ctx := g.Request.Context()
	market, errGet := m.MarketUsecase.GetByName(ctx, name)
	if errGet != nil {
		g.JSON(http.StatusBadRequest, ResponseError{Message: errGet.Error()})
		return
	}

	if market == nil {
		g.JSON(http.StatusNotFound, ResponseError{Message: domain.ErrNotFound.Error()})
		return
	}

	g.JSON(http.StatusOK, market)
}

func (m *MarketHandler) Update(g *gin.Context) {
	reg := g.Param("register")
	if reg == "" {
		g.JSON(http.StatusNotFound, ResponseError{Message: domain.ErrNotFound.Error()})
		return
	}

	var dtoMarket domain.Market
	err := g.Bind(&dtoMarket)

	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		return
	}

	var ok bool
	if ok, err = isRequestValid(&dtoMarket); !ok {
		g.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	ctx := g.Request.Context()
	market, err := m.MarketUsecase.GetByRegister(ctx, reg)
	if err != nil {
		g.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	err = m.MarketUsecase.Update(ctx, market, &dtoMarket)
	if err != nil {
		g.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	if err != nil {
		g.JSON(http.StatusConflict, ResponseError{Message: err.Error()})
		return
	}

	g.JSON(http.StatusOK, market)
}

func (m *MarketHandler) Delete(g *gin.Context) {
	reg := g.Param("register")
	if reg == "" {
		g.JSON(http.StatusNotFound, ResponseError{Message: domain.ErrNotFound.Error()})
		return
	}

	ctx := g.Request.Context()
	err := m.MarketUsecase.Delete(ctx, reg)
	if err != nil {
		g.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	g.JSON(http.StatusNoContent, "")
}

func (m *MarketHandler) GetByRegister(g *gin.Context) {
	reg := g.Param("register")
	if reg == "" {
		g.JSON(http.StatusNotFound, ResponseError{Message: domain.ErrNotFound.Error()})
		return
	}

	ctx := g.Request.Context()
	market, errGet := m.MarketUsecase.GetByRegister(ctx, reg)
	if errGet != nil {
		g.JSON(http.StatusBadRequest, ResponseError{Message: errGet.Error()})
		return
	}

	if market == nil {
		g.JSON(http.StatusNotFound, ResponseError{domain.ErrNotFound.Error()})
		return
	}

	g.JSON(http.StatusOK, market)
}

func (m *MarketHandler) Store(g *gin.Context) {
	var market domain.Market

	err := g.Bind(&market)

	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		return
	}

	var ok bool
	if ok, err = isRequestValid(&market); !ok {
		g.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	ctx := g.Request.Context()
	err = m.MarketUsecase.Store(ctx, &market)

	if err != nil {
		g.JSON(http.StatusConflict, ResponseError{Message: err.Error()})
		return
	}

	g.JSON(http.StatusCreated, market)
}

func isRequestValid(m *domain.Market) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
