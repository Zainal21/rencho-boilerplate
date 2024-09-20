package migration

import (
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/database/postgres"
)

func MigrateDatabase() {
	cfg := config.LoadConfig()
	postgres.DatabaseMigration(cfg)
}
