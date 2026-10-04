package generatesteplist

import (
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/process"
	"solopg/app/shared/types/step"

	tea "charm.land/bubbletea/v2"
)

type Generator[T any] struct {
	process.Process
	steps step.List[T]
}

func Process[T any]() *Generator[T] {
	return &Generator[T]{}
}
func (p *Generator[T]) GetSteps() *step.List[T] {
	return &p.steps
}

func (p *Generator[T]) GetResult() {
// TODO HIGH Remplacer par un warning
	logs.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "session",
	})
}

func (p *Generator[T]) Run(model tea.Model, opts ...tea.ProgramOption) {
	program := tea.NewProgram(model, opts...)
	_, err := program.Run()

	p.SetErr(err)
}

func (p *Generator[T]) Add(steps ...step.Step[T]) *Generator[T] {
	p.steps.Steps = append(p.steps.Steps, steps...)
	return p
}

func (p *Generator[T]) Insert(steps ...step.Step[T]) *Generator[T] {
	if len(steps) == 0 {
		return p
	}

	insertAt := p.steps.CurrentIndex + 1
	if insertAt < 0 {
		insertAt = 0
	}
	if insertAt > len(p.steps.Steps) {
		insertAt = len(p.steps.Steps)
	}

	oldSteps := p.steps.Steps
	newSteps := make([]step.Step[T], 0, len(oldSteps)+len(steps))

	newSteps = append(newSteps, oldSteps[:insertAt]...)
	newSteps = append(newSteps, steps...)
	newSteps = append(newSteps, oldSteps[insertAt:]...)

	p.steps.Steps = newSteps
	return p
}
