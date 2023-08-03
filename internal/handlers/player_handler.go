package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daxpat2000/cricserver/internal/core/domain"
	"github.com/daxpat2000/cricserver/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService ports.PlayerService
}

func NewPlayerHandler(playerService ports.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}

func (ph *PlayerHandler) CreatePlayer(ctx *gin.Context) {
	var player domain.Player
	err := ctx.ShouldBindJSON(&player)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = ph.playerService.CreatePlayer(ctx.Request.Context(), &player)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf(
			"Player %s created successfully", player.String(),
		),
	})
}

func (ph *PlayerHandler) GetPlayer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	player, err := ph.playerService.GetPlayer(ctx.Request.Context(), uint(id))
	if err != nil {
		HandleError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, player)
}

func (ph *PlayerHandler) ListPlayers(ctx *gin.Context) {
	players, err := ph.playerService.ListPlayers(ctx.Request.Context())
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, players)
}
