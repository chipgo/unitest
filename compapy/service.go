package compapy

import (
	"context"
	"errors"
	"log"
	"os"
)

// CompanyRepo https://github.com/golang/mock

//go:generate mockgen -destination=company_repo_mock_test.go -package=company_test . CompanyRepo
type (
	CompanyRepo interface {
		InsertCompany(ctx context.Context, comp *Company) (id string, err error)
		GetCompanyBySymbol(ctx context.Context, symbol string) (*Company, error)
		GetCompanyByID(ctx context.Context, id string) (*Company, error)
	}
)

type (
	CompanyServiceImpl struct {
		repo   CompanyRepo
		logger *log.Logger
	}
)

func NewCompanyService(repo CompanyRepo) *CompanyServiceImpl {
	return &CompanyServiceImpl{
		repo:   repo,
		logger: log.New(os.Stdout, "company-service", log.LstdFlags),
	}
}

func (c *CompanyServiceImpl) CreateCompany(ctx context.Context, comp *Company) (id string, err error) {
	/*
		verify company fields ...
		if OK, then
	*/
	if _, err := c.repo.GetCompanyBySymbol(ctx, comp.Symbol); err != nil && !errors.Is(err, ErrNotFound) {
		return "", err
	}

	id, err = c.repo.InsertCompany(ctx, comp)
	return
}

func (c *CompanyServiceImpl) ReadOneCompanyByID(ctx context.Context, id string) (comp *Company, err error) {

	if len(id) == 0 {
		return nil, errors.New("id should not empty")
	}

	comp, err = c.repo.GetCompanyByID(ctx, id)
	return
}
