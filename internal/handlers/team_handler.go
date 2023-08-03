package handlers

import (
	"fmt"
	"net/http"

	"github.com/daxpat2000/cricserver/internal/core/domain"
	"github.com/daxpat2000/cricserver/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService ports.TeamService
}

func NewTeamHandler(teamService ports.TeamService) *TeamHandler {
	return &TeamHandler{
		teamService: teamService,
	}
}

func (th *TeamHandler) CreateTeam(ctx *gin.Context) {
	var team domain.Team
	err := ctx.ShouldBindJSON(&team)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = th.teamService.CreateTeam(ctx.Request.Context(), &team)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": fmt.Sprintf("Successfully created team %s", team.String()),
	})
}

func (th *TeamHandler) GetTeam(ctx *gin.Context) {
	name := ctx.Param("name")
	teamType := ctx.Param("teamType")
	team, err := th.teamService.GetTeam(ctx.Request.Context(), name, teamType)
	if err != nil {
		HandleError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, team)
}

func (th *TeamHandler) ListTeams(ctx *gin.Context) {
	teams, err := th.teamService.ListTeams(ctx.Request.Context())
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, teams)
}
