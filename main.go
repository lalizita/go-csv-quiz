package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type QuizQuestion struct {
	Question string
	Result   string
	Answer   string
}

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal("Opening file:", err)
	}

	c := csv.NewReader(f)
	r, err := c.ReadAll()
	if err != nil {
		log.Fatal("Erro ao retornar CSV:", err)
	}
	if err == io.EOF {
		log.Fatal("CSV vazio")
	}

	allQuestions := make([]QuizQuestion, 0)
	for _, r := range r {
		allQuestions = append(allQuestions, QuizQuestion{Question: r[0], Result: r[1]})
	}

	totalCorrect := 0
	for index, q := range allQuestions {
		fmt.Printf("Questão %d - %s\n", index+1, q.Question)
		fmt.Scan(&allQuestions[index].Answer)
		if allQuestions[index].Answer == allQuestions[index].Result {
			totalCorrect++
		}
	}

	fmt.Println("TOTAL DE QUESTÕES: ", len(allQuestions))
	fmt.Println("===================")
	fmt.Println("CERTAS: ", totalCorrect)
}
