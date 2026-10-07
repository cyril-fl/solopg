package timer

import "time"

type Record struct {
	Start int64
	Stop  int64
}

type Timer struct {
	init      bool
	running   bool
	timeStart []int64 // unix milisecond TODO faire une note
	timeRun   []int64
}

func New() *Timer {
	return &Timer{
		init:      false,
		running:   false,
		timeStart: []int64{},
		timeRun:   []int64{},
	}
}

// Getters & Setters
func (t *Timer) GetHistory() []Record {
	if !t.init {
		return []Record{}
	}

	history := t.makeHistory()

	if t.running {
		history = append(history, Record{
			Start: getlast(t.timeStart),
			Stop:  getNow(),
		})
	}

	return history
}

func (t *Timer) makeHistory() []Record {
	if !t.init {
		return []Record{}
	}
	maxlength := len(t.timeRun)
	history := make([]Record, maxlength)

	for i, start := range t.timeStart {
		if i >= maxlength {
			break
		}

		history[i] = Record{
			Start: start,
			Stop:  start + t.timeRun[i],
		}
	}
	return history
}

func (t *Timer) GetLastTimeStop() int64 {
	if !t.init {
		return 0
	}
	if t.running {
		return getNow()
	}

	return getlast(t.timeStart) + getlast(t.timeRun)
}

// Methods
func (t *Timer) Start() {
	if t.init {
		return
	}
	t.init = true
	t.running = true
	t.timeStart = append(t.timeStart, getNow())
}

func (t *Timer) Pause() {
	if !t.init {
		return
	}
	if !t.running {
		return
	}
	run := getNow() - getlast(t.timeStart)
	t.timeRun = append(t.timeRun, run)
	t.running = false
}

func (t *Timer) Resume() {
	if !t.init {
		return
	}
	if t.running {
		return
	}
	t.timeStart = append(t.timeStart, getNow())
	t.running = true
}

// Helpers
func getNow() int64 {
	return time.Now().UnixMilli()
}

/*
func getfirst[T any](arr []T) T {
	if len(arr) > 0 {
		return arr[0]
	}
	var zero T
	return zero
}
*/

func getlast[T any](arr []T) T {
	if len(arr) > 0 {
		return arr[len(arr)-1]
	}
	var zero T
	return zero
}

func getRunningTime(history []Record) int64 {
	var total int64
	for _, r := range history {
		total += r.Stop - r.Start
	}
	return total
}
