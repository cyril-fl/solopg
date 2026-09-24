package models

import (
	"solopg/app/services/t"
	"solopg/app/tui"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type RerollModel[T any] struct {
	Attempt int
	Limit   int
	Options list.Model
	Value   T
	Draw    func() (T, error)
}

func NewRerollModel[T any](options list.Model, draw func() (T, error), limit int) RerollModel[T] {
	value, err := draw()
	if err != nil {
		panic("Failed to draw initial value for RerollModel: " + err.Error())
	}

	return RerollModel[T]{
		Attempt: 1,
		Limit:   limit,
		Options: options,
		Value:   value,
		Draw:    draw,
	}
}

func (s *RerollModel[T]) Reroll() {
	value, err := s.Draw()
	if err != nil {
		panic("Failed to draw value for RerollModel: " + err.Error())
	}

	s.Attempt++
	s.Value = value

	s.checkAttemptLimit()
}

func (s *RerollModel[T]) checkAttemptLimit() {
	if s.Attempt < s.Limit {
		return
	}

	s.Options = makeConfirmationModel(s.Options)
}

func (s *RerollModel[T]) IsOutOfLimit() bool {
	return s.Attempt >= s.Limit
}

func (s *RerollModel[T]) HandleEnterInput () tea.Cmd {
	isSelected, ok := s.Options.SelectedItem().(Item[bool])
	if !ok {
		return nil
	}

	if isSelected.Value() || s.IsOutOfLimit() {
		return  func() tea.Msg {
			return tui.ResolutionMsg{Completed: true, Value: s.Value}
		}
	}

	s.Reroll()

	return nil
}

func NewOptionsModel(options rerollOptions) list.Model {
	items := []list.Item{
		NewItem(t.Localize(options.true), "", true),
		NewItem(t.Localize(options.false), "", false),
	}

	return newConfiguredList(items, 0, 0)
}

func makeConfirmationModel(m list.Model) list.Model {
	items := []list.Item{
		NewItem(t.Localize("accept"), "", true),
	}

	return newConfiguredList(items, m.Width(), m.Height())
}

func newConfiguredList(items []list.Item, width, height int) list.Model {
	model := list.New(items, list.NewDefaultDelegate(), width, height)
	ConfigureList(&model)
	return model
}

type rerollOptions struct {
	true  string
	false string
}

var DefaultRerollOptions = rerollOptions{
	true:  "accept",
	false: "reroll",
}
