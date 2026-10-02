package initlogger

import (
	"solopg/app/services/logger"
	"solopg/app/services/mango"
	"solopg/app/services/mango/repository"
	"solopg/app/services/process/processkit"
)

type process struct {
	processkit.Process

	db *mango.Mongo
	
}

func Process(db *mango.Mongo) *process {
	return &process{
		db: db,
	}
}

func (p *process) Run() {
	repo := repository.NewLogSystemRepos()
	repo.SetDb(p.db)

	logger.Init(repo)
}