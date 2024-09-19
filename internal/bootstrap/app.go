package bootstrap

import (
	"firebase.google.com/go/v4/auth"
	"github.com/Zainal21/renco-boilerplate/internal/utils"
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
	app.Env = utils.LoadConfig(".env")
	app.DB = NewPostgresDB(app.Env)
	// app.FirebaseAuth = NewFirebaseAuth(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.DB)
}