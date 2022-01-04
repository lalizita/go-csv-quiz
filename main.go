package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type QuizQuestion struct {
	Question string
	Result   string
}

func readCSVQuiz() ([][]string, error) {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("ERRO AO LER CSV:", err)
	}

	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	return r.ReadAll()
}

func main() {
	allQuestions := make([]QuizQuestion, 0)
	records, err := readCSVQuiz()
	if err != nil {
		log.Fatal("Erro ao retornar CSV:", err)
	}
	if err == io.EOF {
		log.Fatal("CSV vazio")
	}
	for _, r := range records {
		allQuestions = append(allQuestions, QuizQuestion{Question: r[0], Result: r[1]})
	}
}
