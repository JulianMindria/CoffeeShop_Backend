package handlers

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/models"
	"julianmindria/backendCoffee/internal/repositories"
	"julianmindria/backendCoffee/pkg"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) Postdata(ctx *gin.Context) {
	var err error
	user := models.User{
		Role: "user",
	}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	user.Image_file = ctx.MustGet("image").(string)
	user.Pass, err = pkg.HashPassword(user.Pass)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	response, er := h.CreateUser(&user)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) Updatedata(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Image_file = ctx.MustGet("image").(string)
	response, er := h.UpdateUser(&user)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) Deletedata(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.DeleteUser(&user)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) Getdata(ctx *gin.Context) {
	var users models.User
	search := ctx.Query("search")
	if err := ctx.ShouldBind(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Get_Users(&users, search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}
