package handlers

import (
	"julianmindria/backendCoffee/internal/models"
	"julianmindria/backendCoffee/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlerfav struct {
	*repositories.RepoFav
}

func NewFav(r *repositories.RepoFav) *Handlerfav {
	return &Handlerfav{r}
}

func (h *Handlerfav) Addfavorite(ctx *gin.Context) {
	var favorite models.Fav_Product
	if err := ctx.ShouldBind(&favorite); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.Create_favproduct(&favorite)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}
