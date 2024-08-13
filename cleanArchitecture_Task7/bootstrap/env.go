package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                string `mapstructure:"APP_ENV"`
	ServerAddress         string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout        int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                string `mapstructure:"DB_HOST"`
	MongoURI              string `mapstructure:"MONGO_URI"`
	DBName                string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	UserCollection        string `mapstructure:"USER_COLLECTION"`
	TaskCollection        string `mapstructure:"TASK_COLLECTION"`
	Secret                string `mapstructure:"SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	
	viper.SetConfigFile(".env")
	viper.SetConfigName("app")
	viper.AddConfigPath(".") 

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
