// Package progress tracks which verb infinitives the user has already
// practiced across sessions, persisting the state to disk so the coverage
// goal ("eventually see every infinitive") survives program restarts.
package progress

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// fileState mirrors the JSON shape stored on disk.
type fileState struct {
	// Covered holds the set of infinitives already practiced during the
	// current coverage cycle.
	Covered map[string]bool `json:"covered"`
	// Cycles counts how many times the full verb list has been covered
	// and the tracker restarted. Purely informational.
	Cycles int `json:"cycles"`
}

// Tracker records coverage progress and persists it to disk.
type Tracker struct {
	path    string
	Covered map[string]bool
	Cycles  int
}

// progressFileName is the name of the local progress file, kept in the
// same folder the program is run from so it travels with the project.
const progressFileName = ".latvian-progress.json"

// defaultPath returns the on-disk location used to store progress: a
// dotfile in the current working directory.
func defaultPath() string {
	return progressFileName
}

// Load reads the tracker state from disk, returning a fresh empty tracker
// if no file exists yet or if it can't be parsed.
func Load() *Tracker {
	path := defaultPath()
	t := &Tracker{path: path, Covered: map[string]bool{}}

	data, err := os.ReadFile(path)
	if err != nil {
		return t
	}

	var state fileState
	if err := json.Unmarshal(data, &state); err != nil {
		return t
	}

	if state.Covered != nil {
		t.Covered = state.Covered
	}
	t.Cycles = state.Cycles
	return t
}

// Save persists the current tracker state to disk, creating any missing
// parent directories as needed.
func (t *Tracker) Save() error {
	if err := os.MkdirAll(filepath.Dir(t.path), 0o755); err != nil {
		return err
	}

	state := fileState{Covered: t.Covered, Cycles: t.Cycles}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	tmp := t.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, t.path)
}

// IsCovered reports whether the given infinitive has been practiced during
// the current coverage cycle.
func (t *Tracker) IsCovered(infinitive string) bool {
	return t.Covered[infinitive]
}

// Mark records the infinitive as practiced. It returns true if this call
// newly marked it (i.e. it wasn't already covered).
func (t *Tracker) Mark(infinitive string) bool {
	if t.Covered[infinitive] {
		return false
	}
	t.Covered[infinitive] = true
	return true
}

// CoveredCount returns how many distinct infinitives have been marked so
// far in the current cycle.
func (t *Tracker) CoveredCount() int {
	return len(t.Covered)
}

// Reset clears the current cycle's coverage and increments the completed
// cycle counter, so a new full pass over every infinitive can begin.
func (t *Tracker) Reset() {
	t.Covered = map[string]bool{}
	t.Cycles++
}

// IsComplete reports whether every infinitive out of `total` has been
// covered during the current cycle.
func (t *Tracker) IsComplete(total int) bool {
	return total > 0 && t.CoveredCount() >= total
}
