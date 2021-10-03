package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ped-alm/origin-take-home/adapter"
	"github.com/ped-alm/origin-take-home/controller"
	"github.com/ped-alm/origin-take-home/infra/http/model"
	"net/http"
)

type Gin struct {
	UserController controller.User
}

func NewGin(userController controller.User) *Gin {
	return &Gin{userController}
}

const badRequest = "Bad Request"

func (g *Gin) Start(port string) {
	router := gin.Default()

	router.POST("/user/risk", g.userRisk)

	err := router.Run(port)
	if err != nil {
		fmt.Println("error trying to run gin: " + err.Error())
	}
}

func (g *Gin) userRisk(c *gin.Context) {
	var input model.UserProfile

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": badRequest})
		return
	}

	output := adapter.RiskProfileEntityToHttp(g.UserController.CalculateRisk(adapter.UserProfileHttpToEntity(input)))

	c.JSON(http.StatusOK, output)
}
