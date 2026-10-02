package repository

import (
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"

	"go.mongodb.org/mongo-driver/bson"
)

// TODO i18N -- register les erreur a la sortie, pas dans le module repository si non ca va merde.
const LoggingCollection mango.Collection = "log_system"

// Methods
type LogSystemRepos struct {
	*MongoRepository
}

func NewLogSystemRepos() *LogSystemRepos {
	return &LogSystemRepos{
		NewRepository(LoggingCollection),
	}
}

// Getters & setters
func (r *LogSystemRepos) GetDb() *mango.Mongo {
	return r.getDb()
}

func (r *LogSystemRepos) SetDb(db *mango.Mongo) *LogSystemRepos {
	r.setDb(db)
	return r
}

// Methods
func (r *LogSystemRepos) Load(
	filter bson.M,
	opts ...FindOption,
) ([]logs.Log, error) {
	return r.fromCollection[logs.Log](filter, opts...)
}

// TODO a la sorier de lui, pas register mais l'afficher direct ou panic ?
func (r *LogSystemRepos) Register(log logs.Log) error {
	return r.toCollection(&log)
}
