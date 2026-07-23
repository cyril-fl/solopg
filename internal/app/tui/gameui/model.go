package gameui

import (
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

const footerLines = 1

type screen int

const (
	screenMain screen = iota
	screenSettings
)

type model struct {
	itemList list.Model
	width    int
	height   int
	screen   screen
}

func (root model) Init() tea.Cmd {
	return nil
}

func (root model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if winSize, ok := msg.(tea.WindowSizeMsg); ok {
		root.width = winSize.Width
		root.height = winSize.Height
		root.syncListSize()
		return root, nil
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.String() == tui.KeyCtrlC {
			return root, tea.Quit
		}
		updated, cmd, skipListUpdate := root.routeKey(keyMsg.String())
		if skipListUpdate {
			return updated, cmd
		}
		root = updated
	}

	return root.updateItemList(msg)
}

// routeKey returns skipListUpdate=true when Update must return without delegating to the list
// (handled shortcut, or unhandled key on the settings screen).
func (root model) routeKey(key string) (updated model, cmd tea.Cmd, skipListUpdate bool) {
	if root.screen == screenSettings {
		if key == tui.KeyQuit {
			return root, tea.Quit, true
		}
		if key == tui.KeyBack || key == tui.KeyEsc {
			root.screen = screenMain
			root.syncListSize()
			return root, nil, true
		}
		return root, nil, true
	}

	if key == tui.KeySettings {
		root.screen = screenSettings
		return root, nil, true
	}

	return root, nil, false
}

func (root model) updateItemList(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	root.itemList, cmd = root.itemList.Update(msg)
	return root, cmd
}

func (root *model) syncListSize() {
	if root.height < footerLines+1 {
		root.itemList.SetSize(max(0, root.width), 0)
		return
	}
	root.itemList.SetSize(root.width, root.height-footerLines)
}
