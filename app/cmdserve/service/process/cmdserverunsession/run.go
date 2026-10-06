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
	interfass "solopg/app/shared/types/interface"

	tea "charm.land/bubbletea/v2"
	"github.com/spf13/pflag"
)

type run[T interfass.Stringable] struct {
	process.Process
	cache[T]

	getRepository func() repository.Watchable[T]
	getFlags      func() *pflag.FlagSet
}

type cache[T interfass.Stringable] struct {
	view       *generatesteplist.Generator[cmdservetui.Context]
	repository repository.Watchable[T]
}

type ProcessTemplate[T interfass.Stringable] struct {
	GetRepository func() repository.Watchable[T]
	GetFlags      func() *pflag.FlagSet
}

func Process[T interfass.Stringable](
	params ProcessTemplate[T],
) *run[T] {
	return &run[T]{
		getRepository: params.GetRepository,
		getFlags:      params.GetFlags,
	}
}

func (p *run[T]) Run() {
	p.setRepository()
	p.createTuiView()
	p.runTuiView()
}

func (p *run[T]) GetResult() {
	logs.SilentWarning("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "session",
	})
}

// Methods
func (p *run[T]) setRepository() {
	p.repository = p.getRepository()
}

func (p *run[T]) createTuiView() {
	p.cache.view = generatesteplist.Process[cmdservetui.Context]()
}

func (p *run[T]) runTuiView() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	/*
		FIXME MEDIUM de vrai etre une vue et nom l'equilavent du moteur qui tourne dans CMDRUN
		Deplaver le moteur aiderais a ameliore et fixer ici
		Ne devrai pas rester comme ça ca ca marche mais ma melanger les responsabilté
	*/
	model := cmdservetui.New(cmdservetui.StreamModelTemplate[T]{
		Context:    ctx,
		Steps:      p.cache.view.GetSteps(),
		Repository: p.repository,
		Flags:      p.getFlags(),
	})

	p.cache.view.Run(&model, tea.WithContext(ctx))
	if model.HasErr() {
		p.SetErr(model.GetErr())
	}

	if p.cache.view.HasErr() {
		p.SetErr(p.cache.view.GetErr())
	}
}

// Helpers
