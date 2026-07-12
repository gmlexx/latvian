package data

import (
	_ "embed"
	"encoding/json"
)

//go:embed verbs.json
var verbsJSON []byte

// verbJSON mirrors the shape of entries in verbs.json.
type verbJSON struct {
	Infinitive  string       `json:"infinitive"`
	Translation string       `json:"translation"`
	Class       string       `json:"class"`
	Object      string       `json:"object"`
	Forms       [3][6]string `json:"forms"`
}

// Verbs is a curated set of common Latvian verbs with fully verified
// conjugation tables (present, past, future; all six person/number forms).
//
// The list intentionally favors correctness over raw count: each entry has
// been checked against standard Latvian conjugation patterns. The data
// itself lives in verbs.json - to grow the practice set toward the full
// "200 most popular verbs" goal, add more entries there following the same
// shape; the engine and UI need no changes to support more entries.
var Verbs = loadVerbs()

func loadVerbs() []Verb {
	var raw []verbJSON
	if err := json.Unmarshal(verbsJSON, &raw); err != nil {
		panic("data: failed to parse verbs.json: " + err.Error())
	}

	verbs := make([]Verb, len(raw))
	for i, v := range raw {
		verbs[i] = Verb{
			Infinitive:  v.Infinitive,
			Translation: v.Translation,
			Class:       v.Class,
			Object:      v.Object,
			Forms:       v.Forms,
		}
	}
	return verbs
}
