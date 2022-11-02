//go:build integration
// +build integration

package test

import (
	"context"

	"go-template/internals/config"
	"go-template/internals/entity"
	"go-template/internals/infrastructure/database"
	"go-template/internals/repository/postgres"

	"github.com/stretchr/testify/suite"
)

type Filter struct {
	PlatformID string `gorm:"type:varchar(255)" json:"platform_id"`
	Subject    string `gorm:"type:varchar(255)" json:"subject"`
	Metadata   string `gorm:"type:JSONB" json:"metadata"`
}

type PackageTestSuite struct {
	suite.Suite
	ctx     context.Context
	repo    *postgres.PostgresRepository
	verbose *entity.Info
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	conf := config.NewConfiguration()
	connetBase := database.NewServerBase(conf)
	suite.repo = postgres.NewRepository(connetBase)
}

func (suite *PackageTestSuite) SetupTest() {
	conf := config.NewConfiguration()
	connetBase := database.NewServerBase(conf)
	suite.repo = postgres.NewRepository(connetBase)
}

func (suite *PackageTestSuite) makeTestStruct(name string, detail string) (test *entity.Info) {
	return &entity.Info{
		Name:   name,
		Detail: detail,
	}
}
