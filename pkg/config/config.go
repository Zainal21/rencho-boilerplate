package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppName string `mapstructure:"APP_NAME"`
	Host    string `mapstructure:"APP_HOST"`
	Port    string `mapstructure:"APP_PORT"`

	AppLoggerDebug bool   `mapstructure:"APP_LOGGER_DEBUG"`
	AppLoggerLevel string `mapstructure:"APP_LOGGER_LEVEL"`
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

func LoadConfig() *Env {
	env := Env{}
	err := loadConfig()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Loaded config %+v)", err))
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Environment can't be loaded %+v)", err))
	}

	if env.AppEnv == "development" {
		logrus.Info("The App is running in development mode")
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config %v was change", e.Name)
	})

	return &env
}

func loadConfig() error {
	files, err := getAllConfigFiles("./config")
	if err != nil {
		logrus.Warn(err)
	}

	viper.AddConfigPath("./config")
	for _, file := range files {
		viper.SetConfigType("json")
		viper.SetConfigFile(file)
		err = viper.MergeInConfig()
		if err != nil {
			return err
		}
	}

	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	_ = viper.MergeInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config %v was change", e.Name)
	})

	return nil
}

func getAllConfigFiles(folderPath string) ([]string, error) {
	var configFiles []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			configFiles = append(configFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return configFiles, nil
}
