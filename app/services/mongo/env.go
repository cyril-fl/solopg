package mongo

import (
	"os"
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
	uri := cleanJoin("@", env.getCredentials(), env.getHostAt())
	return cleanJoin("", "mongodb://", uri)
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
	return cleanJoin(":", env.username, env.password)
}

func (env *MongoEnvironment) getHostAt() string {
	if env.host == "" {
		env.host = "localhost"
	}
	if env.port == "" {
		env.port = "27017"
	}

	return cleanJoin(":", env.host, env.port)
}

func cleanJoin(sep string, parts ...string) string {
	res := []string{}
	for _, p := range parts {
		if p != "" {
			res = append(res, p)
		}
	}

	return strings.Join(res, sep)
}
