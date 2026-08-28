package gameboard

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

// TODO refactor
func makeOracleModel() list.Model {
	items := make([]list.Item, 0)
	for _, oracle := range gameplay.GetOracle() {
		if oracle.Visible {
			key := "oracle." + oracle.ID
			name := t.Localizer.MustLocalize(&goi18n.LocalizeConfig{
				MessageID: key,
				DefaultMessage: &goi18n.Message{
					ID:    key,
					Other: oracle.ID,
				},
			})
			items = append(items, tui.NewItem(name, "", oracle))
		}
	}

	model := list.New(items, list.NewDefaultDelegate(), panelWidth-4, oracleMenuHeight)
	tui.ConfigureList(&model)
	return model
}

func handleOracleRoll(m model) (model, tea.Cmd) {
	selected, ok := m.oracleList.SelectedItem().(tui.Item[*gameplay.Oracle])
	if !ok || selected.Value() == nil {
		return m, nil
	}

	oracle := selected.Value()
	result, err := gameplay.RollOracle[any](oracle)
	if err != nil {
		m.messages = append(m.messages, t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "error.oracle_action", TemplateData: map[string]any{"Error": err}}))
	} else {
		critical := ""
		if result.Critical {
			critical = " (" + t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "critical"}) + ")"
		}
		message := fmt.Sprintf("Oracle %s — jet de %d : %v%s", oracle.ID, result.Roll, result.Result, critical)
		m.engine.AddJournalEntry("Oracle", message)
		m.messages = append(m.messages, message)
	}

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	m.viewport.GotoBottom()
	return m, nil
}
