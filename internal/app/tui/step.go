package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type Step struct {
	Model   tea.Model
	Resolve func(ctx *Context, value any) error
}

type step int

const (
	step1 step = iota
	step2
	step3
)

func resolveStep(m model, msg tea.Msg) {
	resolution, ok := msg.(ResolutionMsg)
	skip := !ok || resolution.Err != nil || !resolution.Completed

	if skip {
		return
	}

	 m.steps[m.step].Resolve(&m.context, resolution.Value)
}


func changeStep(m model, msg tea.Msg) (model, tea.Cmd) {
	resolution, ok := msg.(ResolutionMsg)
	if !ok {
		return m, nil
	}

	if resolution.Err != nil {
		return m, tea.Quit
	}

	if !resolution.Completed {
		return m, nil
	}


if m.step < step(len(m.steps)-1) {
	m.step++

	fmt.Println("NEW STEP", m.step)

	return m, m.current().Init()
}
	return m, tea.Quit
}
