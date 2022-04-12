package compapy

import (
	"github.com/gin-gonic/gin"
	"log"
)

type (
	CompanyController interface {
		ListingCompany(ctx *gin.Context)
		GetCompany(ctx *gin.Context)
	}
)

type (
	CompanyRoute struct {
		ctl    CompanyController
		logger *log.Logger
	}
)

func NewCompanyRoute(ctl CompanyController) *CompanyRoute {
	return &CompanyRoute{
		ctl: ctl,
	}
}

func (c *CompanyRoute) Route(e *gin.Engine) {
	api := e.Group("/companies")

	api.POST("", c.ctl.ListingCompany)
	api.GET("/:id", c.ctl.GetCompany)
}
