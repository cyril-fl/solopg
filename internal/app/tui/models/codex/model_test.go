package codex

import (
	"testing"

	"solopg/internal/app/tui/models/codexform"
)

func TestFormKindMapsEveryCodexPage(t *testing.T) {
	tests := []struct {
		name string
		kind CodexKind
		want codexform.Kind
	}{
		{name: "npcs", kind: CodexNPCs, want: codexform.NPCs},
		{name: "monsters", kind: CodexMonsters, want: codexform.Monsters},
		{name: "locations", kind: CodexLocations, want: codexform.Locations},
		{name: "objects", kind: CodexObjects, want: codexform.Objects},
		{name: "objectives", kind: CodexObjectifs, want: codexform.Objectifs},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := Model{active: tt.kind}
			if got := model.formKind(); got != tt.want {
				t.Fatalf("formKind() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestScreenStateHelpers(t *testing.T) {
	model := Model{screen: ScreenClosed}
	if model.PageOpen() || model.FormOpen() || model.HandlesEscape() {
		t.Fatal("closed model should not report an open screen")
	}

	model.screen = ScreenPage
	if !model.PageOpen() || model.FormOpen() || model.HandlesEscape() {
		t.Fatal("page model reported an invalid screen state")
	}

	model.screen = ScreenForm
	if model.PageOpen() || !model.FormOpen() || !model.HandlesEscape() {
		t.Fatal("form model reported an invalid screen state")
	}
}
