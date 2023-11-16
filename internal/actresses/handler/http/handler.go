package http

import (
	"net/http"

	"github.com/Lyalyashechka/actresses/internal/actresses"
	"github.com/Lyalyashechka/actresses/internal/models"
	"github.com/Lyalyashechka/actresses/internal/tools/logger"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	logger      logger.Logger
	actressesUC actresses.UseCase
}

func New(logger logger.Logger, actressesUC actresses.UseCase) *Handler {
	return &Handler{
		logger:      logger,
		actressesUC: actressesUC,
	}
}

func (h *Handler) GetAllActresses(echoCtx echo.Context) error {
	ctx := h.logger.WithCaller(echoCtx.Request().Context())

	actresses, err := h.actressesUC.GetAllActresses(ctx)
	if err != nil {
		h.logger.WithCtx(ctx).WithError(err).Errorf("Error getAllActresess")
		return echoCtx.JSON(http.StatusInternalServerError, "Invalid server error")
	}

	return echoCtx.JSON(http.StatusOK, actresses)
}

func (h *Handler) GetTopActresses(echoCtx echo.Context) error {
	ctx := h.logger.WithCaller(echoCtx.Request().Context())

	actresses, err := h.actressesUC.GetTopActresses(ctx)
	if err != nil {
		h.logger.WithCtx(ctx).WithError(err).Errorf("Error getTopActresess")
	}

	return echoCtx.JSON(http.StatusOK, actresses)
}

func (h *Handler) VoteActress(echoCtx echo.Context) error {
	ctx := h.logger.WithCaller(echoCtx.Request().Context())

	voteNew := models.Vote{}
	if err := echoCtx.Bind(&voteNew); err != nil {
		h.logger.WithCtx(ctx).WithError(err).Error("Can not bind vote")
		return echoCtx.JSON(http.StatusBadRequest, "Bad json")
	}

	err := h.actressesUC.VoteActress(ctx, voteNew)
	if err != nil {
		h.logger.WithCtx(ctx).WithError(err).Errorf("Error voteActress")
		return echoCtx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return echoCtx.JSON(http.StatusOK, voteNew)
}
