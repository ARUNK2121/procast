package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost                string `mapstructure:"DB_HOST"`
	DBName                string `mapstructure:"DB_NAME"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	AUTHTOKEN             string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ACCOUNTSID            string `mapstructure:"TWILIO_ACCOUNT_SID"`
	SERVICESID            string `mapstructure:"TWILIO_SERVICE_SID"`
	AWS_REGION            string `mapstructure:"AWS_REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "TWILIO_AUTH_TOKEN", "TWILIO_ACCOUNT_SID", "TWILIO_SERVICE_SID", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
