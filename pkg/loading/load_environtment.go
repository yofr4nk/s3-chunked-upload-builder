package loading

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type EnvironmentBody struct {
	Bucket       string
	AwsAccessKey string
	AwsSecretKey string
	Region       string
}

func GetEnvironmentKeys() (EnvironmentBody, error) {
	//Checking environment before load .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return EnvironmentBody{}, errors.New("error loading environment variables " + err.Error())
		}
	}

	envBody := EnvironmentBody{
		Bucket:       os.Getenv("BUCKET"),
		AwsAccessKey: os.Getenv("ACCESS_KEY"),
		AwsSecretKey: os.Getenv("SECRET_KEY"),
		Region:       os.Getenv("REGION"),
	}

	return envBody, nil
}
