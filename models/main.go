package models

import (
	"math/rand"
	"time"
)

type Joke struct {
	Name  string
	Body  string
	Score int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (joke *Joke) Random() error {
	fetched := totallyProductionGradeDatabase[rand.Intn(len(totallyProductionGradeDatabase))]
	joke.Name = fetched.Name
	joke.Body = fetched.Body
	joke.Score = fetched.Score
	return nil
}

func (joke Joke) Save() error {
	totallyProductionGradeDatabase = append(totallyProductionGradeDatabase, joke)
	return nil
}

var totallyProductionGradeDatabase = []Joke{
	Joke{
		Name:  "Why did the chicken cross the road?",
		Body:  "To get to the other side!",
		Score: 1,
	},
	Joke{
		Name:  "Parallel lines have so much in common.",
		Body:  "It's a shame they'll never meet.",
		Score: 10,
	},
}
