package handlers

import (
	"julianmindria/backendCoffee/internal/models"
	"julianmindria/backendCoffee/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlerfav struct {
	*repositories.RepoOrder
}

func NewFav(r *repositories.RepoOrder) *Handlerfav {
	return &Handlerfav{r}
}

func (h *Handlerfav) AddOrder(ctx *gin.Context) {
	var order models.Order_Product
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.Add_Order_product(&order)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *Handlerfav) RemoveOrder(ctx *gin.Context) {
	var order models.Order_Product
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.Delete_Order_product(&order)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}
