package loading

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/yofr4nk/s3-chunked-upload-builder/pkg/domain"
	"os"
)

func GetEnvironmentKeys() (domain.EnvironmentConfig, error) {
	//Checking environment before load .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return domain.EnvironmentConfig{}, errors.New("error loading environment variables " + err.Error())
		}
	}

	envBody := domain.EnvironmentConfig{
		Bucket:       os.Getenv("BUCKET"),
		AwsAccessKey: os.Getenv("ACCESS_KEY"),
		AwsSecretKey: os.Getenv("SECRET_KEY"),
		Region:       os.Getenv("REGION"),
	}

	return envBody, nil
}
