package mango

import (
	"os"
	"solopg/app/shared/utils/transform"
	"solopg/config"
	"strings"
)

type MongoEnvironment struct {
	username string
	password string
	port     string
	host     string
	dbname   string
}

const (
	HOST = "localhost"
	PORT = "27017"
)

func NewEnvironment() *MongoEnvironment {
	return &MongoEnvironment{}
}

func (env *MongoEnvironment) ReadEnv() {
	env.username = strings.TrimSpace(os.Getenv("MONGO_USERNAME"))
	env.password = strings.TrimSpace(os.Getenv("MONGO_PASSWORD"))
	env.port = strings.TrimSpace(os.Getenv("MONGO_PORT"))
	env.host = strings.TrimSpace(os.Getenv("MONGO_HOST"))
	env.dbname = strings.TrimSpace(os.Getenv("MONGO_DBNAME"))
}

func (env *MongoEnvironment) GetEnvURI() string {
	uri := transform.CleanJoin("@", env.getCredentials(), env.getHostAt())
	return transform.CleanJoin("", "mongodb://", uri)
}

func (env *MongoEnvironment) SetEnvURI() {
	uri := env.GetEnvURI()
	os.Setenv("MONGO_URI", uri)
}

func (env *MongoEnvironment) GetEnvDBName() string {
	if env.dbname == "" {
		env.dbname = config.Current.Name
	}

	return env.dbname
}

func (env *MongoEnvironment) getCredentials() string {
	if env.username == "" || env.password == "" {
		return ""
	}
	return transform.CleanJoin(":", env.username, env.password)
}

func (env *MongoEnvironment) getHostAt() string {
	if env.host == "" {
		env.host = HOST
	}
	if env.port == "" {
		env.port = PORT
	}

	return transform.CleanJoin(":", env.host, env.port)
}
