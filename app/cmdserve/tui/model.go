package cmdservetui

import (
	"context"
	"errors"

	"solopg/app/shared/services/mango/repository"
	interfass "solopg/app/shared/types/interface"
	"solopg/app/shared/types/step"

	tea "charm.land/bubbletea/v2"
)

type Context struct{}

type contextStepList = step.List[Context]
type Step = step.Step[Context]

type model[T interfass.Stringable] struct {
	ctx        context.Context
	repository repository.Watchable[T]
	events     []T
	stream     <-chan T

	err  error
	size *tea.WindowSizeMsg
}

type streamReadyMsg[T any] struct {
	stream <-chan T
}

type streamEventMsg[T any] struct {
	event T
}

type streamErrorMsg struct {
	err error
}

type streamClosedMsg struct{}

func NewModel[T interfass.Stringable](ctx context.Context, steps *contextStepList, repository repository.Watchable[T]) model[T] {
	return model[T]{
		ctx:        ctx,
		repository: repository,
	}
}

func (m *model[T]) Init() tea.Cmd {
	return m.startWatch()
}

func (m *model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case streamReadyMsg[T]:
		m.stream = msg.stream
		return m, waitForEvent(m.stream)
	case streamEventMsg[T]:
		m.events = append(m.events, msg.event)
		return m, waitForEvent(m.stream)

	case streamErrorMsg:
		m.err = msg.err
		return m, tea.Quit
	case streamClosedMsg:
		m.err = errors.New("MongoDB stream closed unexpectedly")
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.size = &msg
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		// TODO mettre des keyevent
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model[T]) Err() error {
	return m.err
}

func (m *model[T]) startWatch() tea.Cmd {
	return func() tea.Msg {
		stream, err := m.repository.Watch(m.ctx)
		if err != nil {
			return streamErrorMsg{err: err}
		}

		return streamReadyMsg[T]{stream: stream}
	}
}

func waitForEvent[T any](stream <-chan T) tea.Cmd {
	return func() tea.Msg {
		event, ok := <-stream
		if !ok {
			return streamClosedMsg{}
		}

		return streamEventMsg[T]{event: event}
	}
}
