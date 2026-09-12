package form

import "strings"

func (m *Model) View() string {
	fields := make([]string, len(m.fields))
	
	for i, field := range m.fields {
		fields[i] = field.Label() + ": " + field.Value().(string)
	}
	
	return "Form:\n" + strings.Join(fields, "\n")	
}
