package config

type Env struct {
	AppEnv                    string `mapstructure:"APP_ENV"`
	Host                      string `mapstructure:"HOST"`
	Port                      string `mapstructure:"PORT"`
	FirebaseCredentialPath    string `mapstructure:"FIREBASE_CREDENTIAL_PATH"`
	FirebaseVerifyPasswordURL string `mapstructure:"FIREBASE_VERIFY_PASSWORD_URL"`
	ContextTimeout            int    `mapstructure:"CONTEXT_TIMEOUT"`
	TestDBUrl                 string `mapstructure:"TEST_DB_URL"`
	TestDBUser                string `mapstructure:"TEST_DB_USER"`
	TestDBPassword            string `mapstructure:"TEST_DB_PASSWORD"`
	DBUrl                     string `mapstructure:"DATABASE_URL"`
	DBName                    string `mapstructure:"DB_NAME"`
	AesSecret                 string `mapstructure:"AES_SECRET"`
	AccessTokenExpiryHour     int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour    int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret         string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret        string `mapstructure:"REFRESH_TOKEN_SECRET"`
}
