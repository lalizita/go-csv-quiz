package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type QuizQuestion struct {
	Order    int
	Question string
	Result   string
}

type Answers struct {
	Order      int
	UserAnswer string
	isCorrect  bool
}

func readCSVQuiz() ([][]string, error) {
	csvFile, err := os.Open("test.csv")
	if err != nil {
		log.Fatal("ERRO AO LER CSV:", err)
	}

	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	return r.ReadAll()
}

func main() {
	records, err := readCSVQuiz()

	allQuestions := make([]QuizQuestion, 0)
	answers := make([]Answers, 0)
	totalCorrect := 0

	if err != nil {
		log.Fatal("Erro ao retornar CSV:", err)
	}
	if err == io.EOF {
		log.Fatal("CSV vazio")
	}
	for index, r := range records {
		allQuestions = append(allQuestions, QuizQuestion{Order: index, Question: r[0], Result: r[1]})
		answers = append(answers, Answers{Order: index})
	}

	// totalQuestions := len(allQuestions)
	scanner := bufio.NewScanner(os.Stdin)
	for _, q := range allQuestions {
		fmt.Printf("Questão %d - %s\n", q.Order+1, q.Question)
		scanner.Scan()
		answers[q.Order].UserAnswer = scanner.Text()
	}

	for i, a := range answers {
		if a.UserAnswer == allQuestions[i].Result {
			a.isCorrect = true
			totalCorrect++
		} else {
			a.isCorrect = false
		}
	}

	fmt.Println("TOTAL DE QUESTÕES: ", len(allQuestions))
	fmt.Println("===================")
	fmt.Println("CERTAS: ", totalCorrect)
	fmt.Println("ERRADAS: ", len(allQuestions)-totalCorrect)
}
