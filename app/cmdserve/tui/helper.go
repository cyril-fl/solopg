package cmdservetui

import (
	"solopg/app/shared/services/logs"
)

/*
	FIXME WARNING The following function is a temporary solution to handle errors in the stream.
	It is not a permanent solution and should be replaced with a more robust error handling mechanism in the future.
*/
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
	} else if filter == "" {
	} else if isKind := logs.AssertKind(logs.Kind(filter)); isKind {
	} else {
		logs.Error("error.unexpected:action", map[string]any{
			"Action": "unexpected:action.logs:filter",
			"Error": logs.CeaseError("error.invalid", map[string]any{
				"Subject": "filter",
			}),
		})
	}
}
