package flect

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type tt struct {
	act string
	exp string
}

func Test_LoadInflections(t *testing.T) {
	r := require.New(t)
	m := map[string]string{
		"beatle": "the beatles",
		"xyz":    "zyx",
	}

	b, err := json.Marshal(m)
	r.NoError(err)

	r.NoError(LoadInflections(bytes.NewReader(b)))

	for k, v := range m {
		r.Equal(v, Pluralize(k))
		r.Equal(v, Pluralize(v))
		r.Equal(k, Singularize(k))
		r.Equal(k, Singularize(v))
	}
}

func Test_LoadAcronyms(t *testing.T) {
	r := require.New(t)
	m := []string{
		"ACC",
		"TLC",
		"LSA",
	}

	b, err := json.Marshal(m)
	r.NoError(err)

	r.NoError(LoadAcronyms(bytes.NewReader(b)))

	for _, acronym := range m {
		r.True(baseAcronyms[acronym])
	}
}

var singlePluralAssertions = []tt{
	{"", ""},
	{"human", "humans"},
	{"movie", "movies"},
	{"ox", "oxen"},
	{"user", "users"},
	{"datum", "data"},
	{"search", "searches"},
	{"switch", "switches"},
	{"fix", "fixes"},
	{"ovum", "ova"},
	{"box", "boxes"},
	{"process", "processes"},
	{"address", "addresses"},
	{"case", "cases"},
	{"stack", "stacks"},
	{"wish", "wishes"},
	{"fish", "fish"},
	{"jeans", "jeans"},
	{"funky jeans", "funky jeans"},
	{"category", "categories"},
	{"query", "queries"},
	{"ability", "abilities"},
	{"agency", "agencies"},
	{"movie", "movies"},
	{"archive", "archives"},
	{"index", "indices"},
	{"wife", "wives"},
	{"safe", "saves"},
	{"half", "halves"},
	{"move", "moves"},
	{"salesperson", "salespeople"},
	{"person", "people"},
	{"spokesman", "spokesmen"},
	{"basis", "bases"},
	{"diagnosis", "diagnoses"},
	{"diagnosis_a", "diagnosis_as"},
	{"datum", "data"},
	{"stadium", "stadia"},
	{"analysis", "analyses"},
	{"node_child", "node_children"},
	{"child", "children"},
	{"experience", "experiences"},
	{"day", "days"},
	{"comment", "comments"},
	{"foobar", "foobars"},
	{"newsletter", "newsletters"},
	{"news", "news"},
	{"series", "series"},
	{"species", "species"},
	{"quiz", "quizzes"},
	{"perspective", "perspectives"},
	{"ox", "oxen"},
	{"photo", "photos"},
	{"buffalo", "buffaloes"},
	{"tomato", "tomatoes"},
	{"dwarf", "dwarves"},
	{"elf", "elves"},
	{"information", "information"},
	{"equipment", "equipment"},
	{"bus", "buses"},
	{"status", "statuses"},
	{"Status", "Statuses"},
	{"status_code", "status_codes"},
	{"mouse", "mice"},
	{"louse", "lice"},
	{"house", "houses"},
	{"spouse", "spouses"},
	{"octopus", "octopi"},
	{"virus", "viri"},
	{"alias", "aliases"},
	{"portfolio", "portfolios"},
	{"matrix", "matrices"},
	{"axis", "axes"},
	{"testis", "testes"},
	{"crisis", "crises"},
	{"rice", "rice"},
	{"shoe", "shoes"},
	{"horse", "horses"},
	{"prize", "prizes"},
	{"edge", "edges"},
	{"database", "databases"},
	{"circus", "circuses"},
	{"plus", "pluses"},
	{"fuse", "fuses"},
	{"prometheus", "prometheuses"},
	{"field", "fields"},
	{"custom_field", "custom_fields"},

	// https://www.grammarly.com/blog/plural-nouns/
	{"cat", "cats"},
	{"house", "houses"},
	{"truss", "trusses"},
	{"bus", "buses"},
	{"marsh", "marshes"},
	{"lunch", "lunches"},
	{"tax", "taxes"},
	{"blitz", "blitzes"},
	{"fez", "fezzes"},
	{"gas", "gasses"},
	{"wife", "wives"},
	{"wolf", "wolves"},
	{"roof", "roofs"},
	{"belief", "beliefs"},
	{"chef", "chefs"},
	{"chief", "chiefs"},
	{"city", "cities"},
	{"puppy", "puppies"},
	{"ray", "rays"},
	{"boy", "boys"},
	{"potato", "potatoes"},
	{"tomato", "tomatoes"},
	{"photo", "photos"},
	{"piano", "pianos"},
	{"halo", "halos"},
	{"cactus", "cacti"},
	{"focus", "foci"},
	{"analysis", "analyses"},
	{"ellipsis", "ellipses"},
	{"phenomenon", "phenomena"},
	{"criterion", "criteria"},
	{"sheep", "sheep"},
	{"series", "series"},
	{"species", "species"},
	{"deer", "deer"},
	{"child", "children"},
	{"goose", "geese"},
	{"man", "men"},
	{"woman", "women"},
	{"tooth", "teeth"},
	{"foot", "feet"},
	{"mouse", "mice"},
	{"person", "people"},

	// more tests
	{"video", "videos"},
	{"Cat", "Cats"},
	{"post video", "post videos"},
	{"Post Video", "Post Videos"},
	{"Person", "People"},
	{"PERSON", "PEOPLE"},
	{"FancyPerson", "FancyPeople"},
	{"FANCYPERSON", "FANCYPEOPLE"},
	{"the_bus", "the_buses"},
	{"three_sheep", "three_sheep"},
	{"hello", "hellos"},
}
var pluralSingularAssertions = []tt{}

func init() {
	for k, v := range singleToPlural {
		singlePluralAssertions = append(singlePluralAssertions, tt{k, v})

		// add some variations
		// singlePluralAssertions = append(singlePluralAssertions, tt{strings.ToUpper(k), v})
		// singlePluralAssertions = append(singlePluralAssertions, tt{strings.ToLower(k), v})
		// for i, x := range k {
		// 	n := k[:i] + strings.ToLower(string(x)) + k[i+1:]
		// 	singlePluralAssertions = append(singlePluralAssertions, tt{n, v})
		//
		// 	n = k[:i] + strings.ToUpper(string(x)) + k[i+1:]
		// 	singlePluralAssertions = append(singlePluralAssertions, tt{n, v})
		// }
	}

	for _, a := range singlePluralAssertions {
		pluralSingularAssertions = append(pluralSingularAssertions, tt{act: a.exp, exp: a.act})
	}
}
