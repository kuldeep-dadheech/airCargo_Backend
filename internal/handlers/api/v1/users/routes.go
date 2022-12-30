package users

import (

	"github.com/gin-gonic/gin"
)

type GroupRoutes interface{
	Initialize(prefix string, r gin.IRouter)
}

type Routes struct {
	handler RoutesHandler
}
func New(
	handler RoutesHandler,
) *Routes {
	return &Routes{
		handler : handler,
	}
}

func (ro *Routes) Initialize(prefix string, r gin.IRouter){
	g :=r.Group(prefix)
	
	{
		g.GET("/hey", ro.handler.Hey)
		g.GET("/admin/:id" , ro.handler.Get )
		g.POST("/signup", ro.handler.Create)
		g.POST("/signin", ro.handler.Signin)
	}
}