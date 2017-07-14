package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"falcon-plus/modules/api/app/controller/alarm"
	"falcon-plus/modules/api/app/controller/dashboard_graph"
	"falcon-plus/modules/api/app/controller/dashboard_screen"
	"falcon-plus/modules/api/app/controller/expression"
	"falcon-plus/modules/api/app/controller/graph"
	"falcon-plus/modules/api/app/controller/host"
	"falcon-plus/modules/api/app/controller/mockcfg"
	"falcon-plus/modules/api/app/controller/strategy"
	"falcon-plus/modules/api/app/controller/template"
	"falcon-plus/modules/api/app/controller/uic"
	"falcon-plus/modules/api/app/utils"
)

func StartGin(port string, r *gin.Engine) {
	r.Use(utils.CORS())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm Falcon+ (｡A｡)")
	})
	graph.Routes(r)
	uic.Routes(r)
	template.Routes(r)
	strategy.Routes(r)
	host.Routes(r)
	expression.Routes(r)
	mockcfg.Routes(r)
	dashboard_graph.Routes(r)
	dashboard_screen.Routes(r)
	alarm.Routes(r)
	r.Run(port)
}
