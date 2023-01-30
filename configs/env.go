package configs

import (
	"fiber_mongo_zap/logger"
	"github.com/joho/godotenv"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load("config.env")
	if err != nil {
		logger.Debug("Error loading .env file ", err)
	}
	return os.Getenv("MONGOURI")
}
