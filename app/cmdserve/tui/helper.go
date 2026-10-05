package cmdservetui

import (
	"solopg/app/shared/services/logs"
)

// NOTE WARNING Pas sur que le codi ici fonctionne.
func (m *model[T]) handleStreamError(err error) {
	if m.hasSameStringError(err) {

		return
	}

	m.errorset.Add(err)

	logs.Error("error.unexpected:action", map[string]any{
		"Action": "unexpected:action.logs:display",
		"Error":  err.Error(),
	})
}

func (m model[T]) hasSameStringError(err error) bool {
	for _, stored := range m.errorset.ToSlice() {
		if stored.Error() == err.Error() {
			return true
		}
	}

	return false
}

// ---

func (m *model[T]) assertFlag() {
	m.assertFilterFlag()
}

func (m *model[T]) assertFilterFlag() {
	if filter, err := m.flags.GetString("filter"); err != nil {
		return
	} else if isKind := logs.AssertKind(logs.Kind(filter)); isKind {
		return
	} else {
		logs.Error("error.unexpected:action", map[string]any{
			"Action": "unexpected:action.logs:filter",
			"Error": logs.CeaseError("error.invalid", map[string]any{
				"Subject": "filter",
			}),
		})
	}
}
