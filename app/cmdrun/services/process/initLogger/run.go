package initlogger

import (
	"solopg/app/cmdrun/services/logger"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
)

type initializer struct {
	process.Process

	db *mango.Mongo
	
}

func Process(db *mango.Mongo) *initializer {
	return &initializer{
		db: db,
	}
}

func (p *initializer) Run() {
	repo := repository.NewLogSystemRepos()
	repo.SetDb(p.db)

	logger.Init(repo)
}