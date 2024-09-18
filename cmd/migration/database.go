package migration

import (
	"github.com/Zainal21/renco-boilerplate/internal/utils"
	"github.com/Zainal21/renco-boilerplate/pkg/database/postgres"
)

func MigrateDatabase() {
	cfg := utils.LoadConfig(".env")
	postgres.DatabaseMigration(cfg)
}
