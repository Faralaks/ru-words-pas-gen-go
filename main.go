package ru_words_pas_gen_go

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

var wordBase map[string]*[]string
var Words []string
var WordsLength int
var ShufflePerion time.Duration // Hour by default

func init() {
	wordBase = make(map[string]*[]string)
	wordBaseSource, err := ioutil.ReadFile("./word_base.json")
	if err != nil {
		println("Can't find word_base.json! Switch to internal word base!")
		panic(err)
	}
	err = json.Unmarshal(wordBaseSource, &wordBase)
	if err != nil {
		println("Panic!  Can't unmarshal wordBaseSource | " + err.Error())
		panic(err)
	}
	Words = *wordBase["words"]
	WordsLength = len(Words)
	ShufflePerion = time.Hour
	rand.Seed(time.Now().UnixNano())
	go Shuffle()
}

func Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for {
		rand.Shuffle(WordsLength, func(i, j int) {
			Words[i], Words[j] = Words[j], Words[i]
		})
		time.Sleep(ShufflePerion)
	}
}

func GeneratePas(minLength int, separator string) string {
	var pas string
	for pas = ""; len(pas) < minLength; {
		pas += separator + Words[rand.Intn(WordsLength)]
	}
	return pas[1:]
}
