package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Carregando Perguntas")

	file, err := os.ReadFile("questions.csv")
	if err != nil {
		log.Fatal("Erro ao carregar arquivo", err)
	}
	strFile := string(file)

	questions := strings.Split(strFile, "\n")
	total := len(questions)
	acertos := 0
	reader := bufio.NewReader(os.Stdin)

	for index, q := range questions {
		question := strings.Split(q, ";")
		fmt.Printf("%d) %s = ", index+1, question[0])
		resp, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal("Erro ao ler o resultados", err)
		}

		if question[1] == string(resp) {
			acertos++
		}

	}

	fmt.Printf("Total de acertos: %d/%d\n", acertos, total)
}
