package repository

import (
	"solopg/app/domain/system/logger"
	"solopg/app/services/mango"

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
) ([]logger.Log, error) {
	return r.fromCollection[logger.Log](filter, opts...)
}

// TODO a la sorier de lui, pas register mais l'afficher direct ou panic ?
func (r *LogSystemRepos) Register(log logger.Log) error {
	return r.toCollection(&log)
}

func TestLog() error {
	log := logger.New(logger.Template{
		Type:    logger.INFO,
		Message: "This is a test log entry",
	})

	db := mango.New()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}
	defer db.Disconnect()

	repo := NewLogSystemRepos().SetDb(db)
	
	return repo.Register(log)
}
