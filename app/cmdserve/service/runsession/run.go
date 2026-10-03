package cmdserverunsession

import (
	"context"
	"os"
	"os/signal"

	cmdservetui "solopg/app/cmdserve/tui"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
	"solopg/app/shared/services/process/generatesteplist"

	tea "charm.land/bubbletea/v2"
)

type runner[T any] struct {
	process.Process
	cache[T]

	getRepository func() repository.Watchable[T]
	// getFlags func() *pflag.FlagSet
}

type cache[T any] struct {
	view       *generatesteplist.Generator[cmdservetui.Context]
	repository repository.Watchable[T]
}

func Process[T any](
	getRepository func() repository.Watchable[T],
	// getFlags func() *pflag.FlagSet,
) *runner[T] {
	return &runner[T]{
		getRepository: getRepository,
		// getFlags: getFlags,
	}
}

func (p *runner[T]) Run() {
	p.setRepository()
	p.createTuiView()
	p.runTuiView()
}

func (p *runner[T]) GetResult() {
	logs.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "session",
	})
}

// Methods
func (p *runner[T]) setRepository() {
	p.repository = p.getRepository()
}

func (p *runner[T]) createTuiView() {
	p.cache.view = generatesteplist.Process[cmdservetui.Context]()
}

func (p *runner[T]) runTuiView() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	/*
		FIXME MEDIUM de vrai etre une vue et nom l'equilavent du moteur qui tourne dans CMDRUN
		Deplaver le moteur aiderais a ameliore et fixer ici
		Ne devrai pas rester comme ça ca ca marche mais ma melanger les responsabilté
	*/
	model := cmdservetui.NewModel(ctx, p.cache.view.GetSteps(), p.repository)

	p.cache.view.Run(&model, tea.WithContext(ctx))
	if err := model.Err(); err != nil {
		p.SetErr(err)
	}

	if p.cache.view.HasErr() {
		p.SetErr(p.cache.view.GetErr())
	}
}

// Helpers
