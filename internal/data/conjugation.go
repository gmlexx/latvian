package data

import "strings"

// ConjugationExplanation returns a short learner-facing explanation for the
// conjugation class recorded on a verb. The examples use the present tense.
func (v Verb) ConjugationExplanation() []string {
	class := strings.ToLower(v.Class)
	var explanation []string

	if strings.Contains(class, "neregulārs") {
		explanation = []string{
			"Šis ir neregulārs darbības vārds: tā celms vai formas neatbilst vienai vispārīgai shēmai, tāpēc galvenās formas jāiemācās atsevišķi.",
			"Piemēri: būt — esmu, esi, ir; iet — eju, ej, iet; redzēt — redzu, redzi, redz.",
		}
	} else {
		switch {
		case strings.Contains(class, "1. konjugācija"):
			explanation = []string{
				"1. konjugācijas darbības vārdiem tagadnes formas bieži veido no mainīta vai nemainīta celma. Bieži: 1. persona vienskaitlī -u, 2. un 3. persona vienskaitlī un 3. persona daudzskaitlī bez galotnes, 1. persona daudzskaitlī -am, 2. persona daudzskaitlī -at.",
				"Piemērs: ņemt — ņemu, ņem, ņem, ņemam, ņemat, ņem.",
			}
		case strings.Contains(class, "2. konjugācija"):
			explanation = []string{
				"2. konjugācijas tagadnes celms bieži beidzas ar -ā, -ē, -ī vai -o; dažiem vārdiem pirms galotnes ir -j. Bieži: 1. persona vienskaitlī -u, 2. un 3. persona vienskaitlī un 3. persona daudzskaitlī bez galotnes, 1. persona daudzskaitlī -am, 2. persona daudzskaitlī -at.",
				"Piemēri: runāt — runāju, runā, runā, runājam, runājat, runā; spēlēt — spēlēju, spēlē, spēlē, spēlējam, spēlējat, spēlē.",
			}
		case strings.Contains(class, "3. konjugācija"):
			explanation = []string{
				"3. konjugācijai ir vairāki tagadnes celma tipi. Vienā biežā modelī galotnes ir -u, -i, -a, -ām, -āt, -a (lasu, lasi, lasa, lasām, lasāt, lasa). Citā modelī 1. persona vienskaitlī ir -u, 2. persona vienskaitlī -i, 3. persona vienskaitlī un daudzskaitlī bez galotnes, 1. persona daudzskaitlī -am un 2. persona daudzskaitlī -at (ticu, tici, tic, ticam, ticat, tic).",
				"Piemēri: lasīt — lasu, lasi, lasa; rakstīt — rakstu, raksti, raksta; ticēt — ticu, tici, tic; gulēt — guļu, guli, guļ.",
			}
		default:
			explanation = []string{"Šim darbības vārdam nav pievienots konjugācijas noteikums; pareizās formas jāapgūst kopā ar piemēriem."}
		}
	}

	if strings.Contains(class, "atgriezeniskā") {
		explanation = append(explanation, "Atgriezeniskie darbības vārdi infinitīvā beidzas ar -ties. Tagadnē pie celma pievieno atgriezeniskās galotnes -os, -ies, -as, -amies, -aties, -as; celma patskanis un garums var mainīties: mazgāties — mazgājos, mazgājies, mazgājas, mazgājamies, mazgājaties, mazgājas; mācīties — mācos, mācies, mācās.")
	}

	return explanation
}
