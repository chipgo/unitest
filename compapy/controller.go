package compapy

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type (
	CompanyService interface {
		CreateCompany(ctx context.Context, comp *Company) (id string, err error)
		ReadOneCompanyByID(ctx context.Context, id string) (comp *Company, err error)
	}
)

type (
	CompanyControllerImpl struct {
		svc    CompanyService
		logger *log.Logger
	}
)

func NewCompanyController(svc CompanyService) *CompanyControllerImpl {
	return &CompanyControllerImpl{svc: svc, logger: log.New(os.Stdout, "company-controller:", log.LstdFlags)}
}

func (c *CompanyControllerImpl) ListingCompany(ctx *gin.Context) {

	comp := new(Company)
	if err := ctx.ShouldBindJSON(comp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "1200",
			"error":  err.Error(),
		})
		ctx.Abort()
		return
	}

	_, err := c.svc.CreateCompany(ctx, comp)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "1300",
			"error":  err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "1000",
		"data":   comp,
	})
}

func (c *CompanyControllerImpl) GetCompany(ctx *gin.Context) {

	id := ctx.Param("id")
	comp, err := c.svc.ReadOneCompanyByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "1400",
			"error":  err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "1000",
		"data":   comp,
	})
}
