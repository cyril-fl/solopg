package cmdservetui

import (
	"context"
	"errors"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

type fakeWatchable[T any] struct {
	stream <-chan T
	err    error
	calls  int
	ctx    context.Context
}

func (f *fakeWatchable[T]) Watch(ctx context.Context) (<-chan T, error) {
	f.calls++
	f.ctx = ctx
	return f.stream, f.err
}

func TestModelReceivesEventsWithoutOpeningMultipleStreams(t *testing.T) {
	stream := make(chan int, 1)
	stream <- 42

	repo := &fakeWatchable[int]{stream: stream}
	m := NewModel(context.Background(), nil, repo)

	ready := m.Init()()
	_, wait := m.Update(ready)
	event := wait()
	_, next := m.Update(event)

	if got := m.events; len(got) != 1 || got[0] != 42 {
		t.Fatalf("unexpected events: %v", got)
	}
	if repo.calls != 1 {
		t.Fatalf("Watch called %d times, want 1", repo.calls)
	}
	if next == nil {
		t.Fatal("expected a command waiting for the next event")
	}
}

func TestModelReturnsErrorWhenWatchFails(t *testing.T) {
	want := errors.New("watch failed")
	repo := &fakeWatchable[int]{err: want}
	m := NewModel(context.Background(), nil, repo)

	msg := m.Init()()
	_, quit := m.Update(msg)

	if !errors.Is(m.Err(), want) {
		t.Fatalf("model error = %v, want %v", m.Err(), want)
	}
	if _, ok := quit().(tea.QuitMsg); !ok {
		t.Fatal("expected tea.Quit after Watch failure")
	}
}

func TestModelReturnsErrorWhenStreamCloses(t *testing.T) {
	stream := make(chan int)
	close(stream)
	repo := &fakeWatchable[int]{stream: stream}
	m := NewModel(context.Background(), nil, repo)

	ready := m.Init()()
	_, wait := m.Update(ready)
	closed := wait()
	_, quit := m.Update(closed)

	if m.Err() == nil {
		t.Fatal("expected an error when the stream closes")
	}
	if _, ok := quit().(tea.QuitMsg); !ok {
		t.Fatal("expected tea.Quit after stream closure")
	}
}

func TestViewDoesNotOpenTheRepositoryStream(t *testing.T) {
	repo := &fakeWatchable[int]{}
	m := NewModel(context.Background(), nil, repo)

	_ = m.View()

	if repo.calls != 0 {
		t.Fatalf("View called Watch %d times", repo.calls)
	}
}

func TestViewRendersEventsAndErrors(t *testing.T) {
	repo := &fakeWatchable[string]{}
	m := NewModel(context.Background(), nil, repo)
	m.events = []string{"first", "second"}
	m.err = errors.New("stream failed")

	content := m.View().Content
	if !strings.Contains(content, "first") || !strings.Contains(content, "second") {
		t.Fatalf("view does not contain events: %q", content)
	}
	if !strings.Contains(content, "stream failed") {
		t.Fatalf("view does not contain the error: %q", content)
	}
}

func TestModelQuitsOnQ(t *testing.T) {
	m := NewModel(context.Background(), nil, &fakeWatchable[int]{})

	_, quit := m.Update(tea.KeyPressMsg(tea.Key{Text: "q"}))

	if _, ok := quit().(tea.QuitMsg); !ok {
		t.Fatal("expected tea.Quit after q")
	}
}

func TestModelUsesTheProvidedContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := &fakeWatchable[int]{stream: make(chan int)}
	m := NewModel(ctx, nil, repo)
	_ = m.Init()()

	if repo.ctx != ctx {
		t.Fatal("Watch did not receive the model context")
	}

	cancel()
	select {
	case <-repo.ctx.Done():
	default:
		t.Fatal("provided context was not cancelled")
	}
}
