package v1

import (
	"sagebackend/internal/handlers/api/v1/countries"
	"sagebackend/internal/handlers/api/v1/currencies"
	"sagebackend/internal/handlers/api/v1/users"
	"sagebackend/internal/handlers/api/v1/enquiries"

    "sagebackend/internal/handlers/api/v1/bq"
	"github.com/gin-gonic/gin"
)

type GroupRoutes interface {
	Initialize(prefix string, r gin.IRouter)
}

type Routes struct {
	countryRoutes  countries.GroupRoutes
	currencyRoutes currencies.GroupRoutes
	usersRoutes    users.GroupRoutes
	enquiriesRoutes enquiries.GroupRoutes

    bqRoutes bq.GroupRoutes
}

func New(
	countryRoutes countries.GroupRoutes,
	currencyRoutes currencies.GroupRoutes,
	usersRoutes users.GroupRoutes,
	enquiriesRoutes enquiries.GroupRoutes,

    bqRoutes bq.GroupRoutes,
) *Routes {
	return &Routes{
		countryRoutes:  countryRoutes,
		currencyRoutes: currencyRoutes,
		usersRoutes:    usersRoutes,
		enquiriesRoutes: enquiriesRoutes,

        bqRoutes: bqRoutes,
	}
}

func (ro *Routes) Initialize(prefix string, r gin.IRouter) {

	r.GET("/data/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	v1 := r.Group(prefix)
	{
		ro.countryRoutes.Initialize("/countries", v1)
		ro.currencyRoutes.Initialize("/currencies", v1)
		ro.usersRoutes.Initialize("/users", v1)
		ro.enquiriesRoutes.Initialize("/enquiries", v1)

        ro.bqRoutes.Initialize("/bq", v1)
	}
}
