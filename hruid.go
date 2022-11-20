package hruid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/s-torneo/hruid-golang/models"
)

// return a random number beetwen max e min value passed as arguments
func randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// return a random element from a slice
func randomElement(slice []string) string {
	return slice[randomNumber(0, len(slice))]
}

// return a random uid with this structure: number-adjective-noun-verb-adverb
func generateUid(obj models.Entities) string {
	return fmt.Sprintf("%d-%s-%s-%s-%s",
		randomNumber(2, 33),
		randomElement(obj.Adjectives),
		randomElement(obj.Nouns),
		randomElement(obj.Verbs),
		randomElement(obj.Adverbs))
}

// return a random uid using default entities
func Generate() string {
	obj := models.Entities{
		Adjectives: models.Adjectives,
		Nouns:      models.Nouns,
		Verbs:      models.Verbs,
		Adverbs:    models.Adverbs,
	}
	return generateUid(obj)
}

// return a random uid using entities obtained from file passed as argument
func GenerateByFile(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var obj models.Entities
	err = json.Unmarshal([]byte(file), &obj)
	if err != nil {
		return "", err
	}
	return generateUid(obj), nil
}

// return a random uid using entities passed as argument
func GenerateByEntities(obj models.Entities) string {
	return generateUid(obj)
}
