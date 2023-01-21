package main

import (
	// "github.com/gin-gonic/gin"
	"encoding/csv"
	"fmt"
	"os"
)

type trainingData struct {
	Text       string
	Suspicious bool
	Racism     bool
	Sexism     bool
	Hate       bool
	Suicidal   bool
}

func parseData() (examples []trainingData) {
	csvFile, err := os.Open("server/pkg/classified_tweets.csv") //may not be correct directory
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		emp := trainingData{
			Text:       line[0],
			Suspicious: (line[1] == "1"),
			Racism:     (line[2] == "1"),
			Sexism:     (line[2] == "2"),
			Hate:       (line[3] == "1"),
			Suicidal:   (line[4] == "1"),
		}
		examples = append(examples, emp)
	}

	return
}

// func ContentMod(content string) bool {

// }

// func main() {
// port := ":3000"

// r := gin.Default()

// r.GET("/", func(c *gin.Context) {
//     c.String(200, "Hello, world!")
// })

// r.Run(port)

// co, err := cohere.CreateClient(COHERE_KEY)
// if err != nil {
//     fmt.Println(err)
//     return
// }
