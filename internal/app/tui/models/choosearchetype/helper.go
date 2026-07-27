package choosearchetype

// func makeItems[T archetypes.Archetype](data []T) []list.Item {
// 	items := make([]list.Item, 0, len(data))
// 	for _, i := range data {
// 		if !i.IsPlayable() {
// 			continue
// 		}

// 		newItem := tui.NewItem(i.GetName(), "", i)

// 		items = append(items, newItem)
// 	}

// 	return items
// }

// func makeModel[T archetypes.Archetype](data []T) list.Model {
// 	items := makeItems(data)
// 	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
// 	tui.ConfigureList(&model)
// 	return model
// }
