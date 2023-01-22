package main

import (
	// "github.com/gin-gonic/gin"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/cohere-ai/cohere-go"
)

var co, err = cohere.CreateClient("VkilF9VBq2pTD9ImdKNsVXWeoojQRLlUSIfBZdwB")

type classifyResult struct {
	Text       string
	Suspicious string
	Discrim    string
	Hate       string
	Suicidal   string
}

func parseData() (examples []classifyResult) {
	// currentDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var csvDir string
	// csvDir = currentDir + "/server/pkg/classified_tweets.csv"
	// csvFile, err := os.Open(csvDir)
	csvFile, err := os.Open("/Users/angelinazhai/Desktop/Github-Repositories/palette_pal/server/pkg/classified_tweets.csv") //replace with abs path if debugging
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		emp := classifyResult{
			Text:       line[0],
			Suspicious: line[1],
			Discrim:    line[2],
			Hate:       line[3],
			Suicidal:   line[4],
		}
		if emp.Suspicious == emp.Discrim && emp.Discrim == emp.Hate && emp.Hate == emp.Suicidal && emp.Suicidal == "0" {
			continue
		} else {
			examples = append(examples, emp)
			// fmt.Println(emp)
			if len(examples) == 100 {
				break
			}
		}
	}

	return
}

func ImportExamples(checkType int, examples []classifyResult) (exampleSet []cohere.Example) {

	// var examples = parseData()
	label := ""

	for i := 0; i < 100; i++ {
		switch checkType {
		// default:
		// 	label = "na"
		case 1: //Suspicious
			if examples[i].Suspicious == "1" {
				label = "suspicious"
			} else {
				label = "na"
			}

		case 2: //Racism
			if examples[i].Discrim == "1" {
				label = "discrim"
			} else if examples[i].Discrim == "2" {
				label = "discrim"
			} else {
				label = "na"
			}

		case 3: //Hate
			if examples[i].Hate == "1" {
				label = "hate"
			} else {
				label = "na"
			}

		case 4: //Suicidal
			if examples[i].Suicidal == "1" {
				label = "suicidal"
			} else {
				label = "na"
			}

		}

		exampleSet = append(exampleSet, cohere.Example{Text: examples[i].Text, Label: label})
	}

	return
}

func ClassifyStringParse(tmp []string, category int) {
	var examples = parseData() //
	response, err := co.Classify(cohere.ClassifyOptions{
		Inputs:          tmp,
		Examples:        ImportExamples(category, examples),
		OutputIndicator: "tmp",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(response)
	}
	// prediction = response.Classifications
	return
}

func Pretrain() (exampleSet []cohere.Example) {
	trainingSet := []cohere.Example{
		{Text: "mark Ya! Dick head old man!", Label: "insult"},
		{Text: "why do you eat so much you fat bastard", Label: "insult"},
		{Text: "whooo.  did you have to choke a bitch?", Label: "insult"},
		{Text: "Testosterone Man!!!  I met Mega Bitch Soccer Mom the other day. She took the time to stop  honk her horn  & flip me off", Label: "insult"},
		{Text: "I dont want to be here anymore", Label: "negative"},
		{Text: "WOW I AM ADDING THAY TO MY LIST OF SHIT TO LISTEN TO WHEN I WANT TO DIE", Label: "negative"},
		{Text: "I'm tired of living in a world", Label: "negative"},
		{Text: "almost broke down at class just because the most negative thoughts are kicking in and I'm mf tired of living hahahahahahaha", Label: "negative"},
		{Text: "I'm so fucking tired of living I'm not smart I'm not athletic I'm not attractive I'm not talented I just have this fucking t", Label: "negative"},
		{Text: "How often do you get sick?", Label: "neutral"},
		{Text: "There there *hugs* Nobody sucks in relationship. Maybe it's about timing and the people :) It's OK", Label: "neutral"},
		{Text: "Just broke 10,000 in Xbox 360 Achievements", Label: "neutral"},
		{Text: "I could but then there would be too much people around admiring me Can you send me the address of that cava bar pleaaase?", Label: "neutral"},
		{Text: "mark Ya! Dick head old man!", Label: "neutral"},
	}
	exampleSet = append(exampleSet, trainingSet...)
	// fmt.Println(exampleSet)
	return
}

func ClassifyString() (results []cohere.Classification) {
	var exampleSet = Pretrain() //preferably preload this
	// fmt.Println(exampleSet)
	var tester = []string{"let me out let me out let me out", "i hate women", "I am happy for you"}
	// var err error
	response, err := co.Classify(cohere.ClassifyOptions{
		Inputs:          tester,
		Examples:        exampleSet,
		OutputIndicator: "tmp",
	})

	if err != nil {
		fmt.Println(err)
	}
	results = response.Classifications

	return
}

//// Keep this in to implement when frontend info input is finished
// func ContentMod(tmp []string)(labels []string) {

// }
