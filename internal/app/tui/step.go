package tui

import (
	tea "charm.land/bubbletea/v2"
)

type Step struct {
	Model   tea.Model
	Resolve func(ctx *Context, value any) error
	Skip    func(ctx *Context) bool
}

type step int

const (
	step1 step = iota
	step2
	step3
)

func resolveStep(m *model, value any) error {
	if len(m.steps) == 0 || int(m.step) >= len(m.steps) {
		return nil
	}
	if m.steps[m.step].Resolve == nil {
		return nil
	}
	return m.steps[m.step].Resolve(&m.context, value)
}

func advanceStep(m *model) (tea.Model, tea.Cmd) {
	for int(m.step) < len(m.steps)-1 {
		m.step++
		if m.steps[m.step].Skip == nil || !m.steps[m.step].Skip(&m.context) {
			break
		}
	}
	if int(m.step) >= len(m.steps)-1 && m.steps[m.step].Skip != nil && m.steps[m.step].Skip(&m.context) {
		return *m, tea.Quit
	}

	init := m.current().Init()
	if m.size == nil {
		return *m, init
	}

	resize := func() tea.Msg {
		return *m.size
	}
	return *m, tea.Sequence(init, resize)
}
