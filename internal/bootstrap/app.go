package bootstrap

import (
	"firebase.google.com/go/v4/auth"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Env          *config.Env
	DB           *sqlx.DB
	FirebaseAuth *auth.Client
}

func App() Application {
	app := &Application{}
	app.Env = config.LoadConfig()
	app.DB = NewPostgresDB(app.Env)
	RegistryLogger(app.Env)

	// app.FirebaseAuth = NewFirebaseAuth(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.DB)
}
