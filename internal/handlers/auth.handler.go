package handlers

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/repositories"
	"julianmindria/backendCoffee/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"pass" json:"pass,omitempty"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data *User
	if ers := ctx.ShouldBind(&data); ers != nil {
		pkg.NewRes(500, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthUser(data.Username)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(users.Pass, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Wrong password",
		}).Send(ctx)
		return
	}

	jsonToken := pkg.NewToken(users.Id_user, users.Role)
	tokens, err := jsonToken.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
}
