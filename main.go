package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"latvian-grammar/internal/data"
	"latvian-grammar/internal/progress"
)

const sessionSize = 10

// question is one exercise instance: a verb form to conjugate in a given
// tense/person/gender slot, tagged with its original position in the
// 10-question session (1-based) so retries can reference the same number.
type question struct {
	num    int // 1..sessionSize, stable across retries
	verb   data.Verb
	object string
	tense  data.Tense
	person data.PersonNumber
	fem    bool
}

// randomTense picks a tense weighted to roughly match how often each tense
// is actually used in everyday Latvian speech/writing: present tense
// dominates, past tense is fairly common, and future tense is comparatively
// rare.
func randomTense() data.Tense {
	switch n := rand.Intn(100); {
	case n < 55:
		return data.Present
	case n < 90:
		return data.Past
	default:
		return data.Future
	}
}

// pickVerb chooses a verb for the next question. To make sure every
// infinitive in the dataset eventually gets practiced, it prefers verbs the
// tracker hasn't covered yet in the current cycle; once every verb has been
// covered, it falls back to picking from the whole list uniformly.
func pickVerb(tracker *progress.Tracker) data.Verb {
	var uncovered []data.Verb
	for _, v := range data.Verbs {
		if !tracker.IsCovered(v.Infinitive) {
			uncovered = append(uncovered, v)
		}
	}
	if len(uncovered) == 0 {
		return data.Verbs[rand.Intn(len(data.Verbs))]
	}
	return uncovered[rand.Intn(len(uncovered))]
}

func newQuestion(num int, tracker *progress.Tracker) question {
	verb := pickVerb(tracker)
	return question{
		num:    num,
		verb:   verb,
		object: verb.RandomObject(),
		tense:  randomTense(),
		person: data.PersonNumber(rand.Intn(6)),
		fem:    rand.Intn(2) == 1,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	tracker := progress.Load()
	totalVerbs := len(data.Verbs)
	if tracker.IsComplete(totalVerbs) {
		fmt.Println(green(bold(fmt.Sprintf(
			"Apsveicam! Iepriekšējā ciklā tika praktizēti visi %d darbības vārdi. Sākam jaunu ciklu (Nr.%d).",
			totalVerbs, tracker.Cycles+1))))
		tracker.Reset()
		if err := tracker.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Brīdinājums: neizdevās saglabāt progresu: %v\n", err)
		}
		fmt.Println()
	}

	fmt.Println(bold("=== Latviešu valodas darbības vārdu treniņš ==="))
	fmt.Printf("Šai kārtā ir %d jautājumi. Nepareizi atbildētie jautājumi tiks atkārtoti beigās,\n", sessionSize)
	fmt.Println("līdz visi ir atbildēti pareizi.")
	fmt.Println("Ievadiet pareizo darbības vārda formu tukšajā vietā.")
	fmt.Println("Rakstiet 'exit' vai 'quit', lai beigtu jebkurā brīdī.")
	fmt.Printf("Kopējais progress: %d/%d darbības vārdi jau praktizēti šajā ciklā (cikls Nr.%d).\n",
		tracker.CoveredCount(), totalVerbs, tracker.Cycles+1)
	fmt.Println()

	queue := make([]question, sessionSize)
	for i := range queue {
		q := newQuestion(i+1, tracker)
		if tracker.Mark(q.verb.Infinitive) {
			if err := tracker.Save(); err != nil {
				fmt.Fprintf(os.Stderr, "Brīdinājums: neizdevās saglabāt progresu: %v\n", err)
			}
		}
		queue[i] = q
	}

	attempts := make(map[int]int) // question num -> attempts so far
	solved := make(map[int]bool)  // question num -> answered correctly at least once
	totalAttempts, totalCorrectAttempts := 0, 0
	quit := false

	for len(queue) > 0 && !quit {
		q := queue[0]
		queue = queue[1:]
		attempts[q.num]++

		remaining := len(queue) + 1 // this one plus what's left in queue
		retryTag := ""
		if attempts[q.num] > 1 {
			retryTag = yellow(fmt.Sprintf(" [atkārtojums, mēģinājums %d]", attempts[q.num]))
		}
		fmt.Printf("%s%s\n", bold(fmt.Sprintf("Jautājums Nr.%d", q.num)), retryTag)
		fmt.Printf("(atlikuši šai kārtā: %d, kopā jautājumu: %d)\n", remaining, sessionSize)

		subject := q.person.Label(q.fem)
		adverb := data.TimeAdverb(q.tense)
		context := q.object
		if context == "" {
			fmt.Printf("%s %s ___ (%s).\n", subject, adverb, cyan(q.verb.Infinitive))
		} else {
			fmt.Printf("%s %s ___ (%s) %s.\n", subject, adverb, cyan(q.verb.Infinitive), context)
		}
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nInput closed, exiting.")
			quit = true
			break
		}
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "exit") || strings.EqualFold(input, "quit") {
			quit = true
			break
		}
		if input == "" {
			fmt.Println("Lūdzu, ierakstiet atbildi (vai 'exit').")
			queue = append([]question{q}, queue...) // put back at front, doesn't count as an attempt
			attempts[q.num]--
			continue
		}

		totalAttempts++
		want := q.verb.Form(q.tense, q.person)
		acceptableTenses := q.tense.AcceptableTenses()

		var matchedTense data.Tense
		matched := false
		for _, t := range acceptableTenses {
			if strings.EqualFold(input, q.verb.Form(t, q.person)) {
				matched = true
				matchedTense = t
				break
			}
		}

		if matched {
			totalCorrectAttempts++
			solved[q.num] = true
			if matchedTense == q.tense {
				fmt.Println(green(bold("✓ Pareizi!")))
			} else {
				fmt.Printf("%s %s\n", green(bold("✓ Pareizi!")),
					yellow(fmt.Sprintf("(pieņemts arī %s, jo \"%s\" pieļauj abas formas)", matchedTense.String(), adverb)))
			}
			fmt.Println()
			continue
		}

		fmt.Println(red(bold("✗ Nepareizi.")))
		fmt.Printf("Tava atbilde: %s\n", red(input))
		if len(acceptableTenses) > 1 {
			alts := make([]string, len(acceptableTenses))
			for i, t := range acceptableTenses {
				alts[i] = fmt.Sprintf("%s (%s)", green(bold(q.verb.Form(t, q.person))), t.String())
			}
			fmt.Printf("Pareizās formas: %s\n", strings.Join(alts, " vai "))
		} else {
			fmt.Printf("Pareizā forma: %s\n", green(bold(want)))
		}
		fmt.Printf("Skaidrojums: darbības vārds %s (%s, %s) laika formā %s prasa %s.\n",
			cyan(q.verb.Infinitive), q.verb.Translation, yellow(q.verb.Class), yellow(q.tense.String()), yellow(q.person.Description()))
		fmt.Printf("Konjugācijas noteikums (%s):\n", yellow(q.verb.Class))
		for _, line := range q.verb.ConjugationExplanation() {
			fmt.Printf("  %s\n", line)
		}
		if context == "" {
			fmt.Printf("Pilns teikums: %s %s %s.\n", subject, adverb, green(want))
		} else {
			fmt.Printf("Pilns teikums: %s %s %s %s.\n", subject, adverb, green(want), context)
		}
		fmt.Println(yellow("Šis jautājums tiks atkārtots kārtas beigās."))
		fmt.Println()

		queue = append(queue, q)
	}

	fmt.Println()
	fmt.Println(bold("=== Rezultāts ==="))
	if quit {
		fmt.Printf("Sesija pārtraukta. Atrisināti %d/%d jautājumi pirms iziešanas.\n", len(solved), sessionSize)
	} else {
		fmt.Printf("Visi %d jautājumi atrisināti pareizi!\n", sessionSize)
	}

	if totalAttempts > 0 {
		pct := float64(totalCorrectAttempts) / float64(totalAttempts) * 100
		line := fmt.Sprintf("Kopā mēģinājumu: %d, no tiem pareizi: %d (%.0f%%)", totalAttempts, totalCorrectAttempts, pct)
		if pct >= 70 {
			fmt.Println(green(line))
		} else if pct >= 40 {
			fmt.Println(yellow(line))
		} else {
			fmt.Println(red(line))
		}

		firstTry := 0
		for n := 1; n <= sessionSize; n++ {
			if solved[n] && attempts[n] == 1 {
				firstTry++
			}
		}
		fmt.Printf("Jautājumi, kas atbildēti pareizi no pirmās reizes: %d/%d\n", firstTry, sessionSize)
	}

	fmt.Printf("Kopējais progress: %d/%d darbības vārdi praktizēti šajā ciklā (cikls Nr.%d).\n",
		tracker.CoveredCount(), totalVerbs, tracker.Cycles+1)
	if tracker.IsComplete(totalVerbs) {
		fmt.Println(green(bold(fmt.Sprintf(
			"Viss %d darbības vārdu klāsts ir praktizēts! Nākamajā palaišanas reizē sāksies jauns cikls.", totalVerbs))))
	}

	fmt.Println("Uz redzēšanos!")
}
