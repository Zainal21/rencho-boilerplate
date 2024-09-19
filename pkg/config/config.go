package config

type Env struct {
	AppEnv string `mapstructure:"APP_ENV"`
	Host   string `mapstructure:"APP_HOST"`
	Port   string `mapstructure:"APP_PORT"`
	// firebase
	FirebaseCredentialPath    string `mapstructure:"FIREBASE_CREDENTIAL_PATH"`
	FirebaseVerifyPasswordURL string `mapstructure:"FIREBASE_VERIFY_PASSWORD_URL"`
	ContextTimeout            int    `mapstructure:"CONTEXT_TIMEOUT"`
	// database testing
	TestDBUrl      string `mapstructure:"TEST_DB_URL"`
	TestDBUser     string `mapstructure:"TEST_DB_USER"`
	TestDBPassword string `mapstructure:"TEST_DB_PASSWORD"`
	// database real
	DBConnection string `mapstructure:"DB_CONNECTION"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBDatabase   string `mapstructure:"DB_DATABASE"`
	DBUsername   string `mapstructure:"DB_USERNAME"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBSSLMode    string `mapstructure:"DB_SSL_MODE"`

	// DBUrl string `mapstructure:"DATABASE_URL"`

	AesSecret string `mapstructure:"AES_SECRET"`
	// jwt token
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}
