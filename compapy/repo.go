package compapy

import (
	"context"
	"database/sql"
	"log"
	"os"
)

type (
	CompanyOracleRepo struct {
		db     *sql.DB
		logger *log.Logger
	}
)

func NewCompanyOracleRepo(db *sql.DB) *CompanyOracleRepo {
	return &CompanyOracleRepo{
		db:     db,
		logger: log.New(os.Stdout, "company-oracle-repo", log.LstdFlags),
	}
}

func (c *CompanyOracleRepo) InsertCompany(ctx context.Context, comp *Company) (id string, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *CompanyOracleRepo) GetCompanyBySymbol(ctx context.Context, symbol string) (*Company, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CompanyOracleRepo) GetCompanyByID(ctx context.Context, id string) (*Company, error) {
	//TODO implement me
	panic("implement me")
}
