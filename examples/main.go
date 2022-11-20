package main

import (
	"fmt"
	"log"

	"github.com/s-torneo/hruid-golang"
	"github.com/s-torneo/hruid-golang/models"
)

func main() {
	// generate with default entities values
	fmt.Println(hruid.Generate())
	// generate passing entities values in a file
	uid, err := hruid.GenerateByFile("entities_values.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(uid)
	// generate passing entities values
	obj := models.Entities{
		Adjectives: []string{"beautiful", "ugly", "big", "small"},
		Nouns:      []string{"cat", "dog", "pig", "cow"},
		Verbs:      []string{"plan", "do", "check", "act"},
		Adverbs:    []string{"fast", "really", "actually", "recently"},
	}
	fmt.Println(hruid.GenerateByEntities(obj))
}
