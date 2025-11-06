package config

import (
	"field-service/common/util"
	"os"

	"github.com/sirupsen/logrus"
)

var Config AppConfig

type AppConfig struct {
	Port                       int             `json:"port"`
	AppName                    string          `json:"appName"`
	AppEnv                     string          `json:"appEnv"`
	SignatureKey               string          `json:"signatureKey"`
	Database                   Database        `json:"database"`
	RateLimiterMaxRequest      float64         `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond      int             `json:"rateLimiterTimeSecond"`
	JwtSecretKey               string          `json:"jwtSecretKey"`
	JwtExpirationTime          int             `json:"jwtExpirationTime"`
	InternalService            InternalService `json:"InternalService"`
	GCSType                    string          `json:"GCSType"`
	GCSProjectID               string          `json:"GCSProjectID"`
	GCSPrivateKeyID            string          `json:"GCSPrivateKeyID"`
	GCSPrivateKey              string          `json:"GCSPrivateKey"`
	GCSClientEmail             string          `json:"GCSClientEmail"`
	GCSClientID                string          `json:"GCSClientID"`
	GCSAuthURI                 string          `json:"GCSAuthURI"`
	GCSTokenURI                string          `json:"GCSTokenURI"`
	GCSAuthProviderX509CertURL string          `json:"gcsAuthProviderX509CertURL"`
	GCSClientX509CertURL       string          `json:"gcsClientX509CertURL"`
	GCSUniverseDomain          string          `json:"gcsUniverseDomain"`
	GCSBucketName              string          `json:"gcsBucketName"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

type InternalService struct {
	User User `json:"user"`
}

type User struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind config: %v", err)
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}
