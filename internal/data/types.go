package data

import "math/rand"

// Tense represents one of the three main Latvian tenses practiced here.
type Tense int

const (
	Present Tense = iota
	Past
	Future
)

func (t Tense) String() string {
	switch t {
	case Present:
		return "tagadne (present)"
	case Past:
		return "pagātne (past)"
	case Future:
		return "nākotne (future)"
	}
	return "?"
}

// PersonNumber indexes the six person/number combinations of Latvian,
// in the conventional conjugation table order.
type PersonNumber int

const (
	Es   PersonNumber = iota // 1st sg.
	Tu                       // 2nd sg.
	Vins                     // 3rd sg. (viņš/viņa)
	Mes                      // 1st pl.
	Jus                      // 2nd pl.
	Vini                     // 3rd pl. (viņi/viņas)
)

// Subject is the pronoun/label shown in the exercise sentence for a given
// person/number slot. Third person alternates between masc/fem randomly.
var subjectsMasc = [6]string{"Es", "Tu", "Viņš", "Mēs", "Jūs", "Viņi"}
var subjectsFem = [6]string{"Es", "Tu", "Viņa", "Mēs", "Jūs", "Viņas"}

func (p PersonNumber) Label(fem bool) string {
	if fem {
		return subjectsFem[p]
	}
	return subjectsMasc[p]
}

func (p PersonNumber) Description() string {
	switch p {
	case Es:
		return "1. persona vsk. (es)"
	case Tu:
		return "2. persona vsk. (tu)"
	case Vins:
		return "3. persona vsk. (viņš/viņa)"
	case Mes:
		return "1. persona dsk. (mēs)"
	case Jus:
		return "2. persona dsk. (jūs)"
	case Vini:
		return "3. persona dsk. (viņi/viņas)"
	}
	return "?"
}

// Verb holds full conjugation data for one infinitive, across the three
// tenses and six person/number slots, plus metadata used to build practice
// sentences and grammar explanations.
type Verb struct {
	Infinitive  string   // e.g. "iegūt"
	Translation string   // English gloss, e.g. "to obtain"
	Class       string   // short conjugation-class note shown in explanations
	Objects     []string // trailing sentence contexts, e.g. "jaunus draugus"
	Forms       [3][6]string
}

// Form returns the correct conjugated form for a tense/person slot.
func (v Verb) Form(t Tense, p PersonNumber) string {
	return v.Forms[t][p]
}

// RandomObject returns one of the verb's sentence-context objects at
// random, or an empty string if none are defined.
func (v Verb) RandomObject() string {
	if len(v.Objects) == 0 {
		return ""
	}
	return v.Objects[rand.Intn(len(v.Objects))]
}

// TimeAdverb returns the natural time word used per tense in sentences.
func TimeAdverb(t Tense) string {
	switch t {
	case Present:
		return "šodien"
	case Past:
		return "vakar"
	case Future:
		return "rīt"
	}
	return ""
}

// AcceptableTenses returns every tense whose conjugated form is a legitimate
// answer for the given adverb/tense slot. Latvian "šodien" (today) is
// ambiguous: it can describe something ongoing/habitual (present tense) or
// something that will still happen later that day (future tense), so both
// are accepted. "vakar" (yesterday) and "rīt" (tomorrow) are unambiguous.
func (t Tense) AcceptableTenses() []Tense {
	if t == Present {
		return []Tense{Present, Future}
	}
	return []Tense{t}
}
