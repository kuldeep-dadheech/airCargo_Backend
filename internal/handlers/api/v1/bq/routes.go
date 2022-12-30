package bq

import "github.com/gin-gonic/gin"

type GroupRoutes interface {
	Initialize(prefix string, r gin.IRouter)
}

type Routes struct {
	handler RoutesHandler
}

func New(
	handler RoutesHandler,
) *Routes {
	return &Routes{
		handler: handler,
	}
}

func (ro *Routes) Initialize(prefix string, r gin.IRouter) {
	g := r.Group(prefix)

	{
		g.POST("/create", ro.handler.Create)
		g.PUT("/update", ro.handler.Update)
		g.DELETE("/delete/:id", ro.handler.Delete)
		g.GET("/list", ro.handler.List)
		g.GET("/get/:id", ro.handler.Get)
	}
}