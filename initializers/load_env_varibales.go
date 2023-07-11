package initializers

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
    Port           int    `mapstructure:"PORT"`
    DBURL          string `mapstructure:"DB_URL"`
    ClientID       string `mapstructure:"CLIENT_ID"`
    ClientSecret   string `mapstructure:"CLIENT_SECERT"`
    PostgresPass   string `mapstructure:"POSTGRES_PASSWORD"`
    PostgresUser   string `mapstructure:"POSTGRES_USER"`
    PostgresDB     string `mapstructure:"POSTGRES_DB"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}