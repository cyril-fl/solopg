package interfass

import (
	"solopg/app/cmdrun/types/size"

	tea "charm.land/bubbletea/v2"
)


type Page interface {
	GetView() tea.View
	SetSize(size *tea.WindowSizeMsg)
	SetBody(body string)
	SetFooter(footer string)
	GetSize() size.Size
	GetAvailableSize() size.Size
}

type Presentable interface {
	SetPageSize(size *tea.WindowSizeMsg)
}